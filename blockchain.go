package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type blockchain struct {
	chain              []Block
	currentTransaction []Transaction
}

func (b *blockchain) lastBlock() *Block {
	return &b.chain[len(b.chain)-1]
}

func (b *blockchain) newBlock(previousHash string, proof int) Block {
	block := Block{
		index:        len(b.chain) + 1,
		timestamp:    time.Now().UnixNano(),
		transactions: b.currentTransaction,
		proof:        proof,
		previousHash: previousHash,
	}
	if len(previousHash) == 0 {
		block.previousHash = hash(b.lastBlock())
	}
	b.currentTransaction = []Transaction{}
	b.chain = append(b.chain, block)
	return block
}

func (b *blockchain) newTransaction(sender string, recipient string, amount int) {
	transaction := Transaction{sender, recipient, amount}
	b.currentTransaction = append(b.currentTransaction, transaction)
	b.lastBlock().index++
}

func hash(block *Block) string {
	jsonBytes, err := json.Marshal(block)
	if err != nil {
		fmt.Println("json marshal error")
		os.Exit(1)
	}
	digest := sha256.Sum256(jsonBytes)
	return string(digest[:])
}

func newBlockchain() *blockchain {
	b := new(blockchain)
	b.newBlock("1", 100)
	return b
}

func main() {
}
