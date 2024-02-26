package main

type User struct {
	Name  string `gorm:"column:name" form:"name" json:"name"`
	Image string `gorm:"column:image" form:"video" json:"video"`
}

func (u *User) TableName() string {
	return "upload_image"
}
