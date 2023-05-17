package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	bc "github.com/aymene01/go-blockchain/pkg/blockchain"
)

func main() {
	var memoryBc bc.Blockchain

	genesisBlock := bc.Block{0, time.Now().String(), []bc.Transaction{}, "", ""}
	genesisBlock.Hash = bc.CalculateHash(genesisBlock)

	memoryBc.AddBlock(genesisBlock)

	for {
		fmt.Println("Enter sender (or 'q' to quit and blocks to list all the blocks):")
		var sender string
		fmt.Scanln(&sender)

		if sender == "q" {
			fmt.Println("Exiting...")
			os.Exit(0)
		}

		if sender == "blocks" {
			bc := memoryBc.GetBlockchain()

			for _, block := range bc {
				fmt.Println("-------------------")
				if block.Index != 0 {
					fmt.Println("Transactions:")
				}
				for _, transaction := range block.Transactions {
					fmt.Println("- Sender:", transaction.Sender)
					fmt.Println("  Receiver:", transaction.Receiver)
					fmt.Println("  Amount:", transaction.Amount)
				}
				fmt.Println("Hash", block.Hash)
				if block.Index == 0 {
					fmt.Println("Genesis block")

				} else {
					fmt.Println("PrevHash", block.PrevHash)
				}
			}
			fmt.Println("-------------------")
			continue
		}

		fmt.Println("Enter receiver:")
		var receiver string
		fmt.Scanln(&receiver)

		fmt.Println("Enter amount:")
		var amountStr string
		fmt.Scanln(&amountStr)
		amount, err := strconv.Atoi(amountStr)
		if err != nil {
			fmt.Println("Invalid amount. Please try again.")
			continue
		}

		transaction := bc.Transaction{Sender: sender, Receiver: receiver, Amount: amount}

		newBlock := bc.GenerateBlock(memoryBc.Chain[len(memoryBc.Chain)-1], []bc.Transaction{transaction})
		memoryBc.AddBlock(newBlock)

		fmt.Println("Block added!")
		fmt.Println("Index:", newBlock.Index)
		fmt.Println("Timestamp:", newBlock.Timestamp)
		fmt.Println("Transactions:")

		for _, tx := range newBlock.Transactions {
			fmt.Println("- Sender:", tx.Sender)
			fmt.Println("  Receiver:", tx.Receiver)
			fmt.Println("  Amount:", tx.Amount)
		}
		fmt.Println("PrevHash:", newBlock.PrevHash)
		fmt.Println("Hash:", newBlock.Hash)
	}
}
