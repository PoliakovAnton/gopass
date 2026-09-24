package passphrase

import (
	"crypto/rand"
	"errors"
	"math/big"
	"strings"
)

var words = []string{
	"apple",
	"river",
	"stone",
	"forest",
	"cloud",
	"silver",
	"rocket",
	"coffee",
	"orange",
	"winter",
	"summer",
	"mountain",
	"ocean",
	"thunder",
	"shadow",
	"garden",
	"planet",
	"castle",
	"sunset",
	"dragon",
}

func Generate(count int, separator string) (string, error) {
	if count <= 0 {
		return "", errors.New("word count must be greater than 0")
	}

	result := make([]string, count)
	max := big.NewInt(int64(len(words)))

	for i := 0; i < count; i++ {
		index, err := rand.Int(rand.Reader, max)
		if err != nil {
			return "", err
		}

		result[i] = words[index.Int64()]
	}

	return strings.Join(result, separator), nil
}
