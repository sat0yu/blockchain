package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

const difficulty = 4

type blockchain struct {
	chain              []Block
	currentTransaction []Transaction
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
	b.lastBlock().index++
	return b.lastBlock().index
}

func (b *blockchain) proofOfWork(lastProof int) int {
	var p = 0
	for ; validateProof(lastProof, p); p++ {
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
	b := new(blockchain)
	b.newBlock("1", 100)
	return b
}

func main() {
	x := 5
	p := 0
	for ; !validateProof(x, p); p++ {
	}
	fmt.Printf("%d\t%s\n", p, genHashValue(x, p))
}
