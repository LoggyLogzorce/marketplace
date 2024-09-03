package user

import (
	"adminService/internal/db"
	"github.com/google/uuid"
)

type User struct {
	Uid      uuid.UUID `json:"uid" gorm:"primaryKey"`
	Login    string    `json:"login"`
	Password string    `json:"password"`
	Username string    `json:"username"`
	Role     uint8     `json:"role"`
}

func (u *User) TableName() string {
	return "admins"
}

func MigrateUser() {
	err := db.DB().AutoMigrate(User{})
	if err != nil {
		panic(err)
	}
}

func init() {
	db.Add(MigrateUser)
	uuid.New()
}
