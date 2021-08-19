package model

/*
 * @autocode({"table": "posts", "json": "post", "label": "文章"})
 */
type Post struct {
	ID          int64  `json:"id"          gorm:"column=id;primaryKey" auto:"weight=0;searchable=false;sortable=true ;mode=r;required=true; type=text;label=ID"`
	AuthorID    int64  `json:"authorID"    gorm:"column=author_id"     auto:"weight=1;searchable=false;sortable=false;mode=i;required=true; type=text;label=作者ID"`
	Path        string `json:"path"        gorm:"column=path"          auto:"weight=2;searchable=false;sortable=false;mode=rw;required=true;type=text;label=路径"`
	Slug        string `json:"slug"        gorm:"column=slug;unique"   auto:"weight=3;searchable=false;sortable=false;mode=rw;required=true;type=text;label=网址别名"`
	Title       string `json:"title"       gorm:"column=title"         auto:"weight=4;searchable=true ;sortable=false;mode=rw;required=true;type=text;label=标题"`
	Content     string `json:"content"     gorm:"column=content"       auto:"weight=5;searchable=true ;sortable=false;mode=rw;required=true;type=text;label=内容"`
	PublishedAt int64  `json:"publishedAt" gorm:"column=published_at"  auto:"weight=6;searchable=false;sortable=true ;mode=r; required=true;type=text;label=发布于"`
	ModifiedAt  int64  `json:"modifiedAt"  gorm:"column=modified_at"   auto:"weight=7;searchable=false;sortable=true ;mode=r; required=true;type=text;label=编辑于"`
}
