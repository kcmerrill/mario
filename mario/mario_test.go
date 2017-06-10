package mario

import (
	"testing"
)

func n() *Mario {
	return New([]string{"mario"})
}

func TestNew(t *testing.T) {
	m := n()
	if m.Name != "mario" {
		t.Logf("The name of this applicaton should be mario")
		t.Logf(file)
		t.FailNow()
	}
}
