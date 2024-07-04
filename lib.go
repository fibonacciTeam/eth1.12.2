package eth1_12_2

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/metachris/eth-go-bindings/erc20"
	"math/big"
)

func Dial(rawurl string) (*ethclient.Client, error) {
	return ethclient.Dial(rawurl)
}

func GetBalance(rawurl string, contractAddress common.Address, account common.Address) (*big.Int, error) {
	cli, err := ethclient.Dial(rawurl)
	if err != nil {
		return &big.Int{}, err
	}

	contractIns, err := erc20.NewErc20(contractAddress, cli)
	if err != nil {
		return &big.Int{}, err
	}

	b, err := contractIns.BalanceOf(nil, account)
	if err != nil {
		return &big.Int{}, err
	}

	return b, nil
}
