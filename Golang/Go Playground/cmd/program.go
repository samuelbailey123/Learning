package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

var input string

func Program() {
	for input != "4" {
		printMenu()
		menu()
	}
}

func printMenu() {
	fmt.Printf("\n")
	fmt.Println("Main menu")
	fmt.Println("1. Add")
	fmt.Println("2. List")
	fmt.Println("3. Mark as complete")
	fmt.Println("4. Exit")
}

func menu() {
	fmt.Print("Enter text: ")
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return
	}

	input = strings.TrimSuffix(input, "\n")
	fmt.Println(input)
	fmt.Println()

	switch input {
	case "1":
		fmt.Println("What would you like to add to the list?")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occured while reading input. Please try again", err)
			return
		}
		input = strings.TrimSuffix(input, "\n")
		Add(input)
	case "2":
		List()
	case "3":
		fmt.Println("What item would you like to mark as complete?")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occured while reading input. Please try again", err)
			return
		}
		input = strings.TrimSuffix(input, "\n")
		if Contains(ToDo, input) {
			fmt.Println("Marking item as complete")
			MarkAsComplete(input)
		} else {
			fmt.Println("Item not found")
		}
	case "4":
		fmt.Println("Exiting...")
		time.Sleep(time.Second * 1)
		os.Exit(0)
	default:
		fmt.Println("Invalid input. Please try again")
	}
}

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
