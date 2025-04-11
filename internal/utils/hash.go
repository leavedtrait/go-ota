package utils

import (
	"encoding/json"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}


type CustomResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
	Status  int    `json:"status"`
}

func CustomJsonResponse(msg string, data any, status int) ([]byte, error) {
	res := CustomResponse{msg, data, status}
	return json.Marshal(res)
}
