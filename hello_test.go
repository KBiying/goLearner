package main

import "testing"

func TestHello(t *testing.T){
	got:=Hello("kb")
	want := "Hello, kb"

	if got!= want{
		t.Errorf("got '%q' want '%q'",got, want)
	}
}