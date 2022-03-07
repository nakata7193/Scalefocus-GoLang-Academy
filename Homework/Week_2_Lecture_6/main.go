package main

import "fmt"

type Item struct {
	value    int
	prevItem *Item
}

type MagicList struct {
	lastItem *Item
}

func add(list *MagicList, value int) {
	i := Item{
		value: value,
	}

	if list.lastItem == nil {
		list.lastItem = &i
	} else {
		i.prevItem = list.lastItem
		list.lastItem = &i
	}
}

func toSlice(list *MagicList) []int {
	item := list.lastItem
	var itemList []int
	for true {
		if item != nil {
			itemList = append(itemList, item.value)
			item = item.prevItem
		} else {
			break
		}
	}
	return itemList
}

func main() {
	magicList := &MagicList{}
	add(magicList, 12)
	add(magicList, 44)
	add(magicList, 33)
	add(magicList, 14)
	add(magicList, 56)
	add(magicList, 78)
	add(magicList, 43)
	add(magicList, 15)
	add(magicList, 61)
	add(magicList, 71)
	fmt.Print(toSlice(magicList))
}
