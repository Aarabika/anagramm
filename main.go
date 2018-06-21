package main

import (
	_ "github.com/wcharczuk/go-chart"
	"os"
	"strconv"
	"golang.org/x/text/encoding/charmap"
	"fmt"
)

const DictPath  = "./dict/dict1.txt"

func main() {
	dict := GetDict(DictPath)

	fmt.Print("Введите строку: ")

	rusWord, _ := charmap.Windows1251.NewEncoder().String("кот кот кот кот")

	words, word := PrepareString(rusWord)
	numberVal := DecomposeNumber(len(word))
	GetAnagrams(word, words, numberVal, dict)


	//fmt.Print("Введите строку: ")
	//scanner := bufio.NewScanner(os.Stdin)
	//scanner.Scan()
	//str := scanner.Text()
	//
	//rusWord, _ := charmap.Windows1251.NewEncoder().String(str)
	//
	//words, word := PrepareString(rusWord)
	//numberVal := DecomposeNumber(len(word))
	//combination := GetAnagrams(word, words, numberVal, dict)
	//
	//for k, v := range combination {
	//	for key, word := range v {
	//		str, _ := charmap.Windows1251.NewDecoder().String(word)
	//		v[key] = string(str)
	//	}
	//
	//	fmt.Print(strconv.Itoa(k)  + ": " + strings.Join(v," "))
	//}
}

func GetAnagrams(str string, badWords []string, number [][]int, dictionary map[int]map[string][]string) [][]string {

	value := RunWithSort(str, number, dictionary)


	var items  [][]string

	for _, child := range value {
		items = collapse(child, []string{}, items)
	}

	combinations := GetWords(dictionary, items)

	items = [][]string{}
	for _,combination := range combinations {
		items = GetAllCombinations(combination, []string{}, 0, items, badWords)
	}

	return DeleteEquals(items)
}

func collapse(item *anagram , current []string, all [][]string) [][]string {
	for _, child := range item.children {
		var values = append(current, child.mask)
		all = collapse(child, values, all)
	}

	if len(item.children) == 0 {
		return append(all, current)
	}

	return all
}

func GetWords(dictionary map[int]map[string][]string, combinations [][]string) [][][]string {
	words := make([][][]string, len(combinations))

	for key1, combination := range combinations {
		words[key1] = make([][]string, len(combination))

		for key2, word  := range combination {
			words[key1][key2] = dictionary[len(word)][word]
		}
	}

	return words
}

func GetAllCombinations(combinations [][]string, current []string, currentIndex int, result [][]string, badWords []string) [][]string {
	if currentIndex == len(combinations) {
		return append(result, current)
	}

	for _, val := range combinations[currentIndex] {
		if  !findString(badWords, val) {
			result = GetAllCombinations(combinations, append(current, val), currentIndex + 1, result, badWords)
		}
	}

	return result
}

func DeleteEquals(combinations [][]string) [][]string {
	var result [][]string

	for k := range combinations {
		if !findArrayString(result, combinations[k]) {
			result = append(result, combinations[k])
		}
	}

	return result
}

func getCombinations(dictionary map[int]map[string][]string, combinations [][]string) []map[string][]string {
	words := make([]map[string][]string, len(combinations))

	for key, combination := range combinations {
		words[key] = make(map[string][]string)

		for _, word  := range combination {
			words[key][word] = dictionary[len(word)][word]
		}
	}

	return words
}



func output(items []*anagram , file *os.File) {
	for key, val := range items {
		file.WriteString(strconv.Itoa(key) + " : " + val.mask + string([]byte{13, 10}))
		if len(val.children) != 0 {
			output(val.children, file)
		}
	}

}

func outputStrings(items [][]string , file *os.File) {
	for key, val := range items {
		file.WriteString(strconv.Itoa(key) + " : " + string([]byte{13, 10}))

		for _, word := range val {
			file.WriteString(word + ", ")
		}

		file.WriteString(string([]byte{13, 10}))
	}

}

func outputMap(items [][][]string , file *os.File) {
	for key, val := range items {
		file.WriteString(strconv.Itoa(key) + " : " + string([]byte{13, 10}))

		for _, words := range val {
			file.WriteString("[ ")
			for _, word := range words {
				file.WriteString(word + ", ")
			}
			file.WriteString("] ")
		}

		file.WriteString(string([]byte{13, 10}))
	}

}
func length(items []*anagram, count int) int {
	for _, val := range items {
		if len(val.children) != 0 {
			count = length(val.children, count + 1)
		}
	}
	return count
}

func collect(item *anagram, strings [][]string , currentWords[]string)  {

	if item.mask != "" {
		currentWords = append(currentWords, item.mask)
	}

	if len(item.children) != 0 {
		words := make([]string, len(currentWords))
		copy(words, currentWords)
		for _, val := range item.children {
			collect(val, strings, words)
		}
	} else {
		items := append(strings, currentWords)
		copy(strings, items)
	}
}