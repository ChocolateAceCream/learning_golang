package db

import (
	"database/sql"
	"fmt"
	"math/rand"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type MyDB sql.DB

func Conn() *sql.DB {
	dbHost := "localhost"
	dbPort := 3306
	dbUser := "root"
	dbPass := "di880817"
	dbName := "blog_test"
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	db, err := sql.Open("mysql", dbURI)
	if err != nil {
		panic(err)
	}
	// defer db.Close()

	// Check if the connection is successful
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to the MySQL database!")
	return db
}

type Data struct {
	key1         string
	key2         int
	key3         string
	key_part1    string
	key_part2    string
	key_part3    string
	common_field string
}

func RandInsert(db *sql.DB) {
	// Insert data into the users table
	k2 := 1
	count := 0
	for count < 10001 {
		d := Data{
			key1:         RandString(),
			key2:         k2,
			key3:         RandString(),
			key_part1:    RandString(),
			key_part2:    RandString(),
			key_part3:    RandString(),
			common_field: RandString(),
		}
		r, err := db.Exec("INSERT INTO single_table (key1, key2, key3, key_part1, key_part2, key_part3, common_field) VALUES (?, ?, ?, ?, ?, ?, ?)", d.key1, d.key2, d.key3, d.key_part1, d.key_part2, d.key_part3, d.common_field)
		if err != nil {
			panic(err)
		}
		id, err := r.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Inserted data with ID %d\n", id)
		k2 += RandNum()
		count++
	}
}

func RandNum() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(3) + 1
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func RandString() string {
	rand.Seed(time.Now().UnixNano())
	l := rand.Intn(5) + 5
	b := make([]byte, l)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
