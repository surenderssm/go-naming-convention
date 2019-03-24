package repository

import "testing"

// the the word repo
func getTestWords() []string {

	tokens := []string{"one", "two", "three", "hello", "world"}
	return tokens
}

func TestNewWordRepository(t *testing.T) {
	words := getTestWords()
	wordRepo := NewWordRepository(words, nil)

	if _, ok := wordRepo.WordMap[words[0]]; !ok {

		t.Error("Wordmap is not initialized in wordRepo")
	}
}

func TestIsValidWordInWordRepository(t *testing.T) {
	words := getTestWords()
	wordRepo := NewWordRepository(words, nil)

	for _, value := range words {
		if ok, _ := wordRepo.IsValidWord(value); !ok {
			t.Error("ISValidWord Failed for expected word")
		}
	}
}
