package repl

import (
	"fmt"
	"os"
)

func commandExit() error {
	exitMessage := "Exiting the Pokedex"
	fmt.Println(exitMessage)
	os.Exit(0)
	return nil
}
