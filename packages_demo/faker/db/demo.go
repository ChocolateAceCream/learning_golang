package db

import (
	"database/sql"
	"fmt"
)

type User struct {
	ID        int    `json:"userId"`
	Name      string `json:"name"`
	Age       int    `json:"age"`
	Email     string `json:"email"`
	Address   string `json:"address"`
	CreatedAt string `json:"createdAt"`
	DeletedAt string `json:"deletedAt"`
}

func CreateDemoTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
    user_id INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
    name VARCHAR(50) NOT NULL,
    age INT NOT NULL,
    email VARCHAR(255) NOT NULL,
    address VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
		UNIQUE KEY unique_email (user_id, email)  -- Composite unique index
	) PARTITION BY range (user_id) (
			PARTITION p0 VALUES LESS THAN (10000),
			PARTITION p1 VALUES LESS THAN (20000),
			PARTITION p2 VALUES LESS THAN (30000),
			PARTITION p3 VALUES LESS THAN (40000),
			PARTITION p4 VALUES LESS THAN (50000),
			PARTITION p5 VALUES LESS THAN MAXVALUE
	);
	`
	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("fail to create table: %w", err)
	}
	return nil
}

func Insert(myDB *sql.DB, users []User) error {
	fmt.Println("----------Insert-----------")
	tx, err := myDB.Begin()
	if err != nil {
		return fmt.Errorf("failed to start tx: %w", err)
	}
	query := `INSERT INTO users (user_id, name, age, email, address) VALUES `
	values := make([]interface{}, 0, len(users)*7)
	for _, user := range users {
		query += "(?,?,?,?,?),"
		values = append(values, user.ID, user.Name, user.Age, user.Email, user.Address)
	}
	query = query[:len(query)-1] // remove last ,
	statement, err := tx.Prepare(query)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer statement.Close()

	_, err = statement.Exec(values...)
	// lastId, _ := result.LastInsertId()
	// fmt.Println("------result--", lastId)
	if err != nil {
		return fmt.Errorf("failed to execute statement: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	fmt.Println("Records inserted successfully")
	return nil
}
