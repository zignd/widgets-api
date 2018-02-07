package business

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/zignd/widgets-api/models"
)

func GetUsers(db *gorm.DB) ([]*models.User, error) {
	var users []*models.User
	if err := db.Find(&users).Error; err != nil {
		return nil, errors.Wrap(err, "failed to query for users")
	}
	return users, nil
}

func GetUserByID(db *gorm.DB, id int) (*models.User, error) {
	var user models.User
	if err := db.Where(`id = ?`, id).First(&user).Error; err != nil {
		return nil, errors.Wrap(err, "failed to query for one user")
	}
	return &user, nil
}
