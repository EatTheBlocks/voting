package handler

import (
	"net/http"
	"strconv"

	"github.com/gosimple/slug"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Proposal struct {
	Author   string   `json:"author"`
	Title    string   `json:"title"`
	Body     string   `json:"body"`
	Choices  []string `json:"choices"`
	Created  int64    `json:"created"`
	Start    int64    `json:"start"`
	End      int64    `json:"end"`
	Snapshot int64    `json:"snapshot"`
}

type SignedProposalIFPS struct {
	ID        string `json:"id"`
	Signature string `json:"signature"`
	Proposal
}

type ProposalIFPS struct {
	ID string `json:"id"`
	Proposal
}

func (h handler) PostProposal(c echo.Context) error {
	ctx := c.Request().Context()
	collection := h.db.Client().Database("etb-voting-app").Collection("proposals")

	var signedProposal struct {
		Signature string   `json:"signature"`
		Proposal  Proposal `json:"proposal"`
	}

	if err := c.Bind(&signedProposal); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	name := slug.MakeLang(signedProposal.Proposal.Title, "en")
	hash, err := h.ipfs.PinJSON("etb/proposal/"+strconv.FormatInt(signedProposal.Proposal.Created, 10)+"-"+name, signedProposal)
	if err != nil {
		//log.Fatal(err) // TODO
		return c.JSON(http.StatusInternalServerError, err)
	}

	var final SignedProposalIFPS
	final.ID = hash
	final.Signature = signedProposal.Signature
	final.Proposal = signedProposal.Proposal

	_, err = collection.InsertOne(ctx, final)
	if err != nil {
		//log.Fatal(err) // TODO
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.String(http.StatusOK, hash)
}

func (h handler) GetProposals(c echo.Context) error {
	ctx := c.Request().Context()
	collection := h.db.Client().Database("etb-voting-app").Collection("proposals")

	cur, err := collection.Find(ctx, bson.D{}, &options.FindOptions{
		Sort: bson.D{{"_id", -1}},
	})
	if err != nil {
		//log.Fatal(err) // TODO
		return c.JSON(http.StatusInternalServerError, err)
	}
	defer cur.Close(ctx)

	var proposals []ProposalIFPS
	for cur.Next(ctx) {
		var proposal ProposalIFPS
		err = cur.Decode(&proposal)
		if err != nil {
			//log.Fatal(err) // TODO
			return c.JSON(http.StatusInternalServerError, err)
		}

		proposals = append(proposals, proposal)
	}
	if err = cur.Err(); err != nil {
		//log.Fatal(err) // TODO
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, struct {
		Proposals []ProposalIFPS `json:"proposals"`
	}{
		Proposals: proposals,
	})
}

func (h handler) GetProposal(c echo.Context) error {
	ctx := c.Request().Context()
	collection := h.db.Client().Database("etb-voting-app").Collection("proposals")

	id := c.Param("id")

	var proposal SignedProposalIFPS
	err := collection.FindOne(ctx, bson.D{{"id", id}}).Decode(&proposal)
	if err == mongo.ErrNoDocuments {
		return c.JSON(http.StatusNotFound, "space not found")
	} else if err != nil {
		//log.Fatal(err) // TODO
		return c.JSON(http.StatusInternalServerError, err)
	}

	collectionVotes := h.db.Client().Database("etb-voting-app").Collection("votes")

	cur, err := collectionVotes.Find(ctx, bson.D{{"vote.proposal", proposal.ID}})
	if err != nil {
		//log.Fatal(err) // TODO
		return c.JSON(http.StatusInternalServerError, err)
	}
	defer cur.Close(ctx)

	var votes []VoteIFPS
	for cur.Next(ctx) {
		var vote VoteIFPS
		err = cur.Decode(&vote)
		if err != nil {
			//log.Fatal(err) // TODO
			return c.JSON(http.StatusInternalServerError, err)
		}

		votes = append(votes, vote)
	}
	if err = cur.Err(); err != nil {
		//log.Fatal(err) // TODO
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, struct {
		Proposal SignedProposalIFPS `json:"proposal"`
		Votes    []VoteIFPS         `json:"votes"`
	}{
		Proposal: proposal,
		Votes:    votes,
	})
}
