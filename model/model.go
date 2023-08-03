package model

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
)

type PlacedAssets struct {
	Seller  common.Address `json:"seller"`
	TokenId uint256.Int    `json:"tokenId"`
	Price   uint256.Int    `json:"price"`
	Time    uint256.Int    `json:"time"`
}

type CanceledAssets struct {
	Seller  common.Address `json:"seller"`
	TokenId uint256.Int    `json:"tokenId"`
	Price   uint256.Int    `json:"price"`
	Time    uint256.Int    `json:"time"`
}

type AcceptedOffers struct {
	Seller  common.Address `json:"seller"`
	Buyer   common.Address `json:"buyer"`
	TokenId uint256.Int    `json:"tokenId"`
	Price   uint256.Int    `json:"price"`
	Time    uint256.Int    `json:"time"`
}

type BoughtAssets struct {
	Bidder  common.Address `json:"bidder"`
	TokenId uint256.Int    `json:"tokenId"`
	Price   uint256.Int    `json:"price"`
	Time    uint256.Int    `json:"time"`
}

type PlacedBids struct {
	Bidder  common.Address `json:"bidder"`
	TokenId uint256.Int    `json:"tokenId"`
	Price   uint256.Int    `json:"price"`
	Time    uint256.Int    `json:"time"`
}
