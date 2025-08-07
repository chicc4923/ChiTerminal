package main

import (
	"testing"
)

func TestGreeting(t *testing.T) {
	out, _ := makeUI()

	if out.Text != "Hello world!" {
		t.Error("Incorrect initial greeting")
	}
}
