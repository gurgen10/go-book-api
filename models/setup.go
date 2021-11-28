package models

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

type Person struct {
	gorm.Model
	Name  string
	Email string `gorm:"typevarchar(100);unique_index"`
}
type Book struct {
	gorm.Model
	Title      string
	Author     string
	CallNumber string `gorm:"unique_index"`
}

var (
	people    = &Person{Name: "Gug", Email: "gurgen@retailmerchanygroup.com"}
	booksData = []Book{
		{Title: "Masha & Bear", Author: "Karen", CallNumber: "1516155"},
		{Title: "Life and fight", Author: "Khoren", CallNumber: "145585"},
	}
)

var DB *gorm.DB

func Connection() {
	fmt.Println("Connection modil")
	godotenv.Load(".env")
	dialect := os.Getenv("DIALECT")
	host := os.Getenv("HOSTNAME")
	port := os.Getenv("PORT")
	user := os.Getenv("DBUSER")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DBNAME")
	dbURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", host, port, user, dbname, password)
	fmt.Println("host=localhost port=5432 user=postgres dbname=book_keeper password=123")
	fmt.Println(user)
	db, err := gorm.Open(dialect, dbURI)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connection was done successfully.")
	}

	// db.AutoMigrate(&Person{}, &Book{})
	// for i, book := range booksData {
	// 	fmt.Print(fmt.Print("dsdsdsdsd", i))
	// 	db.Create(&book)
	// }

	// db.Create(people)

	// defer db.Close()
	DB = db

}
