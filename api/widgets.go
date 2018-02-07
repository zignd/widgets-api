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

func GetWidgets(c *gin.Context) {
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

	widgets, err := business.GetWidgets(db)
	if err != nil {
		log.Error(errors.Wrap(err, "failed to get the widgets"))
		c.AbortWithStatusJSON(500, &viewmodels.Error{
			Message: viewmodels.GenericError,
		})
		return
	}

	vm := make([]*viewmodels.Widget, len(widgets))
	for i, widget := range widgets {
		vm[i] = &viewmodels.Widget{
			ID:   widget.ID,
			Name: widget.Name,
			Color: &viewmodels.Color{
				ID:   widget.Color.ID,
				Name: widget.Color.Name,
			},
			Price:     widget.Price,
			Melts:     widget.Melts,
			Inventory: widget.Inventory,
			CreatedAt: widget.CreatedAt,
		}
	}

	c.JSON(200, vm)
}

func GetWidgetByID(c *gin.Context) {
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

	widget, err := business.GetWidgetByID(db, id)
	if err != nil {
		log.Error(errors.Wrapf(err, "failed to get a widget with the id %d", id))
		c.AbortWithStatusJSON(500, &viewmodels.Error{
			Message: viewmodels.GenericError,
		})
		return
	}

	vm := &viewmodels.Widget{
		ID:   widget.ID,
		Name: widget.Name,
		Color: &viewmodels.Color{
			ID:   widget.Color.ID,
			Name: widget.Color.Name,
		},
		Price:     widget.Price,
		Melts:     widget.Melts,
		Inventory: widget.Inventory,
		CreatedAt: widget.CreatedAt,
	}

	c.JSON(200, vm)
}
