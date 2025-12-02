package day8

import "testing"

func TestNotifieruser(t *testing.T) {
	fake := &Fakenotifier{}

	Notifyuser(fake, "hey!")
	if fake.Received != "hey!" {
		t.Errorf("Expected hey! got %s", fake.Received)
	}

}
