package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Space struct {
	Name    string   `json:"name"`
	Token   string   `json:"token"`
	Address string   `json:"address"`
	Network string   `json:"network"`
	Admins  []string `json:"admins"`
}

func (h handler) GetSpace(c echo.Context) error {
	ctx := c.Request().Context()
	collection := h.db.Client().Database("etb-voting-app").Collection("space")

	var space Space
	err := collection.FindOne(ctx, bson.D{}).Decode(&space)
	if err == mongo.ErrNoDocuments {
		return c.JSON(http.StatusNotFound, "space not found")
	} else if err != nil {
		log.Error(errors.Wrap(err, "GetSpace find space"))
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, space)
}
