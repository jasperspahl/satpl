package utils

import "math/rand"

func RandomString(length int) string {
	random_str := make([]byte, length)

	for i := 0; i < length; i++ {
		random_str[i] = byte(65 + rand.Intn(25))
	}

	return string(random_str)
}
