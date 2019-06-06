package goinp

import (
	"bytes"
	"strings"
	"syscall"
	"testing"

	"github.com/stretchr/testify/require"
)

//=======================================
// String
//=======================================

func TestAskForStringFromReaderWithDefault(t *testing.T) {
	t.Log("TestAskForString - input, NO default value")
	{
		testUserInput := "this is some text"

		res, err := AskForStringFromReaderWithDefault("Enter some text", "", strings.NewReader(testUserInput))
		require.NoError(t, err)
		require.Equal(t, testUserInput, res)
	}

	t.Log("TestAskForString - input, default value")
	{
		testUserInput := "this is some text"
		defaultValue := "default"

		res, err := AskForStringFromReaderWithDefault("Enter some text", defaultValue, strings.NewReader(testUserInput))
		require.NoError(t, err)
		require.Equal(t, testUserInput, res)
	}

	t.Log("TestAskForString - NO input, NO default value")
	{
		testUserInput := ""

		res, err := AskForStringFromReaderWithDefault("Enter some text", "", strings.NewReader(testUserInput))
		require.Error(t, err)
		require.Equal(t, testUserInput, res)
	}

	t.Log("TestAskForString - NO input, default value")
	{
		testUserInput := ""
		defaultValue := "default"

		res, err := AskForStringFromReaderWithDefault("Enter some text", defaultValue, strings.NewReader(testUserInput))
		require.NoError(t, err)
		require.Equal(t, defaultValue, res)
	}
}

func TestAskForStringFromReader(t *testing.T) {
	t.Log("TestAskForStringFromReader - input")
	{
		testUserInput := "this is some text"

		res, err := AskForStringFromReader("Enter some text", strings.NewReader(testUserInput))
		require.NoError(t, err)
		require.Equal(t, testUserInput, res)
	}

	t.Log("TestAskForStringFromReader - NO input")
	{
		testUserInput := ""

		res, err := AskForStringFromReader("Enter some text", strings.NewReader(testUserInput))
		require.Error(t, err)
		require.Equal(t, testUserInput, res)
	}
}

//=======================================
// Path
//=======================================

func TestAskForPathFromReaderWithDefault(t *testing.T) {
	t.Log("Simple path - input, NO default value")
	{
		testUserInput := "path/without/spaces"
		defaultValue := ""

		res, err := AskForPathFromReaderWithDefault("Enter a path", defaultValue, strings.NewReader(testUserInput))
		require.NoError(t, err)
		require.Equal(t, testUserInput, res)
	}

	t.Log("TestAskForPathFromReaderWithDefault - input, with default value")
	{
		testUserInput := "path/without/spaces"
		defaultValue := "default"

		res, err := AskForPathFromReaderWithDefault("Enter a path", defaultValue, strings.NewReader(testUserInput))
		require.NoError(t, err)
		require.Equal(t, testUserInput, res)
	}

	t.Log("TestAskForPathFromReaderWithDefault - NO input, NO default value")
	{
		testUserInput := ""
		defaultValue := ""

		res, err := AskForPathFromReaderWithDefault("Enter a path", defaultValue, strings.NewReader(testUserInput))
		require.Error(t, err)
		require.Equal(t, defaultValue, res)
	}

	t.Log("TestAskForPathFromReaderWithDefault - input, with default value")
	{
		testUserInput := ""
		defaultValue := "default"

		res, err := AskForPathFromReaderWithDefault("Enter a path", defaultValue, strings.NewReader(testUserInput))
		require.NoError(t, err)
		require.Equal(t, defaultValue, res)
	}
}

func TestAskForPathFromReader(t *testing.T) {
	t.Log("TestAskForPathFromReader - Empty path")
	{
		testUserInput := ""
		res, err := AskForPathFromReader("Enter a path", strings.NewReader(testUserInput))
		require.Error(t, err)
		require.Equal(t, testUserInput, res)
	}

	t.Log("TestAskForPathFromReader - Simple path")
	{
		testUserInput := "path/without/spaces"
		res, err := AskForPathFromReader("Enter a path", strings.NewReader(testUserInput))
		require.NoError(t, err)
		require.Equal(t, "path/without/spaces", res)
	}

	t.Log("TestAskForPathFromReader - Path with simple spaces")
	{
		testUserInput := "path/with simple/spaces"
		res, err := AskForPathFromReader("Enter a path", strings.NewReader(testUserInput))
		require.NoError(t, err)
		require.Equal(t, "path/with simple/spaces", res)
	}

	t.Log("TestAskForPathFromReader - Path with backspace escaped space")
	{
		testUserInput := "path/with\\ spaces/in it"
		res, err := AskForPathFromReader("Enter a path", strings.NewReader(testUserInput))
		require.NoError(t, err)
		require.Equal(t, "path/with spaces/in it", res)
	}
}

//=======================================
// Int
//=======================================

func TestAskForIntFromReaderWithDefault(t *testing.T) {
	t.Log("TestAskForIntFromReaderWithDefault - input, with default value")
	{
		testUserInput := "31"
		defaultValue := 1

		res, err := AskForIntFromReaderWithDefault("Enter a number", defaultValue, strings.NewReader(testUserInput))
		require.NoError(t, err)
		require.Equal(t, int64(31), res)
	}

	t.Log("TestAskForIntFromReaderWithDefault - NO input, with default value")
	{
		testUserInput := ""
		defaultValue := 1

		res, err := AskForIntFromReaderWithDefault("Enter a number", defaultValue, strings.NewReader(testUserInput))
		require.NoError(t, err)
		require.Equal(t, int64(defaultValue), res)
	}
}

func TestAskForIntFromReader(t *testing.T) {
	t.Log("TestAskForIntFromReader - input")
	{
		testUserInput := "31"

		res, err := AskForIntFromReader("Enter a number", strings.NewReader(testUserInput))
		require.NoError(t, err)
		require.Equal(t, int64(31), res)
	}

	t.Log("TestAskForIntFromReader - NO input")
	{
		testUserInput := ""

		res, err := AskForIntFromReader("Enter a number", strings.NewReader(testUserInput))
		require.Error(t, err)
		require.Equal(t, int64(0), res)
	}
}

//=======================================
// Bool
//=======================================

func TestAskForBoolFromReaderWithDefaultValue(t *testing.T) {
	t.Log("TestAskForBoolFromReaderWithDefaultValue - input, default value")
	{
		testUserInput := "y"
		defaultValue := true

		res, err := AskForBoolFromReaderWithDefaultValue("Yes or no?", defaultValue, strings.NewReader(testUserInput))
		require.NoError(t, err)
		require.Equal(t, true, res)
	}

	t.Log("TestAskForBoolFromReaderWithDefaultValue - input, default value")
	{
		testUserInput := "n"
		defaultValue := true

		res, err := AskForBoolFromReaderWithDefaultValue("Yes or no?", defaultValue, strings.NewReader(testUserInput))
		require.NoError(t, err)
		require.Equal(t, false, res)
	}

	t.Log("TestAskForBoolFromReaderWithDefaultValue - NO input, default value")
	{
		testUserInput := ""
		defaultValue := true

		res, err := AskForBoolFromReaderWithDefaultValue("Yes or no?", defaultValue, strings.NewReader(testUserInput))
		require.NoError(t, err)
		require.Equal(t, defaultValue, res)
	}

	t.Log("TestAskForBoolFromReaderWithDefaultValue - INVALID input, default value")
	{
		testUserInput := "invalid"
		defaultValue := true

		res, err := AskForBoolFromReaderWithDefaultValue("Yes or no?", defaultValue, strings.NewReader(testUserInput))
		require.Error(t, err)
		require.Equal(t, false, res)
	}
}

func TestAskForBoolFromReader(t *testing.T) {
	t.Log("TestAskForBoolFromReader - yes")
	{
		testUserInput := "y"
		res, err := AskForBoolFromReader("Yes or no?", strings.NewReader(testUserInput))
		require.NoError(t, err)
		require.Equal(t, true, res)
	}

	t.Log("TestAskForBoolFromReader - no")
	{
		testUserInput := "n"
		res, err := AskForBoolFromReader("Yes or no?", strings.NewReader(testUserInput))
		require.NoError(t, err)
		require.Equal(t, false, res)
	}

	t.Log("TestAskForBoolFromReader - 1")
	{
		testUserInput := "-1"
		res, err := AskForBoolFromReader("Yes or no?", strings.NewReader(testUserInput))
		require.Error(t, err)
		require.Equal(t, false, res)
	}
}

func TestParseBool(t *testing.T) {
	t.Log("TestParseBool - Simple Yes")
	{
		testUserInput := "yes"
		isYes, err := ParseBool(testUserInput)
		require.NoError(t, err)
		require.Equal(t, true, isYes)
	}

	t.Log("TestParseBool - Simple true")
	{
		testUserInput := "true"
		isYes, err := ParseBool(testUserInput)
		require.NoError(t, err)
		require.Equal(t, true, isYes)
	}

	t.Log("TestParseBool - y")
	{
		testUserInput := "y"
		isYes, err := ParseBool(testUserInput)
		require.NoError(t, err)
		require.Equal(t, true, isYes)
	}

	t.Log("TestParseBool - Simple No")
	{
		testUserInput := "no"
		isYes, err := ParseBool(testUserInput)
		require.NoError(t, err)
		require.Equal(t, false, isYes)
	}

	t.Log("TestParseBool - Simple false")
	{
		testUserInput := "false"
		isYes, err := ParseBool(testUserInput)
		require.NoError(t, err)
		require.Equal(t, false, isYes)
	}

	t.Log("TestParseBool - n")
	{
		testUserInput := "n"
		isYes, err := ParseBool(testUserInput)
		require.NoError(t, err)
		require.Equal(t, false, isYes)
	}

	t.Log("TestParseBool - Newline in yes - trim")
	{
		testUserInput := `
 yes
`
		isYes, err := ParseBool(testUserInput)
		require.NoError(t, err)
		require.Equal(t, true, isYes)
	}

	t.Log("TestParseBool - With number - 1")
	{
		testUserInput := "1"
		isYes, err := ParseBool(testUserInput)
		require.NoError(t, err)
		require.Equal(t, true, isYes)
	}

	t.Log("TestParseBool - With number - 0")
	{
		testUserInput := "0"
		isYes, err := ParseBool(testUserInput)
		require.NoError(t, err)
		require.Equal(t, false, isYes)
	}

	t.Log("TestParseBool - With INVALID number - -1")
	{
		testUserInput := "-1"
		isYes, err := ParseBool(testUserInput)
		require.Error(t, err)
		require.Equal(t, false, isYes)
	}
}

//=======================================
// Select
//=======================================

func TestSelectFromStringsFromReaderWithDefault(t *testing.T) {
	availableOptions := []string{"first", "second", "third"}
	defaultValue := 3

	t.Log("TestSelectFromStringsFromReaderWithDefault - input, with default value")
	{
		res, err := SelectFromStringsFromReaderWithDefault("Select something", defaultValue, availableOptions, strings.NewReader("1"))
		require.NoError(t, err)
		require.Equal(t, "first", res)
	}

	t.Log("TestSelectFromStringsFromReaderWithDefault - NO input, with default value")
	{
		res, err := SelectFromStringsFromReaderWithDefault("Select something", defaultValue, availableOptions, strings.NewReader(""))
		require.NoError(t, err)
		require.Equal(t, "third", res)
	}

	t.Log("TestSelectFromStringsFromReaderWithDefault - INVALID input, with default value")
	{
		res, err := SelectFromStringsFromReaderWithDefault("Select something", defaultValue, availableOptions, strings.NewReader("-1"))
		require.Error(t, err)
		require.Equal(t, "", res)
	}
}

func TestSelectFromStringsFromReader(t *testing.T) {
	availableOptions := []string{"first", "second", "third"}

	t.Log("TestSelectFromStringsFromReader - NO input")
	{
		_, err := SelectFromStringsFromReader("Select something", availableOptions, strings.NewReader(""))
		require.Error(t, err)
	}

	t.Log("TestSelectFromStringsFromReader - INVALID input")
	{
		_, err := SelectFromStringsFromReader("Select something", availableOptions, strings.NewReader("-1"))
		require.EqualError(t, err, "invalid option: You entered a number less than 1")
	}

	t.Log("TestSelectFromStringsFromReader - input")
	{
		res, err := SelectFromStringsFromReader("Select something", availableOptions, strings.NewReader("1"))
		require.NoError(t, err)
		require.Equal(t, "first", res)
	}

	t.Log("TestSelectFromStringsFromReader - input")
	{
		res, err := SelectFromStringsFromReader("Select something", availableOptions, strings.NewReader("2"))
		require.NoError(t, err)
		require.Equal(t, "second", res)
	}

	t.Log("TestSelectFromStringsFromReader - input")
	{
		res, err := SelectFromStringsFromReader("Select something", availableOptions, strings.NewReader("3"))
		require.NoError(t, err)
		require.Equal(t, "third", res)
	}

	t.Log("TestSelectFromStringsFromReader - INVALID input")
	{
		_, err := SelectFromStringsFromReader("Select something", availableOptions, strings.NewReader("4"))
		require.EqualError(t, err, "invalid option: You entered a number greater than the last option's number")
	}
}

type testStdin struct {
	Stdin *bytes.Buffer
}

func (testStdin) Fd() uintptr {
	return uintptr(syscall.Stdin)
}

func (t testStdin) Read(b []byte) (int, error) {
	return t.Stdin.Read(b)
}

func testAskOptionWrapper(title string, defaultValue string, optional bool, stdin string, options ...string) (string, string, error) {
	sio := testStdin{
		bytes.NewBufferString(stdin),
	}
	var outBuf bytes.Buffer

	s, err := askOptions(title, defaultValue, optional, sio, &outBuf, options...)
	if err != nil {
		return "", "", err
	}

	return s, outBuf.String(), nil
}

func Test_AskOptions(t *testing.T) {
	//// required user input checks

	answer, stdout, err := testAskOptionWrapper("title", "", false, "single non-optional input\n")
	if err != nil {
		t.Fatal(err)
	}
	if answer != "single non-optional input" {
		t.Fatalf("invalid value returned: %s", answer)
	}
	if stdout != "Enter value for \"title\": " {
		t.Fatalf("invalid stdout: %s", stdout)
	}

	// empty when required
	answer, stdout, err = testAskOptionWrapper("title", "", false, "\nshould return this\n")
	if err != nil {
		t.Fatal(err)
	}
	if answer != "should return this" {
		t.Fatalf("invalid value returned: %s", answer)
	}
	if stdout != `Enter value for "title": [31;1mvalue must be specified[0m
Enter value for "title": ` {
		t.Fatalf("invalid stdout: %s", stdout)
	}

	// optional user input checks
	answer, stdout, err = testAskOptionWrapper("title", "", true, "should return this\n")
	if err != nil {
		t.Fatal(err)
	}
	if answer != "should return this" {
		t.Fatalf("invalid value returned: %s", answer)
	}
	if stdout != `Enter value for "title": ` {
		t.Fatalf("invalid stdout: %s", stdout)
	}

	// before \n is an empty string, should return that and does not care after the rest
	answer, stdout, err = testAskOptionWrapper("title", "", true, "\nshould not return this but empty string")
	if err != nil {
		t.Fatal(err)
	}
	if answer != "" {
		t.Fatalf("invalid value returned: %s", answer)
	}
	if stdout != `Enter value for "title": ` {
		t.Fatalf("invalid stdout: %s", stdout)
	}

	//// required selector tests

	answer, stdout, err = testAskOptionWrapper("title", "", false, "2\n", "option one", "option two")
	if err != nil {
		t.Fatal(err)
	}
	if answer != "option two" {
		t.Fatalf("invalid value returned: %s", answer)
	}
	if stdout != `Select "title" from the list:
[1] : option one
[2] : option two
Type in the option's number, then hit Enter: ` {
		t.Fatalf("invalid stdout: %s", stdout)
	}

	// higher item number than list len
	answer, stdout, err = testAskOptionWrapper("title", "", false, "5\n1\n", "option one", "option two")
	if err != nil {
		t.Fatal(err)
	}
	if answer != "option one" {
		t.Fatalf("invalid value returned: %s", answer)
	}
	if stdout != `Select "title" from the list:
[1] : option one
[2] : option two
Type in the option's number, then hit Enter: [31;1minvalid option number, pick a number 1-2[0m
Type in the option's number, then hit Enter: ` {
		t.Fatalf("invalid stdout: %s", stdout)
	}

	// char instead of number
	answer, stdout, err = testAskOptionWrapper("title", "", false, "a\n2\n", "option one", "option two")
	if err != nil {
		t.Fatal(err)
	}
	if answer != "option two" {
		t.Fatalf("invalid value returned: %s", answer)
	}
	if stdout != `Select "title" from the list:
[1] : option one
[2] : option two
Type in the option's number, then hit Enter: [31;1mfailed to parse option number, pick a number from 1-2[0m
Type in the option's number, then hit Enter: ` {
		t.Fatalf("invalid stdout: %s", stdout)
	}

	// empty instead of number
	answer, stdout, err = testAskOptionWrapper("title", "", false, "\n2\n", "option one", "option two")
	if err != nil {
		t.Fatal(err)
	}
	if answer != "option two" {
		t.Fatalf("invalid value returned: %s", answer)
	}
	if stdout != `Select "title" from the list:
[1] : option one
[2] : option two
Type in the option's number, then hit Enter: [31;1mfailed to parse option number, pick a number from 1-2[0m
Type in the option's number, then hit Enter: ` {
		t.Fatalf("invalid stdout: %s", stdout)
	}

	// auto select only element if required option
	answer, stdout, err = testAskOptionWrapper("title", "", false, "the input can be whatever", "only one option auto selected")
	if err != nil {
		t.Fatal(err)
	}
	if answer != "only one option auto selected" {
		t.Fatalf("invalid value returned: %s", answer)
	}
	if stdout != `` {
		t.Fatalf("invalid stdout: %s", stdout)
	}

	//// optional selector tests

	// DO NOT auto select only element if optional option
	answer, stdout, err = testAskOptionWrapper("title", "", true, "1\n", "only one option selected")
	if err != nil {
		t.Fatal(err)
	}
	if answer != "only one option selected" {
		t.Fatalf("invalid value returned: %s", answer)
	}
	if stdout != `Select "title" from the list:
[1] : only one option selected
[2] : <custom value>
Type in the option's number, then hit Enter: ` {
		t.Fatalf("invalid stdout: %s", stdout)
	}

	// optional option more element
	answer, stdout, err = testAskOptionWrapper("title", "", true, "2\n", "option one", "option two")
	if err != nil {
		t.Fatal(err)
	}
	if answer != "option two" {
		t.Fatalf("invalid value returned: %s", answer)
	}
	if stdout != `Select "title" from the list:
[1] : option one
[2] : option two
[3] : <custom value>
Type in the option's number, then hit Enter: ` {
		t.Fatalf("invalid stdout: %s", stdout)
	}

	// optional option more element and selected custom value
	answer, stdout, err = testAskOptionWrapper("title", "", true, "3\nthis is my custom value\n", "option one", "option two")
	if err != nil {
		t.Fatal(err)
	}
	if answer != "this is my custom value" {
		t.Fatalf("invalid value returned: %s", answer)
	}
	if stdout != `Select "title" from the list:
[1] : option one
[2] : option two
[3] : <custom value>
Type in the option's number, then hit Enter: Enter value for "title": ` {
		t.Fatalf("invalid stdout: %s", stdout)
	}
}
