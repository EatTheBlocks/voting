package main

import (
	"context"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"voting-hub/internal/db"
	"voting-hub/internal/handler"
	"voting-hub/internal/ipfs"
)

func main() {
	ctx := context.Background()

	mongoURI := os.Getenv("MONGO_URI")
	pinataApiKey := os.Getenv("PINATA_API_KEY")
	pinataSecretKey := os.Getenv("PINATA_SECRET_KEY")
	rpcApi := os.Getenv("RPC_API")
	multicallAddress := os.Getenv("MULTICALL_ADDRESS")
	ETBTokenAddress := os.Getenv("ETB_TOKEN_ADDRESS")
	adminList := os.Getenv("ADMIN_LIST")

	//RPC_API=https://apis.ankr.com/c0d871dd3c6d4529b01c9362a9b79e89/6106d4a3ec1d1bcc87ec72158f8fd089/binance/archive/main;MULTICALL_ADDRESS=0x1ee38d535d541c55c9dae27b12edf090c608e6fb;ETB_TOKEN_ADDRESS=0x7ac64008fa000bfdc4494e0bfcc9f4eff3d51d2a;ADMIN_LIST=0x4E48C12cf0ABEf413A2E8994B4A6a743C3f2d296,0xb063d25968b3470F584833a9bfa1F684B9950032,0x539dea3fe88d32b46dd6e07b9ad99c59e81d85be

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
		{"address", ETBTokenAddress},
		{"network", "56"},
		{"admins", strings.Split(adminList, ",")},
	})

	pinata := ipfs.NewPinataClient(pinataApiKey, pinataSecretKey)

	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	e.Use(middleware.CORS())

	h := handler.NewHandler(dbConn, pinata, rpcApi, multicallAddress, ETBTokenAddress)

	e.GET("/", h.GetSpace)
	e.POST("/proposal", h.PostProposal)
	e.GET("/proposals", h.GetProposals)
	e.GET("/proposal/:id", h.GetProposal)
	e.POST("/vote", h.PostVote)
	e.POST("/score", h.GetScore)

	log.Fatalf("server start: %v", e.Start(":80"))
}
