package main

type Block struct {
	index        int
	previousHash string
	proof        int
	timestamp    int64
	transactions []Transaction
}
