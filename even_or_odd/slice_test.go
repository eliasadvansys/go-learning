package main

import "testing"

func TestNewSlice(t *testing.T) {
	s := newSlice(10)

	if len(s) != 10 {
		t.Errorf("Expected 10, got %d", len(s))
	}
}
