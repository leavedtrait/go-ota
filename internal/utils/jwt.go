package utils

import (
        "fmt"
        "os"
        "time"

        "github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET")) // Store your secret securely!

// GenerateJWT generates a JWT token for a user.
func GenerateJWT(userID int64) (string, error) {
        token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
                "user_id": userID,
                "exp":     time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
        })
        tokenString, err := token.SignedString(jwtSecret)
        return tokenString, err
}

// ValidateJWT validates a JWT token and returns the user ID.
func ValidateJWT(tokenString string) (int64, error) {
        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
                if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                        return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
                }
                return jwtSecret, nil
        })

        if err != nil {
                return 0, err
        }

        if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
                userID, ok := claims["user_id"].(float64) // user_id is float64 in claims
                if !ok {
                        return 0, fmt.Errorf("invalid user ID in token")
                }
                return int64(userID), nil
        }
        return 0, fmt.Errorf("invalid token")
}