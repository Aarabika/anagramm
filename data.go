package main

type anagram struct {
	parent *anagram

	value string

	mask string

	children []*anagram
}

func build(mask []int, dictKeys map[int][]string, item *anagram, currentMaskIndex int, currentDictIndex int)  {

	if currentMaskIndex == len(mask) {
		return
	}

	for key, word := range dictKeys[mask[currentMaskIndex]][currentDictIndex:] {

		newValue := delSubstr(item.value, word)

		if newValue == item.value {
			continue
		}

		child := new(anagram)
		child.value = newValue
		child.parent = item
		child.mask = word
		child.children = []*anagram {}


		if currentMaskIndex + 1 < len(mask) && mask[currentMaskIndex] == mask[currentMaskIndex + 1]{
			build(mask, dictKeys, child, currentMaskIndex + 1, key + currentDictIndex)
		} else {
			build(mask, dictKeys, child, currentMaskIndex + 1, 0)
		}


		if child.value == "" || len(child.children) != 0 {
			item.children = append(item.children, child)
		}
	}

	return
}

func Run(value string, number [][]int, dictionary map[int]map[string][]string) []*anagram  {

	var values []*anagram

	for _, mask := range number {
		item := new(anagram)
		item.parent = nil
		item.value = value
		item.children = []*anagram{}

		dictKeys := make(map[int][]string)

		var items []int

		for item, numbersCount := range mask {
			if numbersCount != 0 {
				current := item + minLen
				dictKeys[current] = []string{}

				for key, _ := range dictionary[current]  {
					newValue := delSubstr(value, key)

					if newValue != value {
						dictKeys[current] = append(dictKeys[current], key)
					}

				}
				for numbersCount > 0 {
					items = append(items, current)
					numbersCount --
				}
			}
		}


		build(items, dictKeys, item, 0,0)

		if len(item.children) != 0 {
			values = append(values, item)
		}
	}


	return values
}

func RunWithSort(value string, number [][]int, dictionary map[int]map[string][]string) []*anagram  {

	var values []*anagram

	for _, mask := range number {
		item := new(anagram)
		item.parent = nil
		item.value = value
		item.children = []*anagram{}

		dictKeys := make(map[int][]string)

		var items []int

		for item, numbersCount := range mask {
			if numbersCount != 0 {
				current := item + minLen
				dictKeys[current] = []string{}

				for key, _ := range dictionary[current]  {
					newValue := delSubstr(value, key)

					if newValue != value {
						dictKeys[current] = append(dictKeys[current], key)
					}

				}

				equalLenStringSort(dictKeys[current])

				for numbersCount > 0 {
					items = append(items, current)
					numbersCount --
				}
			}
		}

		build(items, dictKeys, item, 0,0)

		if len(item.children) != 0 {
			values = append(values, item)
		}
	}


	return values
}