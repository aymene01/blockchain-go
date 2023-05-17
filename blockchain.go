package main

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

type Transaction struct {
	Sender   string
	Receiver string
	Amount   int
}

type Block struct {
	Index        int
	Timestamp    string
	Transactions []Transaction
	PrevHash     string
	Hash         string
}

type Blockchain struct {
	chain []Block
}

func (bc *Blockchain) addBlock(newBlock Block) {
	bc.chain = append(bc.chain, newBlock)
}

func (bc *Blockchain) getBlockchain() []Block {
	return bc.chain
}

func calculateHash(block Block) string {
	record := strconv.Itoa(block.Index) + block.Timestamp + block.PrevHash
	hash := sha256.New()
	hash.Write([]byte(record))
	hashed := hash.Sum(nil)
	return hex.EncodeToString(hashed)
}

func generateBlock(prevBlock Block, transactions []Transaction) Block {
	var newBlock Block
	t := time.Now()

	newBlock.Index = prevBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.Transactions = transactions
	newBlock.PrevHash = prevBlock.Hash
	newBlock.Hash = calculateHash(newBlock)

	return newBlock
}
