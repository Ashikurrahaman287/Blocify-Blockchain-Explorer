package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

// Block represents a block in the blockchain
type Block struct {
	Index        int
	Timestamp    time.Time
	PreviousHash string
	Hash         string
	Data         string
}

// Blockchain represents the blockchain
type Blockchain struct {
	Chain []Block
}

// NewBlock creates a new block in the blockchain
func (bc *Blockchain) NewBlock(data string, previousHash string) *Block {
	block := &Block{
		Index:        len(bc.Chain),
		Timestamp:    time.Now(),
		PreviousHash: previousHash,
		Data:         data,
	}
	block.Hash = block.calculateHash()
	return block
}

// AddBlock adds a new block to the blockchain
func (bc *Blockchain) AddBlock(data string) {
	previousBlock := bc.Chain[len(bc.Chain)-1]
	newBlock := bc.NewBlock(data, previousBlock.Hash)
	bc.Chain = append(bc.Chain, *newBlock)
}

// calculateHash calculates the hash of a block
func (b *Block) calculateHash() string {
	record := string(b.Index) + b.Timestamp.String() + b.PreviousHash + b.Data
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

// NewBlockchain creates a new blockchain with the genesis block
func NewBlockchain() *Blockchain {
	genesisBlock := &Block{
		Index:        0,
		Timestamp:    time.Now(),
		PreviousHash: "",
		Data:         "Genesis Block",
	}
	genesisBlock.Hash = genesisBlock.calculateHash()
	return &Blockchain{Chain: []Block{*genesisBlock}}
}
