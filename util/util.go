package util

import (
	"math/rand"
	"time"
)

//生成自定义字符
func RandomString(n int) string {
	var letters = []byte("asdfghjklqwertyuiopzxcvbnm")
	result := make([]byte,n)

	rand.Seed(time.Now().Unix())
	for i := range result{
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}