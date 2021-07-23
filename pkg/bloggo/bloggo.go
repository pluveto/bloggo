package bloggo

// == config ==

type Config struct {
}

// == Blog ==

type SiteSetting map[string]interface {
	// Name            string
	// Description     string
	// BaseUrl         string
	// Copyright       string
}

type SiteSettingServicer interface {
	Get(key string) interface{}
	GetOrDefault(key string, defaultValue interface{}) interface{}
	Set(key string, value interface{})
}

// == USER ==

type User struct {
	ID          uint64
	Username    string
	Password    string
	Email       string
	Description string
	AvatarUrl   string
}

type Token struct {
	UserID     uint64
	TokenValue string
	ExpiredAt  int64
}

type UserServicer interface {
	Login(username string, password string) *Token
	Logout(tokenValue string)
	Create(user *User)
}

// == POST ==

type Post struct {
	ID          uint64
	AuthorID    uint64
	Path        string
	Slug        string
	Title       string
	Content     string
	PublishedAt int64
	ModifiedAt  int64
}

type Comment struct {
	ID          uint64
	PostID      uint64
	Content     string
	AuthorName  string
	AuthorEmail string
	AuthorUrl   string
}
