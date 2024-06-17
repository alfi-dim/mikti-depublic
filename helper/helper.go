package helper

import (
	"math/rand"
	"time"
)

type ResponseToJson map[string]interface{}

func GenerateId(n int, name string) string {
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	charset := []rune("1234567890")
	letters := make([]rune, n)
	for i := range letters {
		letters[i] = charset[r.Intn(len(charset))]
	}

	id := name + "-" + string(letters)
	return id
}
