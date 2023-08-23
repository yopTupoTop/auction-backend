package model

import (
	"github.com/ethereum/go-ethereum/common"
)

type ContractCreated struct {
	Auction common.Address //`json: "auction"`
	Owner   common.Address //`json: "owner"`
}
