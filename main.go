package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	var bc Blockchain

	genesisBlock := Block{0, time.Now().String(), []Transaction{}, "", ""}
	genesisBlock.Hash = calculateHash(genesisBlock)

	bc.addBlock(genesisBlock)

	for {
		fmt.Println("Enter sender (or 'q' to quit):")
		var sender string
		fmt.Scanln(&sender)

		if sender == "q" {
			fmt.Println("Exiting...")
			os.Exit(0)
		}

		if sender == "blocks" {
			bc := bc.getBlockchain()

			for _, block := range bc {
				fmt.Println("-------------------")
				for _, transaction := range block.Transactions {
					fmt.Println("sender:", transaction.Sender)
					fmt.Println("receiver:", transaction.Receiver)
				}
				fmt.Println("Hash", block.Hash)
			}
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

		transaction := Transaction{Sender: sender, Receiver: receiver, Amount: amount}

		newBlock := generateBlock(bc.chain[len(bc.chain)-1], []Transaction{transaction})
		bc.addBlock(newBlock)

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
