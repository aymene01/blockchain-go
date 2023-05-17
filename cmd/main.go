package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	bc "github.com/aymene01/go-blockchain/pkg/blockchain"
)

func main() {
	var memoryBc bc.Blockchain

	genesisBlock := bc.Block{
		Index:        0,
		Timestamp:    time.Now().String(),
		Transactions: []bc.Transaction{},
		PrevHash:     "",
		Hash:         "",
	}

	genesisBlock.Hash = bc.CalculateHash(genesisBlock)

	memoryBc.AddBlock(genesisBlock)

	validCommands := []string{"blocks", "transaction", "q"}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Enter a command:")
		if !scanner.Scan() {
			break
		}
		userInput := strings.TrimSpace(scanner.Text())

		switch userInput {
		case "q":
			fmt.Println("Exiting...")
			os.Exit(0)

		case "blocks":
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

		case "transaction":
			fmt.Println("Enter sender:")
			if !scanner.Scan() {
				break
			}
			sender := strings.TrimSpace(scanner.Text())

			fmt.Println("Enter receiver:")
			if !scanner.Scan() {
				break
			}
			receiver := strings.TrimSpace(scanner.Text())

			fmt.Println("Enter amount:")
			if !scanner.Scan() {
				break
			}
			amountStr := strings.TrimSpace(scanner.Text())
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

		default:
			fmt.Println("Invalid command. Valid commands are:")
			for _, command := range validCommands {
				fmt.Println("-", command)
			}
		}
	}
}
