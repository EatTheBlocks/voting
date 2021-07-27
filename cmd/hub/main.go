package main

import (
	"context"
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"voting-hub/internal/db"
	"voting-hub/internal/handler"
	"voting-hub/internal/ipfs"
)

func main() {
	ctx := context.Background()

	mongoURI := os.Getenv("MONGO_URI")
	pinataApiKey := os.Getenv("PINATA_API_KEY")
	pinataSecretKey := os.Getenv("PINATA_SECRET_KEY")

	dbConn, err := db.NewClient(ctx, db.Options{URI: mongoURI})
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = dbConn.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	pinata := ipfs.NewPinataClient(pinataApiKey, pinataSecretKey)

	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	e.Use(middleware.CORS())

	h := handler.NewHandler(dbConn, pinata)

	e.GET("/", h.GetSpace)
	e.POST("/proposal", h.PostProposal)
	e.GET("/proposals", h.GetProposals)
	e.GET("/proposal/:id", h.GetProposal)

	e.Logger.Fatal(e.Start(":80"))
}
