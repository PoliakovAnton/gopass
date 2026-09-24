package interactive

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/PoliakovAnton/gopass/internal/generator"
)

func Run() error {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("gopass interactive mode")
	fmt.Println()

	length, err := askInt(reader, "Password length", 16)
	if err != nil {
		return err
	}

	useUppercase, err := askBool(reader, "Include uppercase letters", false)
	if err != nil {
		return err
	}

	useNumbers, err := askBool(reader, "Include numbers", false)
	if err != nil {
		return err
	}

	useSymbols, err := askBool(reader, "Include symbols", false)
	if err != nil {
		return err
	}

	options := generator.Options{
		Length:       length,
		UseLowercase: true,
		UseUppercase: useUppercase,
		UseNumbers:   useNumbers,
		UseSymbols:   useSymbols,
	}

	password, err := generator.Generate(options)
	if err != nil {
		return err
	}

	fmt.Println()
	fmt.Println("Generated password:")
	fmt.Println(password)

	return nil
}

func askInt(reader *bufio.Reader, question string, defaultValue int) (int, error) {
	for {
		fmt.Printf("%s [%d]: ", question, defaultValue)

		input, err := reader.ReadString('\n')
		if err != nil {
			return 0, err
		}

		input = strings.TrimSpace(input)

		if input == "" {
			return defaultValue, nil
		}

		value, err := strconv.Atoi(input)
		if err != nil || value <= 0 {
			fmt.Println("Please enter a positive number.")
			continue
		}

		return value, nil
	}
}

func askBool(reader *bufio.Reader, question string, defaultValue bool) (bool, error) {
	defaultText := "y/N"

	if defaultValue {
		defaultText = "Y/n"
	}

	for {
		fmt.Printf("%s [%s]: ", question, defaultText)

		input, err := reader.ReadString('\n')
		if err != nil {
			return false, err
		}

		input = strings.ToLower(strings.TrimSpace(input))

		if input == "" {
			return defaultValue, nil
		}

		switch input {
		case "y", "yes":
			return true, nil
		case "n", "no":
			return false, nil
		default:
			fmt.Println("Please answer y or n.")
		}
	}
}
