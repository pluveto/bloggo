package repository

import (
	"fmt"
	"os"

	"github.com/pluveto/bloggo/pkg/bloggo"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}



func New(c *bloggo.Config) *Repository {
	var err error

	var repo = &Repository{}

	var wd, _ = os.Getwd()
	repo.db, err = gorm.Open(sqlite.Open(wd+"/tmp.db"), &gorm.Config{})

	checkErr(err)
	fmt.Println("Repo is ready.")
	return repo
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
