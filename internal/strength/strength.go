package strength

import (
	"math"
)

type Result struct {
	Length        int
	CharacterSets int
	Entropy       float64
	Level         string
}

func Analyze(password string) Result {
	length := len(password)
	characterSets := 0
	alphabetSize := 0

	if containsLowercase(password) {
		characterSets++
		alphabetSize += 26
	}

	if containsUppercase(password) {
		characterSets++
		alphabetSize += 26
	}

	if containsNumbers(password) {
		characterSets++
		alphabetSize += 10
	}

	if containsSymbols(password) {
		characterSets++
		alphabetSize += 20
	}

	entropy := 0.0

	if alphabetSize > 0 && length > 0 {
		entropy = float64(length) * math.Log2(float64(alphabetSize))
	}

	level := calculateLevel(entropy)

	return Result{
		Length:        length,
		CharacterSets: characterSets,
		Entropy:       entropy,
		Level:         level,
	}
}

func calculateLevel(entropy float64) string {
	switch {
	case entropy < 40:
		return "weak"
	case entropy < 60:
		return "moderate"
	case entropy < 80:
		return "strong"
	default:
		return "very strong"
	}
}

func containsLowercase(password string) bool {
	for _, char := range password {
		if char >= 'a' && char <= 'z' {
			return true
		}
	}

	return false
}

func containsUppercase(password string) bool {
	for _, char := range password {
		if char >= 'A' && char <= 'Z' {
			return true
		}
	}

	return false
}

func containsNumbers(password string) bool {
	for _, char := range password {
		if char >= '0' && char <= '9' {
			return true
		}
	}

	return false
}

func containsSymbols(password string) bool {
	for _, char := range password {
		if !((char >= 'a' && char <= 'z') ||
			(char >= 'A' && char <= 'Z') ||
			(char >= '0' && char <= '9')) {
			return true
		}
	}

	return false
}
