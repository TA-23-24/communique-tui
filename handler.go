// Copyright 2018 The Mellium Contributors.
// Use of this source code is governed by the BSD 2-clause
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"log"
	"time"

	"mellium.im/communique/internal/client"
	"mellium.im/communique/internal/client/event"
	"mellium.im/communique/internal/storage"
	"mellium.im/communique/internal/ui"
	"mellium.im/xmpp/jid"
	"mellium.im/xmpp/roster"
	"mellium.im/xmpp/stanza"
)

// newUIHandler returns a handler for events that are emitted by the UI that
// need to modify the client state.
func newUIHandler(configPath string, pane *ui.UI, db *storage.DB, c *client.Client, logger, debug *log.Logger) func(interface{}) {
	return func(ev interface{}) {
		switch e := ev.(type) {
		case event.StatusAway:
			go func() {
				ctx, cancel := context.WithTimeout(context.Background(), c.Timeout())
				defer cancel()
				if err := c.Away(ctx); err != nil {
					logger.Printf("error setting away status: %v", err)
				}
			}()
		case event.StatusOnline:
			go func() {
				ctx, cancel := context.WithTimeout(context.Background(), c.Timeout())
				defer cancel()
				if err := c.Online(ctx); err != nil {
					logger.Printf("error setting online status: %v", err)
				}
			}()
		case event.StatusBusy:
			go func() {
				ctx, cancel := context.WithTimeout(context.Background(), c.Timeout())
				defer cancel()
				if err := c.Busy(ctx); err != nil {
					logger.Printf("error setting busy status: %v", err)
				}
			}()
		case event.StatusOffline:
			go func() {
				if err := c.Offline(); err != nil {
					logger.Printf("error going offline: %v", err)
				}
			}()
		case event.DeleteRosterItem:
			go func() {
				ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				defer cancel()
				err := roster.Delete(ctx, c.Session, e.JID)
				if err != nil {
					logger.Printf("error removing roster item %s: %v", e.JID, err)
				}
			}()
		case event.UpdateRoster:
			// TODO:
			panic("event.UpdateRoster: not yet implemented")
		case event.ChatMessage:
			go func() {
				ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				defer cancel()

				var err error
				e, err = c.SendMessage(ctx, e)
				if err != nil {
					logger.Printf("error sending message: %v", err)
				}
				if err = writeMessage(pane, configPath, e, false); err != nil {
					logger.Printf("error saving sent message to history: %v", err)
				}
				if err = db.InsertMsg(ctx, e); err != nil {
					logger.Printf("error writing message to database: %v", err)
				}
				// If we sent the message that wasn't automated (it has a body), assume
				// we've read everything before it.
				if e.Sent && e.Body != "" {
					pane.Roster().MarkRead(e.To.Bare().String())
				}
			}()
		case event.OpenChat:
			go func() {
				var firstUnread string
				item, ok := pane.Roster().GetItem(e.JID.Bare().String())
				if ok {
					firstUnread = item.FirstUnread()
				}
				ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
				defer cancel()
				if err := loadBuffer(ctx, pane, db, configPath, e, firstUnread, logger); err != nil {
					logger.Printf("error loading chat: %v", err)
					return
				}
				history := pane.History()
				history.ScrollToEnd()
				pane.Roster().MarkRead(e.JID.Bare().String())
			}()
		case event.CloseChat:
			history := pane.History()
			history.SetText("")
		case event.Subscribe:
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			err := c.Send(ctx, stanza.Presence{
				To:   jid.JID(e),
				Type: stanza.SubscribedPresence,
			}.Wrap(nil))
			if err != nil {
				logger.Printf("error sending presence pre-approval to %s: %v", jid.JID(e), err)
			}
			err = c.Send(ctx, stanza.Presence{
				To:   jid.JID(e),
				Type: stanza.SubscribePresence,
			}.Wrap(nil))
			if err != nil {
				logger.Printf("error sending presence request to %s: %v", jid.JID(e), err)
			}
		default:
			debug.Printf("unrecognized ui event: %q", e)
		}
	}
}

// newClientHandler returns a handler for events that are emitted by the client
// that need to modify the UI.
func newClientHandler(configPath string, pane *ui.UI, db *storage.DB, logger, debug *log.Logger) func(interface{}) {
	return func(ev interface{}) {
		switch e := ev.(type) {
		case event.StatusAway:
			pane.Away()
		case event.StatusBusy:
			pane.Busy()
		case event.StatusOnline:
			pane.Online()
		case event.StatusOffline:
			pane.Offline()
		case event.FetchRoster:
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			err := db.ReplaceRoster(ctx, e)
			if err != nil {
				logger.Printf("error fetching roster ver %q: %v", e.Ver, err)
			}
			err = db.ForRoster(ctx, func(item event.UpdateRoster) {
				pane.UpdateRoster(ui.RosterItem{Item: roster.Item(item.Item)})
			})
			if err != nil {
				logger.Printf("error iterating over roster items: %v", err)
			}
		case event.UpdateRoster:
			pane.UpdateRoster(ui.RosterItem{Item: roster.Item(e.Item)})
			ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
			defer cancel()
			db.UpdateRoster(ctx, e.Ver, e)
		case event.Receipt:
			ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
			defer cancel()
			err := db.MarkReceived(ctx, e)
			if err != nil {
				logger.Printf("error marking message %q as received: %v", e, err)
			}
		case event.ChatMessage:
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			if err := writeMessage(pane, configPath, e, false); err != nil {
				logger.Printf("error writing received message to history: %v", err)
			}
			if err := db.InsertMsg(ctx, e); err != nil {
				logger.Printf("error writing message to database: %v", err)
			}
			// If we sent the message that wasn't automated (it has a body), assume
			// we've read everything before it.
			if e.Sent && e.Body != "" {
				pane.Roster().MarkRead(e.To.Bare().String())
			}
		default:
			debug.Printf("unrecognized client event: %q", e)
		}
	}
}
