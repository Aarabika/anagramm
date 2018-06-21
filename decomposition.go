package main

const minLen  = 2
const maxLen  = 32


func DecomposeNumber(value int) [][]int {
	values := make(map[int][][]int, value - minLen + 1)

	for valueCurrent := minLen; valueCurrent <= value; valueCurrent++ {
		var items [][]int

		maxCurrentLen := maxLen

		if	maxLen > valueCurrent {
			maxCurrentLen = valueCurrent
		}

		for count := minLen; count <= maxCurrentLen; count++ {
			dif := valueCurrent - count

			if dif == 0 && count == valueCurrent {
				var newItem = make([]int, maxLen - minLen + 1)
				newItem[count - minLen] = 1
				items = append(items, newItem)
			}

			if dif >= minLen {

				for currentItems := 0; currentItems < len(values[dif]); currentItems++ {
					var newItem = make([]int, maxLen - minLen + 1)

					var isIncorrect = false

					for currentCoefficient := count - minLen + 1; currentCoefficient < maxLen - 1; currentCoefficient++ {
						if values[dif][currentItems][currentCoefficient] >= 1 {
							isIncorrect = true
							break
						}
					}

					if isIncorrect {
						break
					}

					copy(newItem, values[dif][currentItems])

					newItem[count - minLen]++

					items = append(items, newItem)
				}
			}
		}

		values[valueCurrent] = items
	}

	return values[value]
}

//func decomposeString(string []byte) map[byte]int {
//	value := make(map[byte]int, len(string))
//
//	for index := 0; index < len(string); index ++ {
//		item, _ := value[string[index]]
//		value[string[index]] = item + 1
//	}
//
//
//	return value
//}

