package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/PoliakovAnton/gopass/internal/generator"
)

func main() {
	length := flag.Int("length", 16, "password length")
	flag.Parse()

	password, err := generator.Generate(*length)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	fmt.Println(password)
}
