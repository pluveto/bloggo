package model

type Admin struct {
	ID          uint64
	Username    string
	Password    string
	Salt        string
	Email       string
	Description string
	AvatarUrl   string
}

type AdminLoginReq struct {
	Email    string `form:"email"    json:"email"    validate:"email,required,min=6,max=255"`
	Password string `form:"password" json:"password" validate:"printascii,required"`
	Captcha  string `form:"captcha"  json:"captcha"  validate:"len=6"`
}

type AdminLoginRet struct {
}

type AdminCreateReq struct {
	Email    string `form:"email"    json:"email"    validate:"email,required,min=6,max=255"`
	Password string `form:"password" json:"password" validate:"printascii,required"`
}

type AdminCreateRet struct {
	ID string `json:"id"`
}
