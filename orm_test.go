package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"testing"
)

func TestORM(t *testing.T) {
	db, err := gorm.Open("sqlite3", "test.db")
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
