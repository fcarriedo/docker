package units

import (
	"testing"
	"time"
)

func TestHumanDuration(t *testing.T) {
	assertHumanDuration(t, "Less than a second", HumanDuration(time.Duration(1)*time.Millisecond))
}

func assertHumanDuration(t *testing.T, expected, actual string) {
	if expected != actual {
		t.Errorf("Expected '%s' but got '%s'", expected, actual)
	}
}
