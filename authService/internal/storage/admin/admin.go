package admin

import (
	"authService/internal/db"
	"authService/internal/models"
	"github.com/google/uuid"
)

func GetAdmin(u map[string]string) (string, error) {
	var admin *models.Admin
	db.DB().Where("login = ? and password = ?", u["login"], u["password"]).First(&admin)

	err := uuid.Validate(admin.Uid)

	return admin.Uid, err
}
