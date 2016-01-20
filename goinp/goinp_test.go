package goinp

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAskForStringFromReader(t *testing.T) {
	t.Log("TestAskForString")

	testUserInput := "this is some text"

	res, err := AskForStringFromReader("Enter some text", strings.NewReader(testUserInput))
	require.NoError(t, err)
	if res != testUserInput {
		t.Fatalf("Scanned input (%s) does not match expected (%s)", res, testUserInput)
	}
}

func TestAskForIntFromReader(t *testing.T) {
	t.Log("TestAskForString")

	testUserInput := "31"

	res, err := AskForIntFromReader("Enter a number", strings.NewReader(testUserInput))
	require.NoError(t, err)
	if res != 31 {
		t.Fatalf("Scanned input (%s) does not match expected (%s)", res, testUserInput)
	}
}

func TestAskForBoolFromReader(t *testing.T) {
	t.Log("TestAskForString")

	// yes
	testUserInput := "y"
	res, err := AskForBoolFromReader("Yes or no?", strings.NewReader(testUserInput))
	require.NoError(t, err)
	if res != true {
		t.Fatalf("Scanned input (%s) does not match expected (%s)", res, testUserInput)
	}

	// no
	testUserInput = "no"
	res, err = AskForBoolFromReader("Yes or no?", strings.NewReader(testUserInput))
	require.NoError(t, err)
	if res != false {
		t.Fatalf("Scanned input (%s) does not match expected (%s)", res, testUserInput)
	}
}

func TestParseBool(t *testing.T) {
	t.Log("Simple Yes")
	testUserInput := "y"
	isYes, err := ParseBool("YeS")
	require.NoError(t, err)
	if !isYes {
		t.Fatalf("Scanned input (%s) does not match expected (%s)", isYes, testUserInput)
	}

	t.Log("Simple No")
	testUserInput = "no"
	isYes, err = ParseBool("n")
	require.NoError(t, err)
	if isYes {
		t.Fatalf("Scanned input (%s) does not match expected (%s)", isYes, testUserInput)
	}

	t.Log("Newline in yes - trim")
	testUserInput = `
 yes
`
	isYes, err = ParseBool(testUserInput)
	require.NoError(t, err)
	if !isYes {
		t.Fatalf("Scanned input (%s) does not match expected (%s)", isYes, testUserInput)
	}
}

func TestSelectFromStringsFromReader(t *testing.T) {
	availableOptions := []string{"first", "second", "third"}

	//
	{
		_, err := SelectFromStringsFromReader("Select something", availableOptions, strings.NewReader("-1"))
		require.EqualError(t, err, "Invalid option: You entered a number less than 1")
	}
	//
	{
		_, err := SelectFromStringsFromReader("Select something", availableOptions, strings.NewReader("0"))
		require.EqualError(t, err, "Invalid option: You entered a number less than 1")
	}
	//
	{
		res, err := SelectFromStringsFromReader("Select something", availableOptions, strings.NewReader("1"))
		require.NoError(t, err)
		require.Equal(t, "first", res)
	}
	//
	{
		res, err := SelectFromStringsFromReader("Select something", availableOptions, strings.NewReader("2"))
		require.NoError(t, err)
		require.Equal(t, "second", res)
	}
	//
	{
		res, err := SelectFromStringsFromReader("Select something", availableOptions, strings.NewReader("3"))
		require.NoError(t, err)
		require.Equal(t, "third", res)
	}
	//
	{
		_, err := SelectFromStringsFromReader("Select something", availableOptions, strings.NewReader("4"))
		require.EqualError(t, err, "Invalid option: You entered a number greater than the last option's number")
	}
}
