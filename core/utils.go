package core

import (
	"crypto/rand"
	"github.com/ellioht/btc-miner/common"
	"io"
	mrand "math/rand"
	"time"
)

func RandomHash() common.Hash {
	hash := make([]byte, 32)
	_, err := io.ReadFull(rand.Reader, hash)
	if err != nil {
		panic(err)
	}
	return common.BytesToHash(hash)
}

func CreateRandomHeader() *Header {
	return &Header{
		Version:       1,
		PrevBlockHash: RandomHash(),
		MerkleRoot:    RandomHash(),
		Timestamp:     uint32(time.Now().Unix()),
		Bits:          mrand.Uint32(),
		Nonce:         mrand.Uint32(),
	}
}
