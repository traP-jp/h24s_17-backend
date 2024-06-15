package main

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type TableList struct {
	Tables_in_h24s17 string `json:"tables_in_h24s17,omitempty" db:"Tables_in_h24s17"`
}

func main() {
	conf := mysql.Config{
		User:                 os.Getenv("MYSQL_USER"),
		Passwd:               os.Getenv("MYSQL_PASSWORD"),
		Net:                  "tcp",
		Addr:                 os.Getenv("MYSQL_HOSTNAME") + ":" + os.Getenv("MYSQL_PORT"),
		DBName:               "h24s17",
		AllowNativePasswords: true,
	}

	db, err := sqlx.Open("mysql", conf.FormatDSN())
	if err != nil {
		log.Fatal(err)

		return
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		log.Println("ping failed")

		return
	} else {
		log.Println("ping succeeded")
	}

	var x TableList
	err = db.Get(&x, "SHOW TABLES")

	log.Println(x)

	if errors.Is(err, sql.ErrNoRows) {
		log.Printf("no table in %s\n", conf.DBName)
	}
	if err != nil {
		log.Fatalf("DB error: %s\n", err)

		return
	}

	e := echo.New()

	e.GET("/hello/:name", func(c echo.Context) error {
		name := c.Param("name")

		return c.String(http.StatusOK, "Hello, "+name+".\n")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
