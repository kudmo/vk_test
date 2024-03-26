// A package with functions for working with authorization via jwt
package auth

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"log"
	"strings"
	"time"
	"vk_test/internal/config"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

// Description of JWT fields
type JWTClaims struct {
	jwt.RegisteredClaims
	UserId int `json:"user-id"`
	Type   string
}

// The function of getting the user ID from the transferred token and checking the token for correctness
func TokenGetUserId(c echo.Context) (int, bool) {
	headers := c.Request().Header["Authorization"]
	if len(headers) == 0 {
		return -1, true
	}
	tokenString := strings.ReplaceAll(headers[0], "Bearer ", "")

	// pass your custom claims to the parser function
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(config.SecretKeyJWT), nil
	})
	if err != nil {
		log.Println(err.Error())
		return -1, false
	}
	myClaims := token.Claims.(*JWTClaims)
	return myClaims.UserId, true

}

// Returns encoded token, error.
//
// The current duration of the token is 12 hours
func CalculateToken(userId int) (string, error) {
	claims := &JWTClaims{
		UserId: userId,
		Type:   "access",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 12)),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(config.SecretKeyJWT))
	return t, err
}

// Password hashing function
func HashPassword(password string) string {
	sum := sha256.Sum256([]byte(password))
	return hex.EncodeToString(sum[:])
}
