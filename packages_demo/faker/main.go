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

	"github.com/ChocolateAceCream/faker-demo/db"
	"github.com/jaswdr/faker/v2"
)

var wg sync.WaitGroup

func main() {
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
	badRecordIDs, err := ReadUserData(myDB, filename, 100)
	if err != nil {
		fmt.Println(badRecordIDs)
	}
	// defer GlobalDB.Close()
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

func ReadUserData(myDB *sql.DB, filename string, batchSize int) (badRecords []int, err error) {
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
		fmt.Println(user)
		fmt.Println(len(users))
		if count%batchSize == 0 {
			wg.Add(1)
			go func() {

				err = db.Insert(myDB, users)
				if err != nil {
					for _, u := range users {
						badRecords = append(badRecords, u.ID)
					}
				}
				wg.Done()
			}()

			users = []db.User{}
		}
	}
	if len(users) > 0 {
		err = db.Insert(myDB, users)
		if err != nil {
			for _, u := range users {
				badRecords = append(badRecords, u.ID)
			}
		}
	}
	wg.Wait()
	return badRecords, err
}
