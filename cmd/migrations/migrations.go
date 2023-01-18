package main

import (
	"flag"
	"github.com/MatteoMiotello/goAccounting/internal/bootstrap"
	"log"
	"os"

	_ "github.com/MatteoMiotello/goAccounting/migrations"
	_ "github.com/lib/pq"

	"github.com/pressly/goose/v3"
)

var (
	flags = flag.NewFlagSet("goose", flag.ExitOnError)
	dir   = flags.String("dir", "./migrations", "directory with migration files")
)

func init() {
	bootstrap.InitConfig()
	bootstrap.InitDbMigration()
}

func main() {
	flags.Parse(os.Args[1:])
	args := flags.Args()

	command := args[0]

	dbstring := "postgres://postuser@localhost/postgres?port=1234&password=root&sslmode=disable"

	db, err := goose.OpenDBWithDriver("postgres", dbstring)
	if err != nil {
		log.Fatalf("goose: failed to open DB: %v\n", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("goose: failed to close DB: %v\n", err)
		}
	}()

	arguments := []string{}
	if len(args) > 3 {
		arguments = append(arguments, args[3:]...)
	}

	if err := goose.Run(command, db, *dir, arguments...); err != nil {
		log.Fatalf("goose %v: %v", command, err)
	}
}
