package main

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func Countdown(out io.Writer) {
	fmt.Fprintf(out, "3")
}

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}

	Countdown(buffer)
	got := buffer.String()
	want := "3"
	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
