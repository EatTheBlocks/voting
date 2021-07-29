package handler

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

type SnapshotScoreRequest struct {
	Params SnapshotParams `json:"params"`
}

type SnapshotParams struct {
	Network    string             `json:"network"`
	Snapshot   int64              `json:"snapshot"`
	Strategies []SnapshotStrategy `json:"strategies"`
	Addresses  []string           `json:"addresses"`
}

type SnapshotStrategy struct {
	Name   string    `json:"name"`
	Params ABIParams `json:"params"`
}

type ABIParams struct {
	Symbol    string    `json:"symbol"`
	Address   string    `json:"address"`
	Decimals  int       `json:"decimals"`
	MethodABI ABIMethod `json:"methodABI"`
}

type ABIMethod struct {
	Inputs          []ABIMethodParam `json:"inputs"`
	Name            string           `json:"name"`
	Outputs         []ABIMethodParam `json:"outputs"`
	StateMutability string           `json:"stateMutability"`
	Type            string           `json:"type"`
}

type ABIMethodParam struct {
	InternalType string `json:"internalType"`
	Name         string `json:"name"`
	Type         string `json:"type"`
}

type SnapshotScoreResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		State  string               `json:"state"`
		Scores []map[string]float64 `json:"scores"`
	} `json:"result"`
}

type ScoreRequest struct {
	Snapshot  int64    `json:"snapshot"`
	Addresses []string `json:"addresses"`
}

// GetScore temporary solution that use score.snapshot.org server for archive data
func (h handler) GetScore(c echo.Context) error {
	var scoreRequest ScoreRequest

	if err := c.Bind(&scoreRequest); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	snapshotScoreResponse, err := getScore(scoreRequest)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, snapshotScoreResponse)
}

func getScore(scoreRequest ScoreRequest) (*SnapshotScoreResponse, error) {
	snapshotScoreRequest := SnapshotScoreRequest{Params: SnapshotParams{
		Network:  "56",
		Snapshot: scoreRequest.Snapshot,
		Strategies: []SnapshotStrategy{
			{
				Name: "contract-call",
				Params: ABIParams{
					Symbol:   "ETB",
					Address:  "0x7ac64008fa000bfdc4494e0bfcc9f4eff3d51d2a",
					Decimals: 18,
					MethodABI: ABIMethod{
						Name:            "balanceOf",
						Type:            "function",
						StateMutability: "view",
						Inputs: []ABIMethodParam{
							{
								InternalType: "address",
								Name:         "account",
								Type:         "address",
							},
						},
						Outputs: []ABIMethodParam{
							{
								InternalType: "uint256",
								Name:         "",
								Type:         "uint256",
							},
						},
					},
				},
			},
		},
		Addresses: scoreRequest.Addresses,
	}}

	jsonBody, err := json.Marshal(snapshotScoreRequest)
	if err != nil {
		log.Error(errors.Wrap(err, "GetScore json encode"))
		return nil, err
	}

	req, err := http.NewRequest("POST", "https://score.snapshot.org/api/scores", bytes.NewReader(jsonBody))
	if err != nil {
		log.Error(errors.Wrap(err, "GetScore new request"))
		return nil, err
	}

	req.Header.Set("authority", "score.snapshot.org")
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.115 Safari/537.36")
	req.Header.Set("content-type", "application/json")
	req.Header.Set("accept", "*/*")
	req.Header.Set("sec-gpc", "1")
	req.Header.Set("origin", "https://snapshot.org")
	req.Header.Set("sec-fetch-site", "same-site")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("referer", "https://snapshot.org/")
	req.Header.Set("accept-language", "en-US,en;q=0.9")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Error(errors.Wrap(err, "GetScore send request"))
		return nil, err
	}

	defer resp.Body.Close()

	var snapshotScoreResponse SnapshotScoreResponse
	decoder := json.NewDecoder(resp.Body)
	if err = decoder.Decode(&snapshotScoreResponse); err != nil {
		log.Error(errors.Wrap(err, "GetScore parse json response"))
		return nil, err
	}

	return &snapshotScoreResponse, nil
}
