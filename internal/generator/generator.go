package generator

import (
	"crypto/rand"
	"errors"
	"math/big"
)

const (
	lowercase = "abcdefghijklmnopqrstuvwxyz"
	uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers   = "0123456789"
	symbols   = "!@#$%^&*()-_=+[]{}<>?"
)

type Options struct {
	Length       int
	UseLowercase bool
	UseUppercase bool
	UseNumbers   bool
	UseSymbols   bool
}

func Generate(options Options) (string, error) {
	if options.Length <= 0 {
		return "", errors.New("password length must be greater than 0")
	}

	characters := buildCharacterSet(options)

	if characters == "" {
		return "", errors.New("at least one character set must be enabled")
	}

	password := make([]byte, options.Length)
	max := big.NewInt(int64(len(characters)))

	for i := 0; i < options.Length; i++ {
		randomIndex, err := rand.Int(rand.Reader, max)
		if err != nil {
			return "", err
		}

		password[i] = characters[randomIndex.Int64()]
	}

	return string(password), nil
}

func buildCharacterSet(options Options) string {
	characters := ""

	if options.UseLowercase {
		characters += lowercase
	}

	if options.UseUppercase {
		characters += uppercase
	}

	if options.UseNumbers {
		characters += numbers
	}

	if options.UseSymbols {
		characters += symbols
	}

	return characters
}
