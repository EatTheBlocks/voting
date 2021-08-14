package eth

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/onrik/ethrpc"
)

type Call struct {
	Target   common.Address
	CallData []byte
}

type Result struct {
	Address string
	Balance *big.Int
}

func Multicall(client *ethrpc.EthRPC, multicallAddress string, tokenAddress string, blockNumber uint64, addresses []string) ([]*Result, error) {
	multicallABI, err := abi.JSON(strings.NewReader(`[{"inputs":[{"components":[{"internalType":"address","name":"target","type":"address"},{"internalType":"bytes","name":"callData","type":"bytes"}],"internalType":"structMulticall.Call[]","name":"calls","type":"tuple[]"}],"name":"aggregate","outputs":[{"internalType":"uint256","name":"blockNumber","type":"uint256"},{"internalType":"bytes[]","name":"returnData","type":"bytes[]"}],"stateMutability":"view","type":"function"}]`))
	if err != nil {
		return nil, err
	}

	var calls []Call

	balanceOf, err := balanceOfData()
	if err != nil {
		return nil, err
	}

	for _, address := range addresses {
		data, err := balanceOf(address)
		if err != nil {
			return nil, err
		}
		calls = append(calls, Call{common.HexToAddress(tokenAddress), data})
	}

	multicallData, err := multicallABI.Pack("aggregate", calls)
	if err != nil {
		return nil, err
	}

	tag := "latest"

	if blockNumber > 0 {
		tag = hexutil.EncodeUint64(blockNumber)
	}

	result, err := client.EthCall(ethrpc.T{
		From: multicallAddress,
		To:   multicallAddress,
		Data: hexutil.Encode(multicallData),
	}, tag)
	if err != nil {
		return nil, err
	}

	resultData, err := hexutil.Decode(result)
	if err != nil {
		return nil, err
	}

	out, err := multicallABI.Unpack("aggregate", resultData)
	if err != nil {
		return nil, err
	}

	outstruct := new(struct {
		BlockNumber *big.Int
		ReturnData  [][]byte
	})

	outstruct.BlockNumber = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ReturnData = *abi.ConvertType(out[1], new([][]byte)).(*[][]byte)

	var results []*Result

	for i, address := range addresses {
		balance := &big.Int{}
		balance.SetBytes(outstruct.ReturnData[i])

		results = append(results, &Result{address, balance})
	}

	return results, nil
}

func balanceOfData() (func(string) ([]byte, error), error) {
	erc20ABI, err := abi.JSON(strings.NewReader(`[{"inputs":[{"internalType":"address","name":"account","type":"address"}],"name":"balanceOf","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"}]`))
	if err != nil {
		return nil, err
	}

	return func(address string) ([]byte, error) {
		return erc20ABI.Pack("balanceOf", common.HexToAddress(address))
	}, nil
}
