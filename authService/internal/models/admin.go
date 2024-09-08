package models

import (
	"authService/internal/db"
)

type Admin struct {
	Uid      string `json:"uid" gorm:"primaryKey"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Username string `json:"username"`
}

func (a *Admin) TableName() string {
	return "admins"
}

func AdminMigrate() {
	err := db.DB().AutoMigrate(&Admin{})
	if err != nil {
		panic(err)
	}
}

func init() {
	db.Add(AdminMigrate)
}
