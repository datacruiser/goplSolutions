package main

import (
	"testing"
)

func TestConsts(t *testing.T) {
	if got := KB; got != 1024 {
		t.Errorf("KB = %d, want 1024", got)
	}
	if got := MB; got != 1048576 {
		t.Errorf("MB = %d, want 1048576", got)
	}
	if got := MB / KB; got != 1024 {
		t.Errorf("MB / KB = %d, want 1024", got)
	}
	if got := TB / MB; got != 1048576 {
		t.Errorf("TB / MB = %d, want 1048576", got)
	}
}
