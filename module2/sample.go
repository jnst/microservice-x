package module2

import "math/rand"

func GenerateString() string {
	chars := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	var result []byte
	for i := 0; i < 20; i++ {
		result = append(result, chars[rand.Intn(len(chars))])
	}
	return string(result)
}
