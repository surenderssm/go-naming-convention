package repository

import (
	"go-naming-convention/pkg/common"
	"strconv"
	"strings"
)

// WordReposiotry repo to manage words
type WordReposiotry struct {
	// cache of word
	WordMap map[string]bool
	// TODO : pool of provider and round robin/ based on availabilty can be used to improve the system
	WordProvider *OxfordWordProvider
}

// NewWordRepository returns a new word repo
// repo can be initialized with basic words as well
func NewWordRepository(words []string, wordProvider *OxfordWordProvider) *WordReposiotry {

	wordRepo := new(WordReposiotry)
	wordMap := make(map[string]bool)
	if len(words) > 0 {
		for _, value := range words {
			wordMap[strings.ToLower(value)] = true
		}
	}
	common.Logger.Info("WordRepoInitialized-Length-" + strconv.Itoa(len(wordMap)))
	wordRepo.WordMap = wordMap
	wordRepo.WordProvider = wordProvider
	return wordRepo
}

// IsValidWord returns true if given word a valid word
func (wordRepo *WordReposiotry) IsValidWord(text string) (bool, error) {

	// check word via WordProvider
	// TODO : first priorty should be given to local cache, but as it demands lets go with word provider

	// go to cache mode if word provider is not present
	// this will be nil unit test
	if wordRepo.WordProvider == nil {
		if _, ok := wordRepo.WordMap[text]; ok {
			return true, nil
		}
		return false, nil
	}

	exist, error := wordRepo.WordProvider.WordExist(text)

	// if there are error (not serviceable) fallback to local cache
	if error != nil {

		common.Logger.Error("WordProvider-Failed" + error.Error())

		if _, ok := wordRepo.WordMap[text]; ok {
			return true, nil
		}
		return false, nil
	}

	if exist {

		// update the cache if valid word not in the cache
		if _, ok := wordRepo.WordMap[text]; !ok {
			wordRepo.WordMap[text] = true
		}
		return true, nil
	}
	// cache of not valid words can be expensive, hence decided not to store
	// In future it can be provisioned as it will offload the serviceProvider leading to save cost
	return false, nil
}

// func poupulateCache() {
// 	file, err := os.Open("words.txt")
// 	defer file.Close()

// 	if err != nil {
// 		fmt.Print(err)
// 		return
// 	}

// 	// Start reading from the file with a reader.
// 	reader := bufio.NewReader(file)

// 	var line string
// 	for {

// 		line, err = reader.ReadString('\n')
// 		line = strings.Split(line, "\n")[0]
// 		line = strings.ToLower(line)

// 		dictionary[line] = true
// 		if err != nil {
// 			break
// 		}
// 	}

// 	if err != io.EOF {
// 		fmt.Printf(" > Failed!: %v\n", err)
// 	}

// 	return
// }
