package mood

import (
	"testing"
)

func TestMood(t *testing.T) {
	if Status != "happy" {
		t.Fatal("Not happy, now I am sad")
	}
}
