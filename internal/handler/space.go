package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Space struct {
	Name    string   `json:"name"`
	Address string   `json:"address"`
	Network string   `json:"network"`
	Admins  []string `json:"admins"`
}

func (h handler) GetSpace(c echo.Context) error {
	collection := h.db.Client().Database("etb-voting-app").Collection("space")

	var space Space

	err := collection.FindOne(c.Request().Context(), bson.D{}).Decode(&space)
	if err == mongo.ErrNoDocuments {
		return c.JSON(http.StatusNotFound, "space not found")
	} else if err != nil {
		//log.Fatal(err) // TODO
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, space)
}
