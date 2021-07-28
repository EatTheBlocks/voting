package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"voting-hub/internal/eth"
)

type Vote struct {
	Author    string `json:"author"`
	Proposal  string `json:"proposal"`
	Choice    int    `json:"choice"`
	Timestamp int64  `json:"timestamp"`
}

type SignedVoteIFPS struct {
	ID        string `json:"id"`
	Signature string `json:"signature"`
	Vote
}

type VoteIFPS struct {
	ID string `json:"id"`
	Vote
}

func (h handler) PostVote(c echo.Context) error {
	ctx := c.Request().Context()
	collection := h.db.Client().Database("etb-voting-app").Collection("votes")

	var signedVote struct {
		Signature string `json:"signature"`
		Vote      Vote   `json:"vote"`
	}

	if err := c.Bind(&signedVote); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	msg, err := json.Marshal(signedVote.Vote)
	if err != nil {
		log.Error(errors.Wrap(err, "PostVote decode json vote"))
		return c.JSON(http.StatusInternalServerError, err)
	}

	if !eth.VerifySig(signedVote.Vote.Author, signedVote.Signature, msg) {
		return c.JSON(http.StatusBadRequest, "invalid signature")
	}

	collectionProposals := h.db.Client().Database("etb-voting-app").Collection("proposals")

	var proposal SignedProposalIFPS
	err = collectionProposals.FindOne(ctx, bson.D{{"id", signedVote.Vote.Proposal}}).Decode(&proposal)
	if err == mongo.ErrNoDocuments {
		return c.JSON(http.StatusNotFound, "proposal not found")
	} else if err != nil {
		log.Error(errors.Wrap(err, "PostVote check vote"))
		return c.JSON(http.StatusInternalServerError, err)
	}

	if time.Now().Unix() > proposal.End {
		return c.JSON(http.StatusBadRequest, "proposal closed")
	}

	err = collection.FindOne(ctx, bson.D{
		{"vote.author", signedVote.Vote.Author},
		{"vote.proposal", signedVote.Vote.Proposal},
	}).Err()
	if err != mongo.ErrNoDocuments {
		return c.JSON(http.StatusBadRequest, "you have already voted for this proposal")
	}

	// TODO verify user power based on proposal.snapshot

	hash, err := h.ipfs.PinJSON("etb/vote/"+signedVote.Vote.Proposal+"/"+signedVote.Vote.Author, signedVote)
	if err != nil {
		log.Error(errors.Wrap(err, "PostVote ipfs pin vote"))
		return c.JSON(http.StatusInternalServerError, err)
	}

	var final SignedVoteIFPS
	final.ID = hash
	final.Signature = signedVote.Signature
	final.Vote = signedVote.Vote

	_, err = collection.InsertOne(ctx, final)
	if err != nil {
		log.Error(errors.Wrap(err, "PostVote store vote"))
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.String(http.StatusOK, hash)
}
