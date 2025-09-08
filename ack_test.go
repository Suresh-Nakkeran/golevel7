package golevel7

import (
	"errors"
	"os"
	"testing"
)

func TestAcknowledge(t *testing.T) {
	fname := "./testdata/msg.hl7"
	file, err := os.Open(fname)
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	msgs, err := NewDecoder(file).Messages()
	if err != nil {
		t.Fatal(err)
	}
	mi, err := msgs[0].Info()
	ack := Acknowledge(mi, nil)
	if ack == nil {
		t.Fatal("Expected ACK message got nil")
	}
	ack = Acknowledge(mi, errors.New("This is a test error"))
	if ack == nil {
		t.Fatal("Expected ACK message got nil")
	}
	m := NewMsgInfo()
	m.ReceivingApp = "ORG_REC_APP"
	m.ReceivingFacility = "ORG_REC_FAC"
	m.SendingApp = "ORG_SEND_APP"
	m.SendingFacility = "ORG_SEND_FAC"
	ack = Acknowledge(*m, errors.New("Fatal error"))
	if ack == nil {
		t.Fatal("Expected ACK message got nil")
	}
}
