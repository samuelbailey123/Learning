package cmd

import "fmt"

var ToDo []string

func List() {
	if len(ToDo) == 0 {
		fmt.Println("No items to list")
	} else {
		fmt.Println("The list contains:", ToDo)
	}
}
