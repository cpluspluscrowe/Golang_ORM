package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Highlight struct {
	Text   string
	Url    string
	Book   string
	Posted int
}

var prodDbName string = "highlightTweets.db"

func main() {
	db, err := gorm.Open("sqlite3", prodDbName)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&Highlight{})

	db.Create(&Highlight{Text: "Storing ORM data", Url: "www.google.com", Book: "book of hard knocks", Posted: 1})

	var highlight Highlight
	db.Where("posted = ?", 0).First(&highlight)

	//	db.First(&highlight, 1)                     // find product with id 1
	//	db.First(&highlight, "bookname = ?", "yes") // find product with code l1212

	//db.Model(&highlight).Update("Price", 2000)

}

func Insert(highlight Highlight) {
	insert(highlight, prodDbName)
}

func GetUnpostedHighlights() []Highlight {
	return getUnposted(prodDbName)
}

func insert(highlight Highlight, dbName string) {
	db, err := gorm.Open("sqlite3", prodDbName)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.Create(highlight)
}

func getUnposted(dbName string) []Highlight {
	db, err := gorm.Open("sqlite3", dbName)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	var highlights []Highlight
	db.Where("posted = ?", 0).Find(&highlights)
	return highlights
}
