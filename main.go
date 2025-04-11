package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/leavedtrait/go-ota/cmd/api"
	"github.com/leavedtrait/go-ota/internal/db"
	"github.com/leavedtrait/go-ota/internal/server"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	conn, err := db.ConnectToDB()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	slog.Info("DB connected")
	defer conn.Close(context.Background())

	queries := db.New(conn)
	router := server.NewRouter(queries)

	srv := api.NewServer(":3000", router)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err.Error())
	}
}
