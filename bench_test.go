package main

import (
	"testing"
	"strconv"
)

var word3 = []string{"дом", "меч", "кон", "сад", "вес",
"кот", "сом", "нос", "кий", "зуб"}

var word4 = []string{"мозг", "дичь", "конь", "весы", "мост",
"воск", "пила", "нора", "лиса", "авар"}

var word5 = []string{"абака", "авеню", "айран", "актин", "алтей",
"гагат", "газик", "ганаш", "гарда", "гачек"}

var word6 = []string{"рабкор", "раввин", "развой", "раздор", "размах",
"шаблон", "шагать", "шалить", "шапито", "шашист"}

var word7 = []string{"абдомен", "кабаний", "кавычки", "казимир", "эвглена",
"эйхинин", "экситон", "эластик", "элитный", "эмиссия"}

var word8 = []string{"aриаднин", "шабашник", "шабреный", "шаманить", "габардин",
"газовщик", "гайморит", "галфвинд", "гардения", "гарнитур"}

var word9 = []string{"магазинка", "майданить", "макрокосм", "маловодье", "царапанье",
"цветничок", "целенький", "цементник", "центровой", "цепляться"}

var word10 = []string{"щебетливый", "щелкнуться", "щупленький", "заарканить", "забелеться",
"заблистать", "забористый", "забрезжить", "забубенный", "заваливать"}

var word11 = []string{"садоводство", "салютование", "самогонщица", "дактилозоид", "дауэсизация",
"двоеженство", "двусемянный", "двухголовый", "двухсветный", "двушерстный"}

var word12 = []string{"юрисконсульт", "юмористичный", "юнгштурмовка", "чадорождение", "человеческий",
"червеязычные", "чередоваться", "черепитчатый", "черноголовка", "чернорабочая"}

var word13 = []string{"фабзавкомовец", "фальсификация", "фармакогнозия", "фельетонистка", "ферровольфрам",
"фешенебельный", "заатмосферный", "забинтоваться", "забраковывать", "завербовывать"}

var word14 = []string{"характеристика", "хлеботоргующий", "хлорпикриновый", "холоднокровный", "хормейстерский",
"хромофототипия", "хронологизация", "художественный", "евангелический", "единомышленный"}

var word15 = []string{"галантерейность", "гальванотропизм", "гастролирование", "генеалогический", "германофильский",
"гидроаппаратура", "языкотворческий", "языкотворчество", "яйцеобразование", "дазиметрический"}


func BenchmarkMain(b *testing.B) {
	dict := GetDict(DictPath)
	var words = [][]string {
		word3, word4, word5, word6, word7, word8,
		word9, word10, word11, word12, word13, word14, word15,
	}

	for count, items := range words{
		b.Run("Базовый " + strconv.Itoa(count + 3), func(b *testing.B) {
			benchmarkBase(items, count + 3, dict, b)
		})
		b.Run("Базовый с сортировкой входного слова " + strconv.Itoa(count + 3), func(b *testing.B) {
			benchmarkBaseWithWordSort(items, count + 3, dict, b)
		})
		b.Run("Базовый с сортировкой словар" + strconv.Itoa(count + 3), func(b *testing.B) {
			benchmarkRunWithDictSort(items, count + 3, dict, b)
		})
		b.Run("Базовый с сортировкой словаря и входного слова" + strconv.Itoa(count + 3), func(b *testing.B) {
			benchmarkRunWithDictSortWithWordSort(items, count + 3, dict, b)
		})
	}
}

func benchmarkBase(values []string, len int, dictionary map[int]map[string][]string, b *testing.B) {
	numberVal := DecomposeNumber(len)
	b.ResetTimer()

	for _, value := range values {
		for n := 0; n < b.N; n++ {
			badWords, str := PrepareString(value)
			value := Run(str, numberVal, dictionary)


			var items  [][]string

			for _, child := range value {
				items = collapse(child, []string{}, items)
			}

			combinations := GetWords(dictionary, items)

			items = [][]string{}

			for _,combination := range combinations {
				items = GetAllCombinations(combination, []string{}, 0, items, badWords)
			}

			DeleteEquals(items)
		}
	}

}

func benchmarkBaseWithWordSort(values []string, len int, dictionary map[int]map[string][]string, b *testing.B) {
	numberVal := DecomposeNumber(len)
	b.ResetTimer()

	for _, value := range values {
		for n := 0; n < b.N; n++ {
			value = sortStringBytes(value)
			badWords, str := PrepareString(value)
			value := Run(str, numberVal, dictionary)


			var items  [][]string

			for _, child := range value {
				items = collapse(child, []string{}, items)
			}

			combinations := GetWords(dictionary, items)

			items = [][]string{}

			for _,combination := range combinations {
				items = GetAllCombinations(combination, []string{}, 0, items, badWords)
			}

			DeleteEquals(items)
		}
	}

}

func benchmarkRunWithDictSort(values []string, len int, dictionary map[int]map[string][]string, b *testing.B) {
	numberVal := DecomposeNumber(len)
	b.ResetTimer()

	for _, value := range values {
		for n := 0; n < b.N; n++ {
			badWords, str := PrepareString(value)
			value := RunWithSort(str, numberVal, dictionary)


			var items  [][]string

			for _, child := range value {
				items = collapse(child, []string{}, items)
			}

			combinations := GetWords(dictionary, items)

			items = [][]string{}

			for _,combination := range combinations {
				items = GetAllCombinations(combination, []string{}, 0, items, badWords)
			}

			DeleteEquals(items)
		}
	}

}

func benchmarkRunWithDictSortWithWordSort(values []string, len int, dictionary map[int]map[string][]string, b *testing.B) {
	numberVal := DecomposeNumber(len)
	b.ResetTimer()

	for _, value := range values {
		for n := 0; n < b.N; n++ {
			value = sortStringBytes(value)

			badWords, str := PrepareString(value)
			value := RunWithSort(str, numberVal, dictionary)


			var items  [][]string

			for _, child := range value {
				items = collapse(child, []string{}, items)
			}

			combinations := GetWords(dictionary, items)

			items = [][]string{}

			for _,combination := range combinations {
				items = GetAllCombinations(combination, []string{}, 0, items, badWords)
			}

			DeleteEquals(items)
		}
	}

}


