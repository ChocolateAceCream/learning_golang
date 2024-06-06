package db

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func Init() (*sql.DB, error) {
	username := os.Getenv("MYSQL_USERNAME")
	password := os.Getenv("MYSQL_PASSWORD")
	fmt.Println(username, password)
	dsn := fmt.Sprintf("%s:%s@tcp(localhost:3306)/june", username, password)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// Ping the database to ensure the connection is established
	if err = db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("fail to ping db: %w", err)
	}
	return db, nil
}
