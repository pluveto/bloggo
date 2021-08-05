package model

/*
 * @autocode({"table": "admins", "json": "admin"})
 */
type Post struct {
	ID          int64  `json:"iD" gorm:"column=id" auto:"weight=0;searchable=false;sortable=false;mode=i;required=true;label=ID;type=text"`
	AuthorID    int64  `json:"authorID" gorm:"column=author_id" auto:"weight=1;searchable=false;sortable=false;mode=i;required=true;label=AuthorID;type=text"`
	Path        string `json:"path" gorm:"column=path" auto:"weight=2;searchable=false;sortable=false;mode=rw;required=true;label=路径;type=text"`
	Slug        string `json:"slug" gorm:"column=slug" auto:"weight=3;searchable=false;sortable=false;mode=rw;required=true;label=鼻涕虫;type=text"`
	Title       string `json:"title" gorm:"column=title" auto:"weight=4;searchable=false;sortable=false;mode=rw;required=true;label=标题;type=text"`
	Content     string `json:"content" gorm:"column=content" auto:"weight=5;searchable=false;sortable=false;mode=rw;required=true;label=内容;type=text"`
	PublishedAt int64  `json:"publishedAt" gorm:"column=published_at" auto:"weight=6;searchable=false;sortable=false;mode=rw;required=true;label=PublishedAt;type=text"`
	ModifiedAt  int64  `json:"modifiedAt" gorm:"column=modified_at" auto:"weight=7;searchable=false;sortable=false;mode=rw;required=true;label=ModifiedAt;type=text"`
}
