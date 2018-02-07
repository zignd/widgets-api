package services

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/pkg/errors"
)

func OpenDbConn(connStr string) (*gorm.DB, error) {
	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return db, nil
}
