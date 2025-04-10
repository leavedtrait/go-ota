package main

import (
	"context"
	"fmt"
	"log"
<<<<<<< HEAD
	"log/slog"
	"os"

	"github.com/joho/godotenv"
=======
	"os"

>>>>>>> eaf22056321c4c0cfb9f8c594f068d99eed5d4a2
	"github.com/leavedtrait/go-ota/cmd/api"
	"github.com/leavedtrait/go-ota/internal/db"
	"github.com/leavedtrait/go-ota/internal/server"
)

func main() {
<<<<<<< HEAD
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	conn, err := db.ConnectToDB()
=======
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	conn ,err := db.ConnectToDB() 
>>>>>>> eaf22056321c4c0cfb9f8c594f068d99eed5d4a2
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	slog.Info("DB connected")
	defer conn.Close(context.Background())

	router := server.NewRouter(conn)
<<<<<<< HEAD

=======
	
>>>>>>> eaf22056321c4c0cfb9f8c594f068d99eed5d4a2
	srv := api.NewServer(":3000", router)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err.Error())
	}
}
