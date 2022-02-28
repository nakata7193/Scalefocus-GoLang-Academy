package main

import "fmt"

type Item struct {
	Value    int
	PrevItem *Item
}

type MagicList struct {
	LastItem *Item
}

func add(list *MagicList, value int) {
	i := Item{
		Value: value,
	}

	if list.LastItem == nil {
		list.LastItem = &i
		i.PrevItem = nil
	} else {
		i.PrevItem = list.LastItem
		list.LastItem = &i
	}
}

// func toSlice(list *MagicList) []int {
// 	var numbers []int
// 	for 
// 	return append(numbers, list.LastItem.Value)
// }

func main() {
	magicList := &MagicList{}
	add(magicList, 12)
	add(magicList, 44)
	add(magicList, 33)
	fmt.Println(toSlice(magicList))
}
