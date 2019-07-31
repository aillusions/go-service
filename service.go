package main

import (
	"crypto/rand"
	"fmt"
	"net/http"
	"time"
)

var token string

func main() {

	var err error

	token, err = GenerateRandomString(6)
	if err != nil {
		panic(err)
	}
	fmt.Println(token)

	go bySchedule()

	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)

}

func bySchedule() {
	for {
		fmt.Println("Hello scheduler...", token)
		time.Sleep(1 * time.Second)
	}
}
func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from: %s", token)
	fmt.Println("response completed..", token)
}

func GenerateRandomString(n int) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
	bytes, err := GenerateRandomBytes(n)
	if err != nil {
		return "", err
	}
	for i, b := range bytes {
		bytes[i] = letters[b%byte(len(letters))]
	}
	return string(bytes), nil
}

func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}
