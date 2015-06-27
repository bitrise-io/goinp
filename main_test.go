package main

import (
	"testing"
	"strings"
)

func TestAskForStringFromReader(t *testing.T) {
	t.Log("TestAskForString")

	testUserInput := "this is some text"

	res, err := AskForStringFromReader("Enter some text", strings.NewReader(testUserInput))
	if err != nil {
		t.Fatal(err)
	}
	if res != testUserInput {
		t.Fatalf("Scanned input (%s) does not match expected (%s)", res, testUserInput)
	}
}

func TestAskForIntFromReader(t *testing.T) {
	t.Log("TestAskForString")

	testUserInput := "31"

	res, err := AskForIntFromReader("Enter a number", strings.NewReader(testUserInput))
	if err != nil {
		t.Fatal(err)
	}
	if res != 31 {
		t.Fatalf("Scanned input (%s) does not match expected (%s)", res, testUserInput)
	}
}
