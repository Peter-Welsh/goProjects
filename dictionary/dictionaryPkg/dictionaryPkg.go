package dictionarypkg

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type WordDefs []struct {
	Word     string `json:"word"`
	Meanings []struct {
		PartOfSpeech string `json:"partOfSpeech"`
		Definitions  []struct {
			Definition string `json:"definition"`
		} `json:"definitions"`
	} `json:"meanings"`
}

const baseUrl = "https://api.dictionaryapi.dev/api/v2/entries/en/"

func DefineWord(word string) {
	resp, err1 := http.Get(baseUrl + word)
	if err1 != nil {
		panic(err1)
	}
	readResponse(resp, word)
}

func readResponse(resp *http.Response, word string) {
	defer resp.Body.Close()
	if resp.StatusCode == 404 {
		fmt.Printf("No definitions found for '%v'.", word)
		return
	}
	if resp.StatusCode != 200 {
		fmt.Println("The server isn't responding. Try again later.")
		return
	}
	var wordDefs WordDefs
	err := json.NewDecoder(resp.Body).Decode(&wordDefs)
	if err != nil {
		panic(err)
	}
	for i := range wordDefs {
		printDefinitions(wordDefs, i)
	}
}

func printDefinitions(wordDefs WordDefs, i int) {
	for _, meaning := range wordDefs[i].Meanings {
		fmt.Println("\r\nPart of speech:", meaning.PartOfSpeech)
		fmt.Println("Definition(s):")
		for j, definition := range meaning.Definitions {
			fmt.Printf("%v.", j+1)
			fmt.Println("", definition.Definition)
		}
	}
}
