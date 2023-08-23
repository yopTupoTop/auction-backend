package handler

import (
	//"regexp"
	"testing"

	"auction_backend/event"
	"auction_backend/model"

	"github.com/ethereum/go-ethereum/common"
)

func TestPostHandler(t *testing.T) {
	newContract := model.ContractCreated{
		Auction: common.Address{0x00},
		Owner:   common.Address{0x00},
	}

	event.ContractsCreated = append(event.ContractsCreated, newContract)

}
