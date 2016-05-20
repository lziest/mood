package mood

import (
	"testing"
)

func TestMood(t *testing.T) {
	if Status != "sad" {
		t.Fatal("Not sad, now I am sad")
	}
}
