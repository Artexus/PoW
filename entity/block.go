package entity

import "time"

type Block struct {
	Index        int
	Timestamp    time.Time
	Data         []byte
	Nonce        int
	Hash         string
	PreviousHash string
}
