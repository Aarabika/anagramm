package main

import "sort"

type StringSort struct {
	Items []string
	Length int
}

func (b StringSort) Len() int  {
	return b.Length
}

func (b StringSort) Less(i, j int) bool  {
	for key := range b.Items[i]{
		if b.Items[i][key] < b.Items[j][key] {
			return true
		} else if b.Items[i][key] > b.Items[j][key] {
			return false
		}
	}
	panic(b)
}

func (b StringSort) Swap(i, j int)  {
	container := b.Items[i]
	b.Items[i] = b.Items[j]
	b.Items[j] = container
}


func equalLenStringSort(val []string)   {
	item := StringSort{
		Length: len(val),
		Items: val,
	}
	sort.Sort(item)
}
