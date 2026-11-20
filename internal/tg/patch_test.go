package tg

import (
	"testing"

	tgapi "github.com/gotd/td/tg"
)

func TestReplyToID(t *testing.T) {
	plain := &tgapi.Message{ID: 1}
	if got := replyToID(plain); got != 0 {
		t.Fatalf("no reply: got %d, want 0", got)
	}
	h := &tgapi.MessageReplyHeader{}
	h.SetReplyToMsgID(42)
	reply := &tgapi.Message{ID: 2}
	reply.SetReplyTo(h)
	if got := replyToID(reply); got != 42 {
		t.Fatalf("reply: got %d, want 42", got)
	}
}

func TestScheduledMessage(t *testing.T) {
	m := &tgapi.Message{ID: 7, Date: 1000}
	u := &tgapi.Updates{Updates: []tgapi.UpdateClass{
		&tgapi.UpdateMessageID{ID: 7},
		&tgapi.UpdateNewScheduledMessage{Message: m},
	}}
	if got := scheduledMessage(u); got == nil || got.ID != 7 {
		t.Fatalf("got %v, want message 7", got)
	}
	if got := scheduledMessage(&tgapi.Updates{}); got != nil {
		t.Fatalf("empty updates: got %v, want nil", got)
	}
}
