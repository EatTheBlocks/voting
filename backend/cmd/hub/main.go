package main

import (
	"context"
	"os"
	"strings"

	"github.com/ijustfool/docker-secrets"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"voting-hub/internal/db"
	"voting-hub/internal/handler"
	"voting-hub/internal/ipfs"
)

func main() {
	dockerSecrets, _ := secrets.NewDockerSecrets("")

	log.Info("starting ETB voting hub")

	ctx := context.Background()

	mongoURI, _ := dockerSecrets.Get("MONGO_URI")
	pinataApiKey, _ := dockerSecrets.Get("PINATA_API_KEY")
	pinataSecretKey, _ := dockerSecrets.Get("PINATA_SECRET_KEY")
	rpcApi, _ := dockerSecrets.Get("RPC_API")
	multicallAddress, _ := dockerSecrets.Get("MULTICALL_ADDRESS")
	ETBTokenAddress, _ := dockerSecrets.Get("ETB_TOKEN_ADDRESS")
	adminList, _ := dockerSecrets.Get("ADMIN_LIST")
	port := os.Getenv("PORT")

	if port == "" {
		port = "80"
	}

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

	log.Infof("server start on port %s", port)

	log.Fatalf("server start: %v", e.Start(":"+port))
}
