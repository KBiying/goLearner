package main

import (
	"testing"
)

func Search(dictionary map[string]string, word string) string {
	return dictionary[word]
}

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})
	t.Run("unkown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := ErrNotFound.Error()

		if err == nil {
			t.Fatal("expected to get an error")
		}

		assertStrings(t, err.Error(), want)
		assertError(t, err, ErrNotFound)

	})
}

func TestAdd(t *testing.T) {
	t.Run("new word ", func(t *testing.T) {
		dictionary := Dictionary{}
		key := "test"
		value := "this is just a test"
		dictionary.Add(key, value)
		assertDefinition(t, dictionary, key, value)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this id just a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)

	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		newDeinition := "new definition"
		dictionary.Update(word, newDeinition)

		assertDefinition(t, dictionary, word, newDeinition)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)
		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	t.Run("delete", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{word: "test definision"}
		dictionary.Delete(word)

		_, err := dictionary.Search(word)
		if err != ErrNotFound {
			t.Errorf("Expected '%s' to be deleted", word)
		}
	})
}

func assertDefinition(t *testing.T, dictionary Dictionary, key, value string) {
	t.Helper()
	got, err := dictionary.Search(key)

	if err != nil {
		t.Fatal("should find added word:", err)
	}
	if value != got {
		t.Errorf("got '%s' want '%s'", got, value)
	}
}

func assertStrings(t *testing.T, got, want string) {
	if got != want {
		t.Errorf("got '%s' want '%s' given , '%s'", got, want, "test")
	}
}

func assertError(t *testing.T, got, want error) {
	// t.Helper()

	if got != want {
		t.Errorf("got error '%s' want '%s'", got, want)
	}
}

//  func TestSearch(t *testing.T){
// 	dictionary := map[string]string{"test":"this is just a test"}
// 	got := Search(dictionary, "test")
// 	want:= "this is just a test"

// 	if got!=want{
// 		t.Errorf("got '%s' want '%s' given , '%s'", got, want, "test"

// 	}
//  }
