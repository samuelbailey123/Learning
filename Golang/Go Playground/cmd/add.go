package cmd

import "fmt"

func Add(addToList string) {
	if Contains(ToDo, addToList) {
		fmt.Println("Item already exists")
	} else {
		ToDo = append(ToDo, addToList)
		fmt.Println("Item added")
	}
}
