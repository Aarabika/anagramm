package main

import (
	"sort"
)

const space  = 32

func PrepareString(value string) ([]string, string)  {
	value +=" "
	stringLen := len(value)

	var result []string

	var word string
	var str string

	for count := 0; count < stringLen; count++ {
		if value[count] != space {
			word = string(append([]byte(word), value[count]))
		} else {
			if len(word) >= minLen {
				if !findString(result, word) {
					result = append(result, word)
				}
			}
			str += word
			word = ""
		}

	}


	return result, str
}

func delSubstr(str string, substr string) string {

	if str == substr {
		return ""
	}

	stringLen := len(str)

	var resultStr []byte


	for strIndex := 0; strIndex < stringLen; strIndex++ {
		charIndex := -1

		for substrIndex := 0 ; substrIndex < len(substr); substrIndex++ {
			if substr[substrIndex] == str[strIndex] {
				charIndex = substrIndex
				break
			}
		}

		if charIndex != -1 {
			if charIndex == 0 {
				substr = substr[1:]
			} else if charIndex == len(substr) - 1{
				substr = substr[0:charIndex]
			} else {
				substr = substr[0:charIndex] + substr[charIndex + 1:]
			}
		} else {
			resultStr = append(resultStr, str[strIndex])
		}
	}

	if len(substr) == 0 {
		return string(resultStr)
	} else {
		return  str
	}
}

func findString(all []string, str string) bool  {
	return sort.Search(len(all), func(i int) bool {
		if len(all[i]) == len(str) {
			return str == all[i]
		}
		return false
	}) < len(all)
}

func findArrayString(all [][]string, strings []string) bool  {

	for i, _ := range all {
		if len(all[i]) == len(strings){
			isset := true

			for key := 0 ; key < len(strings); key++ {
				if strings[key] != all[i][key] {
					isset = false
					break
				}
			}

			if isset {
				return true
			}
		}
	}
	return false
}


