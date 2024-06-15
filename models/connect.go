package models

import (
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func Connect() (*sqlx.DB, error) {
	conf := mysql.Config{
		User:                 os.Getenv("NS_MARIADB_USER"),
		Passwd:               os.Getenv("NS_MARIADB_PASSWORD"),
		Net:                  "tcp",
		Addr:                 os.Getenv("NS_MARIADB_HOSTNAME") + ":" + os.Getenv("NS_MARIADB_PORT"),
		DBName:               os.Getenv("NS_MARIADB_DATABASE"),
		AllowNativePasswords: true,
	}

	db, err := sqlx.Open("mysql", conf.FormatDSN())
	if err != nil {
		log.Println("failed to open DB connection")

		return db, err
	}

	err = db.Ping()
	if err != nil {
		log.Println("DB ping failed")

		return db, err
	}

	log.Println("DB ping succeeded")

	return db, nil
}
