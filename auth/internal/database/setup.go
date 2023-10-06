package database

import (
	"auth/config"
	"database/sql"
	"fmt"
	"log"
	"net/url"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"
)

func Setup() (*sql.DB, error) {
	// configs, err := config.LoadPostgresConfig("./config")
	configs, err := config.LoadPostgresConfig()
	if err != nil {
		fmt.Println("error trying to load config variables", err)
		return nil, err
	}
	// construct the conn string
	dsn := url.URL{
		Scheme: "postgres",
		Host:   configs.Postgresdb.Host,
		User:   url.UserPassword(configs.Postgresdb.User, configs.Postgresdb.Password),
		Path:   configs.Postgresdb.Dbname,
	}

	q := dsn.Query()
	q.Add("sslmode", configs.Postgresdb.Sslmode)

	dsn.RawQuery = q.Encode()

	log.Println(dsn.String())
	db, err := sql.Open("pgx", dsn.String())
	if err != nil {
		fmt.Println("error trying to open a postgres connection", err)
		return nil, err
	}

	log.Println("connecting to a database successfully ..")
	return db, nil
}
