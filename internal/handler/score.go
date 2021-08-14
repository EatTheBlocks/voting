package handler

import (
	"math/big"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/onrik/ethrpc"
	"github.com/shopspring/decimal"
	"voting-hub/internal/eth"
)

type ScoreResult struct {
	State  string             `json:"state"`
	Scores map[string]float64 `json:"scores"`
}

type ScoreRequest struct {
	Snapshot  uint64   `json:"snapshot"`
	Addresses []string `json:"addresses"`
}

func (h handler) GetScore(c echo.Context) error {
	var scoreRequest ScoreRequest

	if err := c.Bind(&scoreRequest); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	snapshotScoreResponse, err := getScore(h.rpcApi, h.multicallAddress, h.tokenAddress, scoreRequest)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, snapshotScoreResponse)
}

func getScore(rpcApi, multicallAddress, tokenAddress string, scoreRequest ScoreRequest) (*ScoreResult, error) {
	client := ethrpc.New(rpcApi)

	results, err := eth.Multicall(client, multicallAddress, tokenAddress, scoreRequest.Snapshot, scoreRequest.Addresses)
	if err != nil {
		return nil, err
	}

	var scoreResult ScoreResult
	scoreResult.State = "final" // TODO
	scoreResult.Scores = make(map[string]float64)

	for _, result := range results {
		balance, _ := toDecimal(result.Balance, 18).Float64()
		scoreResult.Scores[result.Address] = balance
	}

	return &scoreResult, nil
}

func toDecimal(ivalue interface{}, decimals int) decimal.Decimal {
	value := new(big.Int)
	switch v := ivalue.(type) {
	case string:
		value.SetString(v, 10)
	case *big.Int:
		value = v
	}

	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	num, _ := decimal.NewFromString(value.String())
	result := num.Div(mul)

	return result
}
