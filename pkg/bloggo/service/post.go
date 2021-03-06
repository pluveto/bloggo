package service

import (
	"github.com/pluveto/bloggo/pkg/bloggo/model"
)


func (s *Service) PostGetByID(id int64) (model *model.Post, err error) {
    return s.repo.PostGetByID(id)
}
func (s *Service) PostGetBySlug(slug string) (model *model.Post, err error) {
    return s.repo.PostGetBySlug(slug)
}

	

    // PostCreate 创建一个帖子
    func (s *Service) PostCreate(
			path string,
			slug string,
			title string,
			content string,
    ) (ret *model.Post, err error) {
        var model = &model.Post{
				Path: path,
				Slug: slug,
				Title: title,
				Content: content,
        }
        err = s.repo.PostCreate(model)
        if err != nil {
            return
        }
        ret = model
        return
    }
	
	

	

	

	

	

	

	

	


	

	

	

	

	

	

	
