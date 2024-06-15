package models

import (
	"database/sql"
	"errors"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type TableList struct {
	TablesInH24s17 string `json:"tables_in_h24s17,omitempty" db:"Tables_in_h24s17"`
}

func Connect() (*sqlx.DB, error) {
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
		log.Println("failed to open DB connection")

		return db, err
	}

	err = db.Ping()
	if err != nil {
		log.Println("ping failed")

		return db, err
	}

	log.Println("ping succeeded")

	var x TableList
	err = db.Get(&x, "SHOW TABLES")

	log.Println("SHOW TABLES:")
	log.Println(x)

	if errors.Is(err, sql.ErrNoRows) {
		log.Printf("no table in %s\n", conf.DBName)
	}
	if err != nil {
		log.Printf("DB error: %s\n", err)

		return db, err
	}

	return db, nil
}
