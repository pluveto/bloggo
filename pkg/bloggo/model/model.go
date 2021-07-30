package model

type Setting struct {
	Key   string
	Value string
	Type  string
}

// == USER ==
/*
type User struct {
	ID          uint64
	Username    string
	Password    string
	Salt        string
	Email       string
	Description string
	AvatarUrl   string
}
*/

/*
type UserServicer interface {
	Login(username string, password string) *Token
	Logout(tokenValue string)
	Create(user *User)
}

*/

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
