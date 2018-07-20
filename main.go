package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Highlight struct {
	gorm.Model
	Id       int `gorm:"primary_key"`
	Text     string
	Url      string
	BookName string
	IsPosted int
}

func main() {
	db, err := gorm.Open("sqlite3", "highlightTweets.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&Highlight{})

	db.Create(&Highlight{Text: "Storing ORM data", Url: "www.google.com", BookName: "book of hard knocks", IsPosted: 0})

	var highlight Highlight
	//	db.First(&highlight, 1)                     // find product with id 1
	//	db.First(&highlight, "bookname = ?", "yes") // find product with code l1212

	//db.Model(&highlight).Update("Price", 2000)

	db.Delete(&highlight)
}
