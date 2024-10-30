package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"log"

	"github.com/Artexus/PoW/entity"
)

type Data struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Amount int    `json:"amount"`
}

func main() {
	chain := entity.Blockchain{
		Difficulty: 4,
		Blocks:     []entity.Block{},
	}

	d := Data{
		From:   hex.EncodeToString(sha256.New().Sum(nil)),
		To:     hex.EncodeToString(sha256.New().Sum(nil)),
		Amount: 100,
	}

	body, _ := json.Marshal(d)
	err := chain.Mine(body)
	if err != nil {
		log.Println(err)
		return
	}

	d = Data{
		From:   hex.EncodeToString(sha256.New().Sum(nil)),
		To:     hex.EncodeToString(sha256.New().Sum(nil)),
		Amount: 300,
	}

	body, _ = json.Marshal(d)
	err = chain.Mine(body)
	if err != nil {
		log.Println(err)
		return
	}

	for _, block := range chain.Blocks {
		data := Data{}

		_ = json.Unmarshal(block.Data, &data)
		log.Println(block.Index, block.Timestamp, data, block.Nonce)
	}
}
