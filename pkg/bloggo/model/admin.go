package model

type Admin struct {
	ID          uint64 `gorm:"column=id;primaryKey"`
	Password    string `gorm:"column=password"`
	Salt        string `gorm:"column=salt"`
	Email       string `gorm:"column=email; unique"`
	Description string `gorm:"column=description"`
	AvatarUrl   string `gorm:"column=avatar_url"`
}