package main

import "fmt"

func callbackHelp() error {
	fmt.Println("Welcome to the Pokedex Help Menu")
	fmt.Println("Please select one of the following help options: ")
	availableCommands := getCommands()

	for _, command := range availableCommands {
		fmt.Printf("- %s: %s\n", command.name, command.description)
	}

	fmt.Println("")
	return nil
}
