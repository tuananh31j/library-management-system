package helper

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"

	"github.com/tuananh31j/library-management-system/model"
	"github.com/tuananh31j/library-management-system/utils"
	"golang.org/x/exp/rand"
	"gorm.io/gorm"
)

func getRandomNumber(n int) int {
	if n <= 0 {
		return 0
	}
	rand.Seed(uint64(time.Now().UnixNano()))
	return rand.Intn(n + 1)
}

func InsertData(db *gorm.DB, fileName string, modal any) {
	file, err := os.Open(fileName)
	if err != nil {
		utils.Log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	// Read file
	records, errRead := csv.NewReader(file).ReadAll()
	if errRead != nil {
		utils.Log.Fatalf("Failed to read file: %v", errRead)
	}
	switch v := modal.(type) {
	case model.Author:
		CreateManyAuthors(db, v, records)
	case model.Book:
		CreateManyBooks(db, v, records)
	case model.Borrower:
		CreateManyBorrowers(db, v, records)
	case model.BorrowerBooks:
		CreateManyBorrowerBooks(db, v, records)
	}

}

func CreateManyAuthors(db *gorm.DB, authors model.Author, records [][]string) {
	// Insert data
	for _, row := range records[1:] {
		// Insert data to database
		record := model.Author{
			Name:      row[1],
			CreatedAt: getDate(row[2]),
			UpdatedAt: getDate(row[3]),
		}

		db.Exec("INSERT INTO authors (name, created_at, updated_at) VALUES (?,?,?)", record.Name, record.CreatedAt, record.UpdatedAt)

	}

}
func CreateManyBooks(db *gorm.DB, authors model.Book, records [][]string) {
	var authorIDs []int
	// Get all author IDs
	db.Raw("SELECT id FROM authors").Scan(&authorIDs)
	for _, row := range records[1:] {
		// Insert data to database
		record := model.Book{
			Name: row[1],
			// AuthorID:  authorIDs[getRandomNumber(len(authorIDs)-1)],
			CreatedAt: getDate(row[3]),
			UpdatedAt: getDate(row[4]),
		}

		db.Exec("INSERT INTO books (name, author_id,created_at, updated_at) VALUES (?,?,?,?)", record.Name, record.AuthorID, record.CreatedAt, record.UpdatedAt)
	}

}

func CreateManyBorrowers(db *gorm.DB, authors model.Borrower, records [][]string) {

	for _, row := range records[1:] {
		// Insert data to database
		record := model.Borrower{
			Name:      row[1],
			Email:     row[2],
			Phone:     row[3],
			Address:   row[4],
			CreatedAt: getDate(row[3]),
			UpdatedAt: getDate(row[4]),
		}

		db.Exec("INSERT INTO borrowers (name, email, phone,address,created_at, updated_at) VALUES (?,?,?,?,?,?)", record.Name, record.Email, record.Phone, record.Address, record.CreatedAt, record.UpdatedAt)
	}

}

func CreateManyBorrowerBooks(db *gorm.DB, authors model.BorrowerBooks, records [][]string) {
	var bookIDs []int
	// Get all author IDs
	db.Raw("SELECT id FROM books").Scan(&bookIDs)

	for _, row := range records[1:] {
		// Insert data to database
		record := model.BorrowerBooks{
			BorrowerID: parseInt(row[1]),
			BookID:     bookIDs[getRandomNumber(len(bookIDs)-1)],
			CreatedAt:  getDate(row[3]),
			UpdatedAt:  getDate(row[4]),
		}

		db.Exec("INSERT INTO borrower_books (borrower_id,book_id,created_at, updated_at) VALUES (?,?,?,?)", record.BorrowerID, record.BookID, record.CreatedAt, record.UpdatedAt)
	}

}

func parseInt(str string) int {
	var value int
	fmt.Sscanf(str, "%d", &value)
	return value
}

func getDate(dateStr string) time.Time {
	layout := "2006-01-02"
	date, err := time.Parse(layout, dateStr)
	if err != nil {
		return time.Now()
	}
	return date
}
