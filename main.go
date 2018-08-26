package main

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://mainnet.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	// Get the latest block
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	blockNumber := header.Number
	fmt.Println("Latest block: " + blockNumber.String())

	// Query a specific block by number
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Block time: " + block.Time().String())
	fmt.Println("Block difficulty: " + block.Difficulty().String())
	fmt.Println("Block hash: " + block.Hash().String())
	fmt.Println("Number of TXs: " + strconv.Itoa(len(block.Transactions())))

	// Query a specific block by hash
	blockByHash, err := client.BlockByHash(context.Background(), block.Hash())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Block time: " + blockByHash.Time().String())
	fmt.Println("Block difficulty: " + blockByHash.Difficulty().String())
	fmt.Println("Block hash: " + blockByHash.Hash().String())
	fmt.Println("Number of TXs: " + strconv.Itoa(len(blockByHash.Transactions())))
}
