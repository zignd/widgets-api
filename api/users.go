package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/zignd/widgets-api/business"
	"github.com/zignd/widgets-api/services"
	"github.com/zignd/widgets-api/viewmodels"
)

func GetUsers(c *gin.Context) {
	connStr := c.GetString("ConnStr")
	db, err := services.OpenDbConn(connStr)
	if err != nil {
		log.Error(errors.Wrap(err, "failed to open a database connection"))
		c.AbortWithStatusJSON(500, &viewmodels.Error{
			Message: viewmodels.GenericError,
		})
		return
	}
	defer db.Close()

	users, err := business.GetUsers(db)
	if err != nil {
		log.Error(errors.Wrap(err, "failed to get the users"))
		c.AbortWithStatusJSON(500, &viewmodels.Error{
			Message: viewmodels.GenericError,
		})
		return
	}

	vm := make([]*viewmodels.User, len(users))
	for i, user := range users {
		vm[i] = &viewmodels.User{
			ID:        user.ID,
			Name:      user.Name,
			Avatar:    user.Avatar,
			Username:  user.Username,
			CreatedAt: user.CreatedAt,
		}
	}

	c.JSON(200, vm)
}

func GetUserByID(c *gin.Context) {
	connStr := c.GetString("ConnStr")
	db, err := services.OpenDbConn(connStr)
	if err != nil {
		log.Error(errors.Wrap(err, "failed to open a database connection"))
		c.AbortWithStatusJSON(500, &viewmodels.Error{
			Message: viewmodels.GenericError,
		})
		return
	}
	defer db.Close()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(400, &viewmodels.Error{
			Message: "id is supposed to be an integer",
		})
		return
	}

	user, err := business.GetUserByID(db, id)
	if err != nil {
		log.Error(errors.Wrapf(err, "failed to get a user with the id %d", id))
		c.AbortWithStatusJSON(500, &viewmodels.Error{
			Message: viewmodels.GenericError,
		})
		return
	}

	vm := &viewmodels.User{
		ID:        user.ID,
		Name:      user.Name,
		Avatar:    user.Avatar,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
	}

	c.JSON(200, vm)
}
