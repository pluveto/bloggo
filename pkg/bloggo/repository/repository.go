package repository

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/pluveto/bloggo/pkg/bloggo"
	"github.com/pluveto/bloggo/pkg/bloggo/model"
	"github.com/pluveto/bloggo/pkg/errcode"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Repository struct {
	db  *gorm.DB
	rdb *redis.Client
}

func New(c *bloggo.Config) *Repository {
	var err error
	var repo = &Repository{}

	wd, err := os.Getwd()
	checkFatalErr(err)

	repo.db, err = gorm.Open(sqlite.Open(wd+"/tmp.db"), &gorm.Config{})
	checkFatalErr(err)
	// TODO: 使用自定义 Logger 为数据库记录日志
	repo.runMigrate()

	repo.rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	var ctx = context.Background()
	var ret = repo.rdb.Set(ctx, "bloggo_started", time.Now().Unix(), 1)
	if ret.Err() != nil {
		panic("Redis may be not running!")
	}
	fmt.Println("[Repo] Repo is ready.")

	return repo
}

func (repo *Repository) runMigrate() {
	repo.db.AutoMigrate(&model.Admin{})
}

func checkFatalErr(err error) {
	if err != nil {
		panic(err)
	}
}

func checkDbError(err error) error {
	if err == nil {
		return nil
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errcode.ApiError(errcode.ErrNoSuchResource)
	}
	return err
}
