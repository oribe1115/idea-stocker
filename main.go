package main

import (
	"net/http"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
)

type Ideas struct {
	ID        uint      `gorm:"primary_key" json:"id,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	Status    string    `json:"status,omitempty"`
	Idea      string    `json:"idea,omitempty"`
}

var (
	db *gorm.DB
)

func main() {
	databaseURL := os.Getenv("DATABASE_URL")
	_db, err := gorm.Open("postgres", databaseURL)
	if err != nil {
		panic("failed to connect database")
	}
	db = _db
	defer db.Close()

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World.\n")
	})

	e.GET("/create", makeTable)

	e.GET("/newidea", addNewIdea)

	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	e.Start(":" + port)
}

func makeTable(c echo.Context) error {
	// 一度作ってみるのに使った後は間違って作りまくってしまわないようにコメントアウトでもしておくこと
	db.CreateTable(&Ideas{})
	return c.String(http.StatusOK, "Created!\n")
}

func addNewIdea(c echo.Context) error {
	newIdea := Ideas{CreatedAt: time.Now(), Status: "nowDoing", Idea: "テスト"}
	db.Create(&newIdea)
	return c.JSON(http.StatusOK, newIdea)
}
