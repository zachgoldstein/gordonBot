package speech

import (
	"reflect"
	"testing"
)

func TestSplitMessage(t *testing.T) {
	msgParts, err := SplitMessage("June Zach Goldstein")
	if err != nil {
		t.Error("Unexpected error ", err)
	}
	if len(msgParts) != 3 {
		t.Error("Expected length of 3 for message parts, got: ", len(msgParts))
	}

	expectedMsgParts := []string{"June", "Zach", "Goldstein"}
	if !reflect.DeepEqual(expectedMsgParts, msgParts) {
		t.Errorf("Expected message parts %s, got: %s", expectedMsgParts, msgParts)
	}

	_, err = SplitMessage("BLARG June Zach Goldstein")
	if err == nil {
		t.Error("Expected error, but no error was created")
	}
}

func TestCreateMessage(t *testing.T) {
	expectedMessage := "Awesome Job You Clever Interdimensional Being"

	msg, err := CreateMessage("June", "Zach", "Goldstein")
	if err != nil {
		t.Error("Unexpected error ", err)
	}
	if msg != expectedMessage {
		t.Errorf("Expected message '%s' but received message ''%s'", expectedMessage, msg)
	}

	_, err = CreateMessage("ImaginaryMonthruary", "Zach", "Goldstein")
	if err == nil {
		t.Errorf("Expected error for non-existent month but did not receive one")
	}
}
