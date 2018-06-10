package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"os"
)

type Block struct {
	Index        int
	PreviousHash string
	Proof        int
	Timestamp    int64
	Transactions []Transaction
}

func (b *Block) Hash() string {
	jsonBytes, err := json.Marshal(*b)
	if err != nil {
		fmt.Println("json marshal error")
		os.Exit(1)
	}
	return fmt.Sprintf("%x", sha256.Sum256(jsonBytes))
}
