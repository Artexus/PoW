package entity

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
	"time"
)

type Blockchain struct {
	Difficulty int
	Blocks     []Block
}

func calculateHash(b Block) string {
	s := fmt.Sprintf("%v%v%v%v", b.Index, b.Timestamp.String(), string(b.Data), b.Nonce)
	v := sha256.New()
	v.Write([]byte(s))
	return hex.EncodeToString(v.Sum(nil))
}

func (b Blockchain) validateTransaction() bool {
	for i, block := range b.Blocks {
		hash := calculateHash(block)
		if b.Blocks[i].PreviousHash != b.Blocks[i-1].PreviousHash &&
			hash != block.Hash {
			return false
		}
	}

	return true
}

func (b Blockchain) calculateNonce(block Block) Block {
	for {
		block.Hash = calculateHash(block)
		if strings.HasPrefix(block.Hash, strings.Repeat("0", b.Difficulty)) {
			return block
		}

		block.Nonce++
	}
}

func (b *Blockchain) PoW(data []byte) Block {
	block := Block{
		Index:     len(b.Blocks),
		Timestamp: time.Now(),
		Data:      data,
	}

	if len(b.Blocks) != 0 {
		block.PreviousHash = b.Blocks[len(b.Blocks)-1].Hash
	}

	block = b.calculateNonce(block)
	return block
}

func (b *Blockchain) Mine(data []byte) (err error) {
	if !b.validateTransaction() {
		err = errors.New("validate transaction")
		return err
	}

	b.Blocks = append(b.Blocks, b.PoW(data))
	return
}
