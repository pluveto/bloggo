package model

/*
 * @autocode({"table": "admins", "json": "admin", "label": "管理员"})
 */
type Admin struct {
	ID          int64  `json:"id"          gorm:"column=id;primaryKey" auto:"weight=0;searchable=false;sortable=true;mode=r;  "`
	Username    string `json:"username"    gorm:"column=username"      auto:"weight=0;searchable=false;sortable=true;mode=rw;required=1;label=用户名  ;type=text"`
	ScreenName  string `json:"screenName"  gorm:"column=screen_name"   auto:"weight=1;searchable=true ;sortable=true;mode=rw;required=1;label=昵称   ;type=text"`
	Password    string `json:"password"    gorm:"column=password"      auto:"weight=4;searchable=false;sortable=true;mode=w; required=1;label=密码   ;type=password"`
	Salt        string `json:"salt"        gorm:"column=salt"          auto:"weight=0;searchable=false;sortable=true;mode=i; required=1;                   "`
	Email       string `json:"email"       gorm:"column=email; unique" auto:"weight=3;searchable=false;sortable=true;mode=rw;required=1;label=邮箱   ;type=email"      binding:"email,required,min=6,max=255"`
	Description string `json:"description" gorm:"column=description"   auto:"weight=9;searchable=false;sortable=true;mode=rw;required=0;label=个人介绍;type=textarea"  binding:"printascii,required"`
	AvatarUrl   string `json:"avatarUrl"   gorm:"column=avatar_url"    auto:"weight=2;searchable=false;sortable=true;mode=rw;required=1;label=头像链接;type=url"`
}
