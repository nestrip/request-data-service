package db

import (
	"database/sql"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/nestrip/request-data-service/ent"
	"log"
	"os"
)

var Client *ent.Client

func Open(databaseUrl string) *ent.Client {
	db, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv))
}

func Connect() {
	Client = Open(os.Getenv("DATABASE_URL"))
	fmt.Println("Connected to database")
}
