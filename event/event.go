package event

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"auction_backend/model"

	//"golang.org/x/tools/go/analysis/passes/errorsas"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var ContractsCreated []model.ContractCreated

func SubscribeToEvents() {
	client, err := ethclient.Dial("wss://eth-sepolia.g.alchemy.com/v2/1ShHEJUdELiemKeraKzLXG3US3yqjZzz")
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress("0x5dD7f6b9d30cA39171f8158fD7311522698f3BE1")
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	abiFile, err := os.Open("contract.abi")
	if err != nil {
		log.Fatal(err)
	}

	defer abiFile.Close()

	byteValue, _ := ioutil.ReadAll(abiFile)

	contractAbi, err := abi.JSON(strings.NewReader(string(byteValue)))
	if err != nil {
		log.Fatal(err)
	}

	logCreatedSig := []byte("ContractCreated(address auction, address owner)")
	logCreatedSigHash := crypto.Keccak256Hash(logCreatedSig)

	var createdEvent model.ContractCreated
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			switch vLog.Topics[0].Hex() {
			case logCreatedSigHash.Hex():
				err := contractAbi.UnpackIntoInterface(createdEvent, "ContractCreated", vLog.Data)
				if err != nil {
					log.Fatal(err)
				}
				ContractsCreated = append(ContractsCreated, createdEvent)
			}
		}
	}
}

//listen and store events to db from auction factory contract
