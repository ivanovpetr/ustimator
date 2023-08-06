package uniswapclient

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// PoolContract wrapper for ABI interface of uniswap v2 pair contract
type PoolContract struct {
	*UniswapPairV2
}

// NewPoolContract creates instance of uniswap v2 pair
func NewPoolContract(address string, url string) (*PoolContract, error) {
	conn, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}

	uniswapPairV2, err := NewUniswapPairV2(common.HexToAddress(address), conn)
	if err != nil {
		return nil, err
	}

	return &PoolContract{uniswapPairV2}, nil
}

// GetReserves returns reserves for uniswap v2 pair tokens, token0 and token1 respectively
func (c *PoolContract) GetReserves() (reserve0, reserve1 *big.Int, err error) {
	reserves, err := c.UniswapPairV2.GetReserves(nil)
	if err != nil {
		return nil, nil, err
	}
	return reserves.Reserve0, reserves.Reserve1, nil
}

// TokenAddresses return addresses of uniswap v2 pair tokens, token0 and token1 respectively
func (c *PoolContract) TokenAddresses() (string, string, error) {
	token0, err := c.Token0(nil)
	if err != nil {
		return "", "", nil
	}
	token1, err := c.Token1(nil)
	if err != nil {
		return "", "", nil
	}
	return token0.Hex(), token1.Hex(), nil
}
