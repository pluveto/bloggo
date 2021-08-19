package api

import (
	"github.com/gin-gonic/gin"
	"github.com/pluveto/bloggo/pkg/bloggo"
	"github.com/pluveto/bloggo/pkg/bloggo/model"
	"github.com/pluveto/bloggo/pkg/errcode"
)

	
type PostCreateReq struct {
	Path    string `form:"path" json:"path" binding:"required"`
	Slug    string `form:"slug" json:"slug" binding:"required"`
	Title    string `form:"title" json:"title" binding:"required"`
	Content    string `form:"content" json:"content" binding:"required"`
}

type PostCreateRet struct {
	ID    int64 
	Slug    string 
}

func (api *API) PostCreate(c *gin.Context) {
	var (
		req PostCreateReq
		err error = c.ShouldBindJSON(&req)
	)

	if err != nil {
		api.retBadParam(c, err)
		return
	}
	
	
	var existingPost *model.Post		
	
	existingPost, _ = api.Service.PostGetBySlug(req.Slug)
	if existingPost != nil {
		api.RetJSON(c, nil, errcode.ApiError(bloggo.ErrConflictPost))
		return
	}		
	newPost, err := api.Service.PostCreate(
			req.Path,
			req.Slug,
			req.Title,
			req.Content,
	)
	if err != nil {
		api.RetJSON(c, nil, err)
		return
	}
	var retModel = PostCreateRet{ID: newPost.ID}	
	api.RetJSON(c, retModel, err)
}

type PostGetReq struct {
	Slug string `form:"slug" json:"slug" binding:"required:true"`
}

type PostGetRet struct {
	Path        string `form:"path" json:"path" binding:"required:true"`
	Slug        string `form:"slug" json:"slug" binding:"required:true"`
	Title       string `form:"title" json:"title" binding:"required:true"`
	Content     string `form:"content" json:"content" binding:"required:true"`
	PublishedAt int64  `form:"publishedAt" json:"publishedAt" binding:"required:true"`
	ModifiedAt  int64  `form:"modifiedAt" json:"modifiedAt" binding:"required:true"`
}

func (api *API) PostGet(c *gin.Context) {
	var (
		req PostCreateReq
		err error = c.ShouldBindJSON(&req)
	)

	if err != nil {
		api.retBadParam(c, err)
		return
	}
	
	
	var existingPost *model.Post		
	
	existingPost, _ = api.Service.PostGetBySlug(req.Slug)
	if existingPost != nil {
		api.RetJSON(c, nil, errcode.ApiError(bloggo.ErrConflictPost))
		return
	}		
	newPost, err := api.Service.PostCreate(
			req.Path,
			req.Slug,
			req.Title,
			req.Content,
	)
	if err != nil {
		api.RetJSON(c, nil, err)
		return
	}
	var retModel = PostCreateRet{ID: newPost.ID}	
	api.RetJSON(c, retModel, err)
}


type PostGetRecentReq struct {
	Slug string `form:"slug" json:"slug" binding:"required:true"`
}

type PostGetRecentRet struct {
	Path        string `form:"path" json:"path" binding:"required:true"`
	Slug        string `form:"slug" json:"slug" binding:"required:true"`
	Title       string `form:"title" json:"title" binding:"required:true"`
	Content     string `form:"content" json:"content" binding:"required:true"`
	PublishedAt int64  `form:"publishedAt" json:"publishedAt" binding:"required:true"`
	ModifiedAt  int64  `form:"modifiedAt" json:"modifiedAt" binding:"required:true"`
}

type PostListReq struct {
}

type PostListRet struct {
	Path        string `form:"path" json:"path" binding:"required:true"`
	Slug        string `form:"slug" json:"slug" binding:"required:true"`
	Title       string `form:"title" json:"title" binding:"required:true"`
	Content     string `form:"content" json:"content" binding:"required:true"`
	PublishedAt int64  `form:"publishedAt" json:"publishedAt" binding:"required:true"`
	ModifiedAt  int64  `form:"modifiedAt" json:"modifiedAt" binding:"required:true"`
}

type PostUpdateReq struct {
	Path    string `form:"path" json:"path" binding:"required:true"`
	Slug    string `form:"slug" json:"slug" binding:"required:true"`
	Title   string `form:"title" json:"title" binding:"required:true"`
	Content string `form:"content" json:"content" binding:"required:true"`
}

type PostUpdateRet struct {
}

type PostSearchReq struct {
}

type PostSearchRet struct {
	Path        string `form:"path" json:"path" binding:"required:true"`
	Slug        string `form:"slug" json:"slug" binding:"required:true"`
	Title       string `form:"title" json:"title" binding:"required:true"`
	Content     string `form:"content" json:"content" binding:"required:true"`
	PublishedAt int64  `form:"publishedAt" json:"publishedAt" binding:"required:true"`
	ModifiedAt  int64  `form:"modifiedAt" json:"modifiedAt" binding:"required:true"`
}

type PostEnableReq struct {
	Slug string `form:"slug" json:"slug" binding:"required:true"`
}

type PostEnableRet struct {
	Path        string `form:"path" json:"path" binding:"required:true"`
	Slug        string `form:"slug" json:"slug" binding:"required:true"`
	Title       string `form:"title" json:"title" binding:"required:true"`
	Content     string `form:"content" json:"content" binding:"required:true"`
	PublishedAt int64  `form:"publishedAt" json:"publishedAt" binding:"required:true"`
	ModifiedAt  int64  `form:"modifiedAt" json:"modifiedAt" binding:"required:true"`
}

type PostDisableReq struct {
	Slug string `form:"slug" json:"slug" binding:"required:true"`
}

type PostDisableRet struct {
	Path        string `form:"path" json:"path" binding:"required:true"`
	Slug        string `form:"slug" json:"slug" binding:"required:true"`
	Title       string `form:"title" json:"title" binding:"required:true"`
	Content     string `form:"content" json:"content" binding:"required:true"`
	PublishedAt int64  `form:"publishedAt" json:"publishedAt" binding:"required:true"`
	ModifiedAt  int64  `form:"modifiedAt" json:"modifiedAt" binding:"required:true"`
}
