package blockchain

import "testing"

func TestBlochain(t *testing.T) {
	bc := Blockchain{}

	genesisBlock := Block{
		Index:        0,
		Timestamp:    "2020-01-01 00:00:00",
		Transactions: []Transaction{},
		PrevHash:     "",
		Hash:         "",
	}

	genesisBlock.Hash = CalculateHash(genesisBlock)

	bc.AddBlock(genesisBlock)

	transaction := Transaction{Sender: "Alice", Receiver: "Bob", Amount: 100}
	newBlock := GenerateBlock(bc.Chain[len(bc.Chain)-1], []Transaction{transaction})
	bc.AddBlock(newBlock)

	if len(bc.Chain) != 2 {
		t.Errorf("Expected blockchain length to be 2, got %d", len(bc.Chain))
	}

	if bc.Chain[1].Index != 1 {
		t.Errorf("Expected block index to be 1, got %d", bc.Chain[1].Index)
	}

	if bc.Chain[1].Transactions[0].Sender != "Alice" {
		t.Errorf("Expected transaction sender to be Alice, got %s", bc.Chain[1].Transactions[0].Sender)
	}

	if bc.Chain[1].Transactions[0].Receiver != "Bob" {
		t.Errorf("Expected transaction receiver to be Bob, got %s", bc.Chain[1].Transactions[0].Receiver)
	}

	if bc.Chain[1].Transactions[0].Amount != 100 {
		t.Errorf("Expected transaction amount to be 100, got %d", bc.Chain[1].Transactions[0].Amount)
	}

	if bc.Chain[1].PrevHash != genesisBlock.Hash {
		t.Errorf("Expected block prev hash to be %s, got %s", genesisBlock.Hash, bc.Chain[1].PrevHash)
	}

	if bc.Chain[1].Hash != CalculateHash(bc.Chain[1]) {
		t.Errorf("Expected block hash to be %s, got %s", CalculateHash(bc.Chain[1]), bc.Chain[1].Hash)
	}
}
