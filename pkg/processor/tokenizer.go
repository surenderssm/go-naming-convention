package processor

import (
	"go-naming-convention/pkg/repository"
	"strings"
)

// Tokenizer .
type Tokenizer struct {
	WordRepo *repository.WordReposiotry
}

type tokenTask struct {
	validMemoize   map[string][]string
	invalidMemoize map[string]bool
	// reference of the word repo embedded during the init
	wordRepo *repository.WordReposiotry
}

// NewTokenizer .
func NewTokenizer(wordRepo *repository.WordReposiotry) *Tokenizer {
	tokenizer := new(Tokenizer)
	tokenizer.WordRepo = wordRepo
	return tokenizer
}

func newTokenTask(wordRepo *repository.WordReposiotry) *tokenTask {
	task := new(tokenTask)
	task.invalidMemoize = make(map[string]bool)
	task.validMemoize = make(map[string][]string)
	task.wordRepo = wordRepo
	return task
}

// ProcessToken extract valid words from the given text and return the token after concatenating them in given format
func (tokenizer *Tokenizer) ProcessToken(text string, caseType NamingConventionCase) string {

	currentTask := newTokenTask(tokenizer.WordRepo)
	data := currentTask.process(text)
	output := Format(data, caseType)
	return output
}

func (task *tokenTask) process(text string) []string {
	tokens := make([]string, 0)

	// DP - check for memoize, to avoid duplicate compute
	if value, ok := task.validMemoize[text]; ok {
		return value
	}

	// allready computed not a valid word
	if _, ok := task.invalidMemoize[text]; ok {
		return tokens
	}

	lengthOfText := len(text)

	// avoid compute for text length then 3
	if lengthOfText < 3 {
		return tokens
	}

	// check if the text is valid word
	if ok, _ := task.wordRepo.IsValidWord(text); ok {
		tokens = append(tokens, text)
		return tokens
	}

	for index := 3; (lengthOfText - index) >= 3; index++ {
		text1 := text[0:index]
		text2 := text[index:lengthOfText]
		if len(text2) < 3 {
			continue
		}

		if _, ok := task.invalidMemoize[text1]; ok {
			continue
		}

		if _, ok := task.invalidMemoize[text2]; ok {
			continue
		}
		// process left node
		result1 := task.process(text1)

		// no need to process right node as left node is not valid
		// prune right side to save compute
		if len(result1) == 0 {
			continue
		}

		// process right node
		result2 := task.process(text2)

		if (len(result1) > 0) && (len(result2) > 0) {
			result1 := append(result1, result2...)
			task.validMemoize[text] = result1
			return result1
		}
	}
	task.invalidMemoize[text] = true
	return tokens
}

// IsValidToken check if the token is valid or not
// 1. Length should be atelast 3
// 2. As of only english token are supported
// 3. Any digit, special character or space will make the token invalid
func IsValidToken(token string) bool {

	if len(token) < 3 {
		return false
	}

	item := strings.ToLower(token)
	for _, alphabet := range item {
		if alphabet < 'a' || alphabet > 'z' {
			return false
		}
	}
	return true
}

//GetWords  N^2 //
 // brute way to get text  -> should not be used
// func bruteWayToGetTokens(text string) []string {

// 	items := make([]string, 0)

// 	end := len(text)

// 	for end > 0 {
// 		for i := 0; i <= end; i++ {
// 			substring := text[i:end]
// 			if ok := IsWord(substring); ok {
// 				items = append(items, substring)
// 				end = i
// 				i = 0
// 			}
// 		}
// 	}
// 	return items
// }
