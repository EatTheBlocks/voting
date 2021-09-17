package main

import (
	"context"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/acme/autocert"
	"voting-hub/internal/db"
	"voting-hub/internal/handler"
	"voting-hub/internal/ipfs"
)

func main() {
	log.Info("starting ETB voting hub")

	ctx := context.Background()

	mongoURI := os.Getenv("MONGO_URI")
	pinataApiKey := os.Getenv("PINATA_API_KEY")
	pinataSecretKey := os.Getenv("PINATA_SECRET_KEY")
	rpcApi := os.Getenv("RPC_API")
	multicallAddress := os.Getenv("MULTICALL_ADDRESS")
	ETBTokenAddress := os.Getenv("ETB_TOKEN_ADDRESS")
	adminList := os.Getenv("ADMIN_LIST")
	email := os.Getenv("EMAIL")
	domain := os.Getenv("DOMAIN")

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
	_, err = collection.InsertOne(ctx, bson.D{
		{"name", "Eat The Blocks"},
		{"token", "ETB"},
		{"address", ETBTokenAddress},
		{"network", "56"},
		{"admins", strings.Split(adminList, ",")},
	})
	if err != nil {
		log.Warn(err)
	}

	pinata := ipfs.NewPinataClient(pinataApiKey, pinataSecretKey)

	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	e.AutoTLSManager.Email = email
	e.AutoTLSManager.HostPolicy = autocert.HostWhitelist(domain)
	e.AutoTLSManager.Cache = autocert.DirCache("/cache")

	e.Pre(middleware.HTTPSRedirect())
	e.Use(middleware.CORS())

	e.Static("/", "static")

	h := handler.NewHandler(dbConn, pinata, rpcApi, multicallAddress, ETBTokenAddress)

	api := e.Group("/api")
	api.GET("", h.GetSpace)
	api.POST("/proposal", h.PostProposal)
	api.GET("/proposals", h.GetProposals)
	api.GET("/proposal/:id", h.GetProposal)
	api.POST("/vote", h.PostVote)
	api.POST("/score", h.GetScore)

	log.Info("server start on port 80 and 443")

	go func() {
		log.Errorf("server start: %v", e.Start(":80"))
	}()

	log.Fatalf("server start: %v", e.StartAutoTLS(":443"))
}
