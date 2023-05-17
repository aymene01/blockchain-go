package blockchain

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
	Chain []Block
}

func (bc *Blockchain) AddBlock(newBlock Block) {
	bc.Chain = append(bc.Chain, newBlock)
}

func (bc *Blockchain) GetBlockchain() []Block {
	return bc.Chain
}

func CalculateHash(block Block) string {
	record := strconv.Itoa(block.Index) + block.Timestamp + block.PrevHash
	hash := sha256.New()
	hash.Write([]byte(record))
	hashed := hash.Sum(nil)
	return hex.EncodeToString(hashed)
}

func GenerateBlock(prevBlock Block, transactions []Transaction) Block {
	var newBlock Block
	t := time.Now()

	newBlock.Index = prevBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.Transactions = transactions
	newBlock.PrevHash = prevBlock.Hash
	newBlock.Hash = CalculateHash(newBlock)

	return newBlock
}
