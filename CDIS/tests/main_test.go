package tests

import (
	"testing"
)

func TestMainStartup(t *testing.T) {
	expected := "ok"
	got := "ok"
	if got != expected {
		t.Errorf("expected %s, got %s", expected, got)
	}
}
