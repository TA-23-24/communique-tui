// Copyright 2019 The Mellium Contributors.
// Use of this source code is governed by the BSD 2-clause
// license that can be found in the LICENSE file.

// Package event contains events that may be emited by the client.
package event // import "mellium.im/communique/internal/client/event"

import (
	"mellium.im/xmpp/commands"
	"mellium.im/xmpp/delay"
	"mellium.im/xmpp/forward"
	"mellium.im/xmpp/jid"
	"mellium.im/xmpp/roster"
	"mellium.im/xmpp/stanza"
)

type (
	// Connected is sent immediately after the connection is esablished.
	Connected struct{}

	// StatusOnline is sent when the user should come online.
	StatusOnline jid.JID

	// StatusOffline is sent when the user should go offline.
	StatusOffline jid.JID

	// StatusAway is sent when the user should change their status to away.
	StatusAway jid.JID

	// StatusBusy is sent when the user should change their status to busy.
	StatusBusy jid.JID

	// LoadingCommands is sent by the UI when the ad-hoc command window opens.
	LoadingCommands jid.JID

	// ExecCommand is sent by the UI when an ad-hoc command should be executed.
	ExecCommand commands.Command

	// FetchRoster is sent when a roster is fetched.
	FetchRoster struct {
		Ver   string
		Items <-chan UpdateRoster
	}

	// DeleteRosterItem is sent when a roster item has been removed (eg. after
	// UpdateRoster triggers a removal or it is removed in the UI).
	DeleteRosterItem roster.Item

	// UpdateRoster is sent when a roster item should be updated (eg. after a
	// roster push).
	UpdateRoster struct {
		roster.Item
		Ver  string
		Room bool
	}

	// ChatMessage is sent when messages of type "chat" or "normal" are received
	// or sent.
	ChatMessage struct {
		stanza.Message

		Body     string          `xml:"body,omitempty"`
		OriginID stanza.OriginID `xml:"urn:xmpp:sid:0 origin-id"`
		SID      []stanza.ID     `xml:"urn:xmpp:sid:0 stanza-id"`
		Delay    delay.Delay     `xml:"urn:xmpp:delay delay"`

		// Sent is true if this message is one that we sent from another device (for
		// example, a message forwarded to us by message carbons).
		Sent bool `xml:"-"`
		// Account is true if this message was sent by the server (empty from, or
		// from matching the bare JID of the authenticated account).
		Account bool `xml:"-"`
	}

	// HistoryMessage is sent on incoming messages resulting from a history query.
	HistoryMessage struct {
		stanza.Message
		Result struct {
			QueryID string `xml:"queryid,attr"`
			ID      string `xml:"id,attr"`
			Forward struct {
				forward.Forwarded
				Msg ChatMessage `xml:"jabber:client message"`
			} `xml:"urn:xmpp:forward:0 forwarded"`
		} `xml:"urn:xmpp:mam:2 result"`
	}

	// Receipt is sent when a message receipt is received and represents the ID of
	// the message that should be marked as received.
	// It may be sent by itself, or in addition to a ChatMessage event (or others)
	// if the payload containing the receipt also contains other events.
	Receipt string

	// OpenChat is sent when a roster item is selected.
	OpenChat roster.Item

	// CloseChat is sent when the chat view is closed.
	CloseChat roster.Item

	// Subscribe is sent when we subscribe to a users presence.
	Subscribe jid.JID

	// PullToRefreshChat is sent when we scroll up while already at the top of
	// the history or when we simply scroll to the top of the history.
	PullToRefreshChat roster.Item
)
