package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"

	"github.com/younesious/go-microservices/authentication/data"
)

const webPort = "8083"

type Config struct {
	Repo   data.Repository
	Client *http.Client
}

func main() {
	log.Println("---------------------------------------------")
	log.Println("Attempting to connect to Postgres...")

	conn := connectToDB()
	if conn == nil {
		log.Panic("can't connect to postgres!")
	}

	app := Config{
		Client: &http.Client{},
	}
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	log.Printf("Starting authentication end service on port %s\n", webPort)
	err := srv.ListenAndServe()

	if err != nil {
		log.Panic(err)
	}
}

func connectToDB() *sql.DB {
	dsn := os.Getenv("DSN")
	count := 0
	for {
		connection, err := openDB(dsn)
		if err != nil {
			log.Println("Postgres not ready...")
			count++
		} else {
			log.Println("Connected to database!")
			return connection
		}

		if count > 10 {
			log.Println(err)
			return nil
		}

		log.Println("Backing off for two seconds...")
		time.Sleep(2 * time.Second)
		continue
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (app *Config) setupRepo(conn *sql.DB) {
	db := data.NewPostgresRepository(conn)
	app.Repo = db
}
