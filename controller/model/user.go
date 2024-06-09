package model

type User struct {
	Id       uint
	Name     string
	Email    string `gorm:"unique"`
	Password []byte
}

type UserResponse struct {
	Id    uint
	Name  string
	Email string
}