package mod

import (
	"testing"
)

func TestEncrypt(t *testing.T) {
	password := []byte("*RTFUGIHOD&TUGGIYKl")
	data := []byte("This encryption algorithm is faster than aes256 up to 40kb but how secure is it?")
	for x, _ := range data {
		p := (password[x%len(password)])
		data[x] = data[x] + byte(p)
	}
	println(string(data))
	for x, _ := range data {
		p := (password[x%len(password)])
		data[x] = data[x] - byte(p)
	}
	println(string(data))
}
