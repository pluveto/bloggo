package model

type Setting struct {
	Key   string
	Value string
	Type  string
}

// == USER ==
/*
type User struct {
	ID          int64
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
type Comment struct {
	ID          int64
	PostID      int64
	Content     string
	AuthorName  string
	AuthorEmail string
	AuthorUrl   string
}
