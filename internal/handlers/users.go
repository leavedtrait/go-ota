package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/leavedtrait/go-ota/internal/db"
	"github.com/leavedtrait/go-ota/internal/utils"
)

type user struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateUserHandler(queries *db.Queries, w http.ResponseWriter, r *http.Request) {
	var user user
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, http.ErrBodyNotAllowed.Error(), 400)
		return
	}
	userExists, _ := queries.CheckUserExists(context.Background(), user.Email)

	if userExists {
		http.Error(w, "User email already exists", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json") // Add content-type header

	password, err := utils.HashPassword(user.Password)
	if err != nil {
		println(err.Error())
		http.Error(w, "Server Error", 500)
		return
	}
	userInDB, err := queries.CreateUser(context.Background(), db.CreateUserParams{
		Email:    user.Email,
		Name:     user.Name,
		Password: password,
	})
	if err != nil {
		http.Error(w, "Please try again", http.StatusFailedDependency)
		return
	}
	type data struct {
		Id    int64  `json:"id"`
		Email string `json:"email"`
	}

	res, _ := utils.CustomJsonResponse("Success. User created", data{Id: userInDB.ID, Email: user.Email}, http.StatusCreated)
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func GetUserByIDHandler(queries *db.Queries, w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	_, err := utils.ValidateJWT(tokenString)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	w.Header().Set("Content-Type", "application/json") // Add content-type header
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid User id", http.StatusBadRequest)
		return
	}
	user, err := queries.GetUser(context.Background(), int64(id))
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Create a new struct without the password field
	type UserResponse struct {
		ID    int64  `json:"id"`
		Email string `json:"email"`
		Name  string `json:"name"`
	}

	userResponse := UserResponse{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
	}

	res, err := utils.CustomJsonResponse("Success,found user", userResponse, http.StatusOK)
	if err != nil {
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}
	w.Write(res)
}

func UpdateUserHandler(queries *db.Queries, w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	_, err := utils.ValidateJWT(tokenString)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	w.Header().Set("Content-Type", "application/json") // Add content-type header
	var user user
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	password, err := utils.HashPassword(user.Password)
	if err != nil {
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}

	arg := db.UpdateUserParams{
		Email:    user.Email,
		Name:     user.Name,
		Password: password,
	}

	userInDB, err := queries.UpdateUser(context.Background(), arg)

	if err != nil {
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}

	type data struct {
		Id    int64  `json:"id"`
		Email string `json:"email"`
	}

	res, err := utils.CustomJsonResponse("Success. User updated", data{Id: userInDB.ID, Email: userInDB.Email}, http.StatusOK)

	if err != nil {
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteUserByIDHandler(queries *db.Queries, w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	_, err := utils.ValidateJWT(tokenString)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json") // Add content-type header
	id, _ := strconv.Atoi(r.PathValue("id"))
	err = queries.DeleteUser(context.Background(), int64(id))
	if err != nil {
		println(err.Error())
		http.Error(w, "Invalid User id", http.StatusBadRequest)
		return
	}
	res, _ := utils.CustomJsonResponse("Success,deleted user", nil, http.StatusOK)
	w.Write(res)
}

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginResponse struct {
	ID    int64  `json:"id"`
	Token string `json:"token"`
}

func LoginHandler(queries *db.Queries, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req loginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user, err := queries.GetUserByEmail(context.Background(), req.Email)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	if !utils.CheckPasswordHash(req.Password, user.Password) { // Compare provided password with stored hash
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token, err := utils.GenerateJWT(user.ID) // Assuming you have JWT generation utils
	if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	response, err := utils.CustomJsonResponse("Login successful", loginResponse{ID: user.ID, Token: token}, http.StatusOK)
	if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	w.Write(response)
}
