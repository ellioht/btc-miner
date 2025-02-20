package rpc

import (
	"encoding/json"
	"github.com/ellioht/btc-miner/common"
)

type Verbosity uint8

const (
	// SerializedBlock 0: returns a serialized core in hex format.
	SerializedBlock Verbosity = iota
	// BlockInfo 1: returns an Object with information about the core.
	BlockInfo
	// BlockTxInfo 2: returns an Object with information about the core and information about each transaction.
	BlockTxInfo
)

type GetBlockResult struct {
	Hash          common.Hash `json:"hash"`
	Confirmations int64       `json:"confirmations"`
	Height        int64       `json:"height"`
	Version       int32       `json:"version"`
	VersionHex    string      `json:"versionHex"`
	MerkleRoot    string      `json:"merkleroot"`
	Time          int64       `json:"time"`
	MedianTime    int64       `json:"mediantime"`
	Nonce         uint32      `json:"nonce"`
	Bits          string      `json:"bits"`
	Difficulty    float64     `json:"difficulty"`
	ChainWork     common.Hash `json:"chainwork"`
	NTx           int32       `json:"nTx"`
	NextBlockHash common.Hash `json:"nextblockhash,omitempty"`
	StrippedSize  int32       `json:"strippedsize"`
	Size          int32       `json:"size"`
	Weight        int32       `json:"weight"`
	Tx            []Tx        `json:"tx,omitempty"`
}

type Tx struct {
	TxId     common.Hash `json:"txid"`
	Hash     common.Hash `json:"hash"`
	Version  int32       `json:"version"`
	Size     int32       `json:"size"`
	VSize    int32       `json:"vsize"`
	Weight   int32       `json:"weight"`
	LockTime int32       `json:"locktime"`
	VIn      []VIn       `json:"vin"`
	VOut     []VOut      `json:"vout"`
	Hex      string      `json:"hex"`
	Fee      float64     `json:"fee,omitempty"`
}

type VIn struct {
	Coinbase  string    `json:"coinbase"`
	Sequence  int64     `json:"sequence"`
	TxWitness []string  `json:"txinwitness,omitempty"`
	TxId      string    `json:"txid,omitempty"`
	VOut      int32     `json:"vout,omitempty"`
	ScriptSig ScriptSig `json:"scriptSig,omitempty"`
}

type ScriptSig struct {
	Asm string `json:"asm"`
	Hex string `json:"hex"`
}

type VOut struct {
	Value        float64      `json:"value"`
	N            int32        `json:"n"`
	ScriptPubKey ScriptPubKey `json:"scriptPubKey"`
}

type ScriptPubKey struct {
	Asm     string `json:"asm"`
	Desc    string `json:"desc"`
	Hex     string `json:"hex"`
	Address string `json:"address"`
	Type    string `json:"type"`
}

func UnmarshalGetBlockResult(data []byte) (*GetBlockResult, error) {
	var result GetBlockResult
	err := json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
