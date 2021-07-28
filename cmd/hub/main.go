package main

import (
	"context"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/bson"
	"voting-hub/internal/db"
	"voting-hub/internal/handler"
	"voting-hub/internal/ipfs"

	log "github.com/sirupsen/logrus"
)

func main() {
	ctx := context.Background()

	mongoURI := os.Getenv("MONGO_URI")
	pinataApiKey := os.Getenv("PINATA_API_KEY")
	pinataSecretKey := os.Getenv("PINATA_SECRET_KEY")

	dbConn, err := db.NewClient(ctx, db.Options{URI: mongoURI})
	if err != nil {
		log.Fatalf("db connect: %v", err)
	}
	defer func() {
		if err = dbConn.Disconnect(ctx); err != nil {
			log.Fatalf("db disconnect: %v", err)
		}
	}()

	// Replace space with last info
	collection := dbConn.Client().Database("etb-voting-app").Collection("space")
	_, _ = collection.DeleteMany(ctx, bson.D{})
	_, _ = collection.InsertOne(ctx, bson.D{
		{"name", "Eat The Blocks"},
		{"token", "ETB"},
		{"address", "0x7ac64008fa000bfdc4494e0bfcc9f4eff3d51d2a"},
		{"network", "56"},
		{"admins", []string{
			"0x4E48C12cf0ABEf413A2E8994B4A6a743C3f2d296",
			"0x539dea3fe88d32b46dd6e07b9ad99c59e81d85be",
		}},
	})

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
	e.POST("/vote", h.PostVote)
	e.POST("/score", h.GetScore)

	log.Fatalf("server start: %v", e.Start(":80"))
}
