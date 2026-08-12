// package utils

// import (
// 	"fmt"
// 	"os"
// 	"time"

// 	"github.com/golang-jwt/jwt/v5"
// )

// func GenerateToken(userID uint) (string, error) {
// 	claims := jwt.MapClaims{
// 		"UserID": userID,
// 		"exp":    time.Now().Add(24 * time.Hour).Unix(),
// 	}
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	return token.SignedString([]byte("JWT_SECRET"))
// }
// func ValidateToken(tokenString string) (uint, error) {

// 	secret := os.Getenv("JWT_SECRET")

// 	token, err := jwt.Parse(
// 		tokenString,
// 		func(token *jwt.Token) (interface{}, error) {

// 			// Check signing algorithm
// 			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 				return nil, fmt.Errorf("unexpected signing method")
// 			}

// 			return []byte(secret), nil
// 		},
// 	)

// 	if err != nil || !token.Valid {
// 		return 0, fmt.Errorf("invalid token")
// 	}

// 	claims, ok := token.Claims.(jwt.MapClaims)

// 	if !ok {
// 		return 0, fmt.Errorf("invalid claims")
// 	}

// 	userIDFloat, ok := claims["user_id"].(float64)

// 	if !ok {
// 		return 0, fmt.Errorf("user_id not found")
// 	}

// 	return uint(userIDFloat), nil
// }

package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const JWTSecret = "JWT_SECRET"

func GenerateToken(userID uint) (string, error) {

	claims := jwt.MapClaims{
		"UserID": userID,
		"exp":    time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	return token.SignedString([]byte(JWTSecret))
}

func ValidateToken(tokenString string) (uint, error) {

	token, err := jwt.Parse(
		tokenString,
		func(token *jwt.Token) (interface{}, error) {

			// Signing method check
			if token.Method != jwt.SigningMethodHS256 {
				return nil, fmt.Errorf("unexpected signing method")
			}

			return []byte(JWTSecret), nil
		},
	)

	if err != nil {
		return 0, err
	}

	if !token.Valid {
		return 0, fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return 0, fmt.Errorf("invalid claims")
	}

	// GenerateToken mein "UserID" use kiya tha
	userIDFloat, ok := claims["UserID"].(float64)

	if !ok {
		return 0, fmt.Errorf("UserID not found")
	}

	return uint(userIDFloat), nil
}
