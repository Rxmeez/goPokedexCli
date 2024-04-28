package repl

import "fmt"

func commandHelp() error {
	fmt.Println("\nWelcome to the Pokedex!")
	fmt.Print("Usage:\n\n")
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}
