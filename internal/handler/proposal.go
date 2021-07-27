package handler

import (
	"net/http"

	"github.com/gosimple/slug"
	"github.com/labstack/echo/v4"
)

type Proposal struct {
	Author  string   `json:"author"`
	Title   string   `json:"title"`
	Body    string   `json:"body"`
	Choices []string `json:"choices"`
	Created int64    `json:"created"`
	Start   int64    `json:"start"`
	End     int64    `json:"end"`
}

type SignedProposalIFPS struct {
	ID        string `json:"id"`
	Signature string `json:"signature"`
	Proposal
}

func (h handler) PostProposal(c echo.Context) error {
	collection := h.db.Client().Database("etb-voting-app").Collection("proposals")

	var signedProposal struct {
		Signature string   `json:"signature"`
		Proposal  Proposal `json:"proposal"`
	}

	if err := c.Bind(&signedProposal); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	name := slug.MakeLang(signedProposal.Proposal.Title, "en")
	hash, err := h.ipfs.PinJSON("etb/"+name, signedProposal)
	if err != nil {
		//log.Fatal(err) // TODO
		return c.JSON(http.StatusInternalServerError, err)
	}

	var final SignedProposalIFPS
	final.ID = hash
	final.Signature = signedProposal.Signature
	final.Proposal = signedProposal.Proposal

	_, err = collection.InsertOne(c.Request().Context(), final)
	if err != nil {
		//log.Fatal(err) // TODO
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.String(http.StatusOK, hash)
}
