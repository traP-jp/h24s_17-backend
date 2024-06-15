package main

import (
	"database/sql"
	"errors"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/traP-jp/h24s_17-backend/routes"
)

type TableList struct {
	TablesInH24s17 string `json:"tables_in_h24s17,omitempty" db:"Tables_in_h24s17"`
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
		log.Println("ping failed")
		log.Fatal(err)

		return
	}

	log.Println("ping succeeded")

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
	routes.SetupRoutes(e)
	e.Logger.Fatal(e.Start(":1323"))
}
