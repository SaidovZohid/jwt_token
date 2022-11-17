package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
)

const secretkey = "superadmin"

// CreateToken creates a new token
func CreateToken(username, email string, duration time.Duration) (string, *Payload, error) {
	payload, err := NewPayload(username, email, duration)
	if err != nil {
		return "",  payload, err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	token, err := jwtToken.SignedString([]byte(secretkey))
	if err != nil {
		return "",  payload, err
	}
	
	return token, payload, err
}