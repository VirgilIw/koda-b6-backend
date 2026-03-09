package lib

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

type CustomClaims struct { // claim dilakukan ketika generate jwt
	Id int `json:"id"`
	jwt.RegisteredClaims
}

func GenerateToken(userId int) (string, error) {
	godotenv.Load()
	mySecret := os.Getenv("SECRET_KEY")

	// claim itu isinya data-data yang disimpan di dalam JWT.
	// generate claim, claims = payload
	claims := CustomClaims{
		userId,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// token akan dikembalikan
	// signingmethodHS256 = algoritma jwt
	// signedstring = signature, diberikan setelah token di generate
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(mySecret)) // memberikan ttd, secret key yang disimpan dari .env

	if err != nil {
		return "", err
	}

	// struktur jwt (Header.Payload.Signature)
	return tokenString, nil
}
