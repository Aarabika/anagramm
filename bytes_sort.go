package main

import "sort"

type BytesSort struct {
	Items []byte
	Length int
}

func (b BytesSort) Len() int  {
	return b.Length
}

func (b BytesSort) Less(i, j int) bool  {
	return b.Items[i] < b.Items[j]
}

func (b BytesSort) Swap(i, j int)  {
	container := b.Items[i]
	b.Items[i] = b.Items[j]
	b.Items[j] = container
}

func sortStringBytes(val string) string  {
	value := []byte(val)
	item := BytesSort{
		Length: len(value),
		Items: value,
	}
	sort.Sort(item)

	return string(value)
}

func sortBytesArray(val []byte)   {
	item := BytesSort{
		Length: len(val),
		Items: val,
	}
	sort.Sort(item)
}
