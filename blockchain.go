package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

const difficulty = 5

type blockchain struct {
	chain              []Block
	currentTransaction []Transaction
}

func (b *blockchain) String() string {
	var chain = "chain: \n"
	for _, block := range b.chain {
		chain += fmt.Sprintf("{\n\tindex: %d\n\tprevHash: %s\n\tts: %d\n\ttransactions: %v\n},\n", block.Index, block.PreviousHash, block.Timestamp, block.Transactions)
	}
	var currentTransaction = fmt.Sprintf("currentTransaction: %v\n", b.currentTransaction)
	return chain + currentTransaction
}

func (b *blockchain) lastBlock() *Block {
	return &b.chain[len(b.chain)-1]
}

func (b *blockchain) newBlock(previousHash string, proof int) Block {
	block := Block{
		Index:        len(b.chain) + 1,
		Timestamp:    time.Now().UnixNano(),
		Transactions: b.currentTransaction,
		Proof:        proof,
		PreviousHash: previousHash,
	}
	if len(previousHash) == 0 {
		block.PreviousHash = b.lastBlock().Hash()
	}
	b.currentTransaction = []Transaction{}
	b.chain = append(b.chain, block)
	return block
}

func (b *blockchain) newTransaction(sender string, recipient string, amount int) int {
	transaction := Transaction{sender, recipient, amount}
	b.currentTransaction = append(b.currentTransaction, transaction)
	return b.lastBlock().Index + 1
}

func (b *blockchain) proofOfWork(lastProof int) int {
	var p = 0
	for ; !validateProof(lastProof, p); p++ {
	}
	return p
}

func genHashValue(a, b int) string {
	str := fmt.Sprintf("%d%d", a, b)
	jsonBytes, err := json.Marshal(str)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return fmt.Sprintf("%x", sha256.Sum256(jsonBytes))
}

func validateProof(lastProof, proof int) bool {
	digest := genHashValue(lastProof, proof)
	for _, ch := range digest[len(digest)-difficulty : len(digest)] {
		if ch != '0' {
			return false
		}
	}
	return true
}

func newBlockchain() *blockchain {
	bc := new(blockchain)
	bc.newBlock("1", 100)
	return bc
}

// func main() {
// 	bc := newBlockchain()

// 	var lastBlock *Block
// 	var lastProof, proof int
// 	var prevHash string

// 	// add transaction data to the second block
// 	bc.newTransaction("a", "b", 5)
// 	// enclose the second block with proof
// 	lastBlock = bc.lastBlock()
// 	lastProof = lastBlock.Proof
// 	proof = bc.proofOfWork(lastProof)
// 	bc.newTransaction("master", "0", 1)
// 	prevHash = lastBlock.Hash()
// 	bc.newBlock(prevHash, proof)

// 	// enclose the third block with proof (no transaction)
// 	lastBlock = bc.lastBlock()
// 	lastProof = lastBlock.Proof
// 	proof = bc.proofOfWork(lastProof)
// 	bc.newTransaction("master", "0", 1)
// 	prevHash = lastBlock.Hash()
// 	bc.newBlock(prevHash, proof)

// 	// add transaction data to the third block
// 	bc.newTransaction("a", "b", 5)
// 	bc.newTransaction("a", "b", 5)
// 	// enclose the fourth block with proof
// 	lastBlock = bc.lastBlock()
// 	lastProof = lastBlock.Proof
// 	proof = bc.proofOfWork(lastProof)
// 	bc.newTransaction("master", "0", 1)
// 	prevHash = lastBlock.Hash()
// 	bc.newBlock(prevHash, proof)

// 	fmt.Println(bc.chain)
// }
