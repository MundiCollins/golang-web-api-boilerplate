package main

import (
	"database/sql"
	"github.com/MundiCollins/golang-web-api-boilerplate/handlers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/mattn/go-sqlite3"
)

var driverName string = "sqlite3"

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//Database
	db := initDB("storage.db")
	migrate(db)

	//Routes
	e.GET("/", func(c echo.Context) error {
		return c.JSON(200, "Index page")
	})

	e.GET("/posts", handlers.GetPosts(db))

	//Start server
	e.Logger.Fatal(e.Start(":9090"))
}

func initDB(filepath string) *sql.DB {
	db, err := sql.Open(driverName, filepath)

	if err != nil {
		panic(err)
	}

	if db == nil {
		panic("db nil")
	}

	return db
}

func migrate(db *sql.DB) {
	query := `
            CREATE TABLE IF NOT EXISTS posts(
                    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
                    title VARCHAR NOT NULL,
                    topic VARCHAR NOT NULL,
                    votes INTEGER NOT NULL
            );

            INSERT INTO posts(title, topic, votes) VALUES('Chapter 1','Introduction', 30);
            INSERT INTO posts(title, topic, votes) VALUES('Chapter 1','Methods', 12);
            INSERT INTO posts(title, topic, votes) VALUES('Chapter 2','Methods', 17);
            INSERT INTO posts(title, topic, votes) VALUES('Chapter 1','Concurrency', 4);
            INSERT INTO posts(title, topic, votes) VALUES('Chapter 2','Concurrency', 10);
       `
	_, err := db.Exec(query)

	if err != nil {
		panic(err)
	}
}
