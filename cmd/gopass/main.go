package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/PoliakovAnton/gopass/internal/generator"
	"github.com/PoliakovAnton/gopass/internal/interactive"
	"github.com/PoliakovAnton/gopass/internal/passphrase"
	"github.com/PoliakovAnton/gopass/internal/strength"
)

func main() {
	length := flag.Int("length", 16, "password length")
	useLowercase := flag.Bool("lowercase", true, "include lowercase letters")
	useUppercase := flag.Bool("uppercase", false, "include uppercase letters")
	useNumbers := flag.Bool("numbers", false, "include numbers")
	useSymbols := flag.Bool("symbols", false, "include symbols")

	usePassphrase := flag.Bool("passphrase", false, "generate a passphrase")
	wordsCount := flag.Int("words", 4, "number of words in passphrase")

	strengthPassword := flag.String("strength", "", "analyze password strength")

	interactiveMode := flag.Bool("interactive", false, "interactive mode")

	flag.Parse()

	if *interactiveMode {
		err := interactive.Run()

		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		}

		return
	}

	if *usePassphrase {
		result, err := passphrase.Generate(*wordsCount, "-")

		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		}

		fmt.Println(result)
		return
	}

	if *strengthPassword != "" {
		result := strength.Analyze(*strengthPassword)

		fmt.Printf("Length: %d\n", result.Length)
		fmt.Printf("Character sets: %d\n", result.CharacterSets)
		fmt.Printf("Estimated entropy: %.1f bits\n", result.Entropy)
		fmt.Printf("Strength: %s\n", result.Level)

		return
	}

	options := generator.Options{
		Length:       *length,
		UseLowercase: *useLowercase,
		UseUppercase: *useUppercase,
		UseNumbers:   *useNumbers,
		UseSymbols:   *useSymbols,
	}

	password, err := generator.Generate(options)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	fmt.Println(password)
}
