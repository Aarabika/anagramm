package main

import (
	"io/ioutil"
	"bytes"
	"os"
)

func readFile(filePath string) ([][]byte, error) {
	file, err := ioutil.ReadFile(filePath)


	if err == nil {
		return bytes.Split(file, []byte{13, 10}), nil
	}

	return nil, err
}

func compareDictionaries(dictionary1, dictionary2[][]byte) [][]byte {
	var diffWords  [][]byte

	for count1 := 0; count1 < len(dictionary1); count1++ {
		notExist := true
		for count2 := 0; count2 < len(dictionary2); count2++ {
			if bytes.Equal(dictionary1[count1], dictionary2[count2]) {
				notExist = false
				break
			}
		}

		if notExist {
			diffWords = append(diffWords, dictionary1[count1])
		}
	}

	return diffWords
}

func compareWords(word1 string, word2 []byte , length int, words [][]byte) {
	var need bool

	for index1 := 0; index1 < length; index1++ {
		need = false

		for index2 := 0; index2 < length; index2++ {
			if word1[index1] == word2[index2] {
				need = true
				break
			}
		}

		if !need {
			break
		}
	}

	if need {
		words = append(words, word2)
	}
}

func findMaxLenWord(dictionary [][]byte) int {
	length := len(dictionary)

	maxLen := 0

	for index := 0; index < length; index++ {
		if len(dictionary[index]) > maxLen {
			maxLen = len(dictionary[index])
		}
	}

	return maxLen
}

func findMinLenWord(dictionary [][]byte) int {
	length := len(dictionary)

	minLen := 9999

	for index := 0; index < length; index++ {
		if len(dictionary[index]) < minLen && len(dictionary[index]) != 0  {
			minLen = len(dictionary[index])
		}
	}

	return minLen
}

func deleteOneSymbolWordFromDictionary(dictionary [][]byte , newDictFileName string) {
	file, _ := os.OpenFile(newDictFileName, os.O_APPEND|os.O_WRONLY, 0600)

	for count := 0; count < len(dictionary); count++ {
		if len(dictionary[count]) > 1 {
			file.WriteString(string(dictionary[count]) + string([]byte{13, 10}))
		}

	}

	file.Close()
}

func GetDict(dictName string) map[int]map[string][]string  {
	dictionary, _ := readFile(dictName)
	dictionaryLen := len(dictionary)

	dictionaryWords := make(map[int]map[string][]string, maxLen - minLen + 1)

	for count := 0; count < dictionaryLen; count++ {
		currentWord := string(dictionary[count])
		currentLen := len(currentWord)
		rCurrentWord := sortStringBytes(currentWord)

		itemsByNumber, ok := dictionaryWords[currentLen]

		if !ok {
			itemsByNumber = make(map[string][]string)
		}

		itemsByWords, _ := dictionaryWords[currentLen][rCurrentWord]

		itemsByNumber[rCurrentWord] = append(itemsByWords, currentWord)

		dictionaryWords[currentLen] = itemsByNumber
	}

	return dictionaryWords
}