package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"testing"
)

var dbTestName string = "database_test.db"

func TestORM(t *testing.T) {
	db, err := gorm.Open("sqlite3", dbTestName)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&Highlight{})

	db.Create(&Highlight{Text: "Storing ORM data", Url: "www.google.com", Book: "book of hard knocks", Posted: 0})

	var highlight Highlight
	var count int
	db.Where("posted = ?", 1).Find(&highlight).Count(&count)
	if count != 0 {
		t.Errorf("Found posted highlights when there weren't any")
	}
	db.Where("posted = ?", 0).Find(&highlight).Count(&count)
	if count != 1 {
		t.Errorf("Did not find posted highlights when there weren't any")
	}
	db.Delete(&highlight)
}

func TestGetUnposted(t *testing.T) {
	db, err := gorm.Open("sqlite3", dbTestName)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.Create(&Highlight{Text: "Storing ORM data", Url: "www.google.com", Book: "book of hard knocks", Posted: 0})
	db.Create(&Highlight{Text: "", Url: "", Book: "", Posted: 0})

	highlights := getUnposted(dbTestName)
	fmt.Println(highlights)
	fmt.Println(len(highlights))
	if len(highlights) != 2 {
		t.Errorf("Did not find two unposted highlights")
	}
	db.Delete(&highlights)
}

func TestInsert(t *testing.T) {

	highlight1 := Highlight{Text: "text", Url: "www.google.com", Book: "book of hard knocks", Posted: 0}
	highlight2 := Highlight{Text: "text", Url: "www.google.com2", Book: "book of hard knocks2", Posted: 0}

	insert(highlight1, dbTestName)
	insert(highlight2, dbTestName)

	db, err := gorm.Open("sqlite3", dbTestName)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	var highlights []Highlight
	db.Find(&highlights)
	if len(highlights) != 2 {
		t.Errorf("Did not find the two inserted highlights, found: %d", len(highlights))
	}

	db.Delete(&highlight1)
	db.Delete(&highlight2)
}
