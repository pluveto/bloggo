package service

import (
	"github.com/pluveto/bloggo/pkg/bloggo"
	"github.com/pluveto/bloggo/pkg/bloggo/repository"
)

// ISettingService mplements SettingServicer
type Service struct {
	repo *repository.Repository
}

func New(c *bloggo.Config) *Service {
	var srv = &Service{}
	srv.repo = repository.New(c)
	return srv
}
