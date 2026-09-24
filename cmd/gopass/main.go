package main

import (
	"fmt"
	"os"

	"github.com/PoliakovAnton/gopass/internal/generator"
)

func main() {
	password, err := generator.Generate(16)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	fmt.Println("Generated password:", password)
}
