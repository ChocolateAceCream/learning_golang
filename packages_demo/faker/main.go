package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/ChocolateAceCream/faker-demo/db"
	"github.com/jaswdr/faker/v2"
)

func main() {
	var wg sync.WaitGroup
	start := time.Now()
	fmt.Println("---")
	fake := faker.New()
	fmt.Println(fake.Person().Name())
	fmt.Println(fake.RandomNumber(12))
	fmt.Println(fake.Address().City())
	filename := "sample.csv"
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal("fail to open file", err)
	}
	reader := csv.NewReader(file)
	_, err = reader.Read()
	if err != nil {
		if err == io.EOF {
			if err = InitFileData(file, 100000); err != nil {
				fmt.Println("fail to init file", err)
			}
		}
	}

	myDB, err := db.Init()
	if err != nil {
		fmt.Println("fail to init db", err)
	}
	err = db.CreateDemoTable(myDB)
	if err != nil {
		fmt.Println("fail to create demo table", err)
	}

	pool := make(chan []db.User, 10)
	results := make(chan int)
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go Worker(pool, results, myDB, &wg)
	}
	badRecordIDs, err := ReadUserData(myDB, filename, 100, pool)
	close(pool)

	close(results)
	wg.Wait()

	// close(pool)

	if err != nil {
		fmt.Println(badRecordIDs)
	}
	duration := time.Since(start)
	fmt.Println(duration)
	// defer GlobalDB.Close()
}

func Worker(jobs <-chan []db.User, badRecords chan<- int, myDB *sql.DB, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		fmt.Println("------------jobs--------------", j)
		db.Insert(myDB, j)
	}
}

func InitFileData(file *os.File, RecordCount int) error {
	writer := csv.NewWriter(file)
	defer writer.Flush()
	header := []string{"ID", "Name", "Age", "Email", "Address"}
	if err := writer.Write(header); err != nil {
		return err
	}
	fake := faker.New()
	for i := 0; i < RecordCount; i++ {
		row := []string{strconv.Itoa(i), fake.Person().Name(), strconv.Itoa(fake.RandomNumber(2)), fake.Person().Contact().Email, fake.Address().Address()}
		if err := writer.Write(row); err != nil {
			return err
		}
	}
	// Check for any error encountered during the write process
	if err := writer.Error(); err != nil {
		return err
	}
	return nil
}

func ReadUserData(myDB *sql.DB, filename string, batchSize int, pool chan []db.User) (badRecords []int, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return badRecords, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	_, err = reader.Read()
	if err != nil {
		return badRecords, fmt.Errorf("failed to read file header: %w", err)
	}
	users := []db.User{}
	count := 0
	for {
		count++
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return badRecords, fmt.Errorf("failed to read file data: %w", err)
		}
		id, err := strconv.Atoi(record[0])
		if err != nil {
			badRecords = append(badRecords, id)
			continue
		}
		name := record[1]
		age, err := strconv.Atoi(record[2])
		if err != nil {
			badRecords = append(badRecords, id)
			continue
		}
		email := record[3]
		address := record[4]
		user := db.User{
			ID:      id,
			Name:    name,
			Age:     age,
			Email:   email,
			Address: address,
		}
		users = append(users, user)
		fmt.Println(len(users))
		if count%batchSize == 0 {
			pool <- users
			users = []db.User{}
		}
	}
	fmt.Println("------left over---------")

	if len(users) > 0 {
		pool <- users
	}
	fmt.Println("------done reading---------")

	return badRecords, err
}
