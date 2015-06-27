package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func AskForStringFromReader(messageToPrint string, inputReader io.Reader) (string, error) {
	scanner := bufio.NewScanner(inputReader)
	fmt.Printf("%s : ", messageToPrint)
	if scanner.Scan() {
		scannedText := scanner.Text()
		return scannedText, nil
	}
	return "", errors.New("Failed to get input - scanner failed.")
}

func AskForString(messageToPrint string) (string, error) {
	return AskForStringFromReader(messageToPrint, os.Stdin)
}

func AskForIntFromReader(messageToPrint string, inputReader io.Reader) (int64, error) {
	userInputStr, err := AskForStringFromReader(messageToPrint, inputReader)
	if err != nil {
		return 0, err
	}
	return strconv.ParseInt(userInputStr, 10, 64)
}

func AskForInt(messageToPrint string) (int64, error) {
	return AskForIntFromReader(messageToPrint, os.Stdin)
}

func AskForBoolFromReader(messageToPrint string, inputReader io.Reader) (bool, error) {
	userInputStr, err := AskForStringFromReader(messageToPrint, inputReader)
	if err != nil {
		return false, err
	}
	lowercased := strings.ToLower(userInputStr)
	if lowercased == "yes" || lowercased == "y" {
		return true, nil
	}
	if lowercased == "no" || lowercased == "n" {
		return false, nil
	}
	return strconv.ParseBool(lowercased)
}

func AskForBool(messageToPrint string) (bool, error) {
	return AskForBoolFromReader(messageToPrint, os.Stdin)
}

func main() {
	retStr, err := AskForString("Please enter some text here")
	if err != nil {
		log.Fatalln("Error:", err)
	}
	fmt.Println("Entered text was:", retStr)

	retInt, err := AskForInt("Please enter a number")
	if err != nil {
		log.Fatalln("Error:", err)
	}
	fmt.Println("Entered:", retInt)

	retBool, err := AskForBool("Yes or no?")
	if err != nil {
		log.Fatalln("Error:", err)
	}
	fmt.Println("Entered:", retBool)
}
