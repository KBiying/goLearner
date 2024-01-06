package main

import (
	"testing"
)

func Search(dictionary map[string]string, word string) string{
	return dictionary[word]
}

 func TestSearch(t *testing.T){
	dictionary := map[string]string{"test":"this is just a test"}
	got := Search(dictionary, "test")
	want:= "this is just a test"

	if got!=want{
		t.Errorf("got '%s' want '%s' given , '%s'", got, want, "test")

	}
 }