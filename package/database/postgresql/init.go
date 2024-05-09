package postgresql

import (
	"eniqlo/config"
	"log"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

func New() *sqlx.DB {
	config := config.PostgreSQLConfig{
		Host:     os.Getenv("DB_HOST"),
		Sslmode:  os.Getenv("DB_PARAMS"),
		User:     os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Port:     os.Getenv("DB_PORT"),
		Database: os.Getenv("DB_NAME"),
	}
	dsn := config.FormatDSN()
	db, err := sqlx.Connect("pgx", dsn)

	if err != nil {
		log.Println("m=GetPool,msg=connection has failed", err)
	}

	if err != nil {
		log.Fatalln(err)
		return nil
	}

	db.SetMaxOpenConns(60)
	db.SetMaxIdleConns(60)

	return db
}
