package repository

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pluveto/bloggo/pkg/bloggo"
)

type Repository struct {
	dbConn *sql.DB
}

func New(c *bloggo.Config) *Repository {
	var err error

	var repo = &Repository{}

	var wd, _ = os.Getwd()
	repo.dbConn, err = sql.Open("sqlite3", wd+"/tmp.db")
	checkErr(err)
	fmt.Println("Repo is ready.")
	return repo
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
