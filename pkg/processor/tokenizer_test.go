package processor

import (
	"go-naming-convention/pkg/repository"
	"testing"
)

// the the word repo
func getTestWords() []string {

	tokens := []string{"one", "two", "three", "hello", "world"}
	return tokens
}

func TestForValidToken(t *testing.T) {

	if ok := IsValidToken("a"); ok {
		t.Error("token should be atelast 3")
	}
	if ok := IsValidToken("one"); !ok {
		t.Error("token should be atelast 3")
	}

	if ok := IsValidToken("one1"); ok {
		t.Error("token can not have a digit")
	}

	if ok := IsValidToken("one@#"); ok {
		t.Error("token can not have any special character")
	}
}

func getTokenProcessor() *Tokenizer {
	words := getTestWords()
	wordRepo := repository.NewWordRepository(words, nil)
	tokenProcessor := NewTokenizer(wordRepo)
	return tokenProcessor
}

func TestForTokenProcess(t *testing.T) {
	tokenProcessor := getTokenProcessor()

	text := "onetwothree"
	result := tokenProcessor.ProcessToken(text, CamelCase)

	if result != "oneTwoThree" {
		t.Error("Expected output not returned")
	}

	result = tokenProcessor.ProcessToken(text, PascalCase)

	if result != "OneTwoThree" {
		t.Error("Expected output not returned")
	}

	// TODO : check if DP is used
	text = "onetwothreeonetwothree"
	result = tokenProcessor.ProcessToken(text, CamelCase)

	if result != "oneTwoThreeOneTwoThree" {
		t.Error("Expected output not returned")
	}
}

func TestForTokenProcessWithInvalidWord(t *testing.T) {
	tokenProcessor := getTokenProcessor()

	// not a valid word
	text := "ewbbseq"
	result := tokenProcessor.ProcessToken(text, CamelCase)

	if result != "" {
		t.Error("no outpput was expected")
	}

	// "tiger" is not a valid word in our test so output should be returned
	text = "tigerworld"

	result = tokenProcessor.ProcessToken(text, CamelCase)

	if result != "" {
		t.Error("Expected output not returned")
	}
}
