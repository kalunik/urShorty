package utils

import (
	"bytes"
	"math/rand"
)

func generateShortURL(full string) (string, error) {

	const (
		base62Chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
		base        = 62
		shortLength = 8
	)

	shortURL := &bytes.Buffer{}

	for i := 0; i < shortLength; i++ {
		shortURL.WriteByte(base62Chars[rand.Intn(base)])
	}

	return shortURL.String(), nil
}

//func main() {
//	str, _ := generateShortURL("https://github.com/AleksK1NG/Go-Clean-Architecture-REST-API/blob/master/pkg/utils/http.go#L119")
//	fmt.Println(str)
//	return
//}
