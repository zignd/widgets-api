package business

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/zignd/widgets-api/models"
)

func GetWidgets(db *gorm.DB) ([]*models.Widget, error) {
	var widgets []*models.Widget
	if err := db.Preload("Color").Find(&widgets).Error; err != nil {
		return nil, errors.Wrap(err, "failed to query for widgets")
	}
	return widgets, nil
}

func GetWidgetByID(db *gorm.DB, id int) (*models.Widget, error) {
	var widget models.Widget
	if err := db.Preload("Color").Where(`id = ?`, id).First(&widget).Error; err != nil {
		return nil, errors.Wrap(err, "failed to query for one widget")
	}
	return &widget, nil
}
