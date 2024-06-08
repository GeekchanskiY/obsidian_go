package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println("Obsidian Vault Reader and parser")
	if len(args) < 2 {
		fmt.Println("No vault path provided")
		fmt.Println("Usage: vault_reader <vault_path>")
		return
	}
	fmt.Println(args[1])
}
