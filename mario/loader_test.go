package mario

import (
	"strings"
	"testing"
)

func TestLoaderLocal(t *testing.T) {
	m := n()

	if m.Name != "mario" {
		t.Logf("The name of this applicaton should be mario")
		t.FailNow()
	}

	if !strings.Contains(string(m.Instructions), "tdd:") {
		t.Logf("Unable to load the instructions properly")
		t.FailNow()
	}
}

func TestLoaderRemote(t *testing.T) {
	m := nRemote()
	if m.Name != "alfred" {
		t.Logf("The name of this applicaton should be alfred ;)")
		t.FailNow()
	}

	if !strings.Contains(string(m.Instructions), "setup:") {
		t.Logf("Fetching remote URL instructions failed")
		t.FailNow()
	}
}
