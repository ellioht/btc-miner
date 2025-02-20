package core

import "github.com/ellioht/btc-miner/common"

type Tx struct {
	TxId          common.Hash
	Hash          common.Hash
	Version       int32
	Size          int32
	VSize         int32
	Weight        int32
	LockTime      int64
	VIn           []VIn
	VOut          []VOut
	Hex           string
	BlockHash     common.Hash
	Confirmations int64
	Time          int64
	BlockTime     int64
}

type VIn struct{}

type VOut struct{}
