package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Highlight struct {
	Text   string
	Url    string
	Book   string
	Posted int
}

func main() {
	db, err := gorm.Open("sqlite3", "highlightTweets.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&Highlight{})

	db.Create(&Highlight{Text: "Storing ORM data", Url: "www.google.com", Book: "book of hard knocks", Posted: 1})

	var highlight Highlight
	db.Where("posted = ?", 0).First(&highlight)
	fmt.Println(highlight)
	//	db.First(&highlight, 1)                     // find product with id 1
	//	db.First(&highlight, "bookname = ?", "yes") // find product with code l1212

	//db.Model(&highlight).Update("Price", 2000)

	db.Delete(&highlight)
}
