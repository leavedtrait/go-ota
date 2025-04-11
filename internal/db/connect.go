package db

import (
	"context"
	"os"
	"github.com/jackc/pgx/v5"
)
func ConnectToDB() (*pgx.Conn ,error){
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	return conn,err
}