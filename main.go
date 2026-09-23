package main

import (
	"crypto/rand"
	"fmt"
)

const characters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generatePassword(length int) (string, error) {
	password := make([]byte, length)

	for i := 0; i < length; i++ {
		randomByte := make([]byte, 1)

		_, err := rand.Read(randomByte)
		if err != nil {
			return "", err
		}

		randomIndex := int(randomByte[0]) % len(characters)

		password[i] = characters[randomIndex]
	}

	return string(password), nil
}

func main() {
	password, err := generatePassword(16)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Generated password:", password)
}
