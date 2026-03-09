package lib

import (
	"github.com/matthewhartstonge/argon2"
)

var argon = argon2.DefaultConfig()

func HashPassword(password string) (string, error) {

	encoded, err := argon.HashEncoded([]byte(password))
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func VerifyPassword(password string, encodedHash string) bool {
	ok, err := argon2.VerifyEncoded([]byte(password), []byte(encodedHash))
	if err != nil {
		return false
	}

	return ok
}
