type AuthGetUserInfoRet struct {
	Username    string `json:"username"`
	ScreenName  string `json:"screenName"`
	ID          int64  `json:"id"`
	Email       string `json:"email"`
	Description string `json:"description"`
	AvatarUrl   string `json:"avatarUrl"`
}