package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"github.com/leavedtrait/go-ota/cmd/api"
	"github.com/leavedtrait/go-ota/internal/db"
	"github.com/leavedtrait/go-ota/internal/server"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	conn, err := db.ConnectToDB()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	slog.Info("DB connected")
	defer conn.Close(context.Background())

	router := server.NewRouter(conn)

	srv := api.NewServer(":3000", router)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err.Error())
	}
}
