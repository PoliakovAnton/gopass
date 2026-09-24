package generator

import (
	"crypto/rand"
	"math/big"
)

const characters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func Generate(length int) (string, error) {
	password := make([]byte, length)

	max := big.NewInt(int64(len(characters)))

	for i := 0; i < length; i++ {
		randomIndex, err := rand.Int(rand.Reader, max)
		if err != nil {
			return "", err
		}

		password[i] = characters[randomIndex.Int64()]
	}

	return string(password), nil
}
