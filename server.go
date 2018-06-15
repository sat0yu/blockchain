package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// =================================
	//  initialize block chain
	// =================================

	bc := newBlockchain()
	var lastBlock *Block
	var lastProof, proof int

	// =================================
	// launch http server
	// =================================

	handleChain := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, fmt.Sprintln(bc))
	})
	handleNewTransaction := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var t Transaction
		err := decoder.Decode(&t)
		if err != nil {
			panic(err)
		}
		bc.newTransaction(t.Sender, t.Recipient, t.Amount)
		fmt.Fprintf(w, fmt.Sprintln(bc))
	})
	handleProofOfWork := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lastBlock = bc.lastBlock()
		lastProof = lastBlock.Proof
		proof = bc.proofOfWork(lastProof)
		fmt.Fprintf(w, fmt.Sprintln(proof))
	})
	handleCommit := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var c Commit
		err := decoder.Decode(&c)
		if err != nil {
			panic(err)
		}
		bc.newTransaction("master", "0", 1)
		bc.newBlock(c.PreviousHash, c.Proof)
		fmt.Fprintf(w, fmt.Sprintln(bc))
	})
	http.Handle("/chain", handleChain)
	http.Handle("/new_transaction", handleNewTransaction)
	http.Handle("/proof_of_work", handleProofOfWork)
	http.Handle("/commit", handleCommit)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
