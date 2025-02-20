package rpc

import (
	"encoding/json"
	"github.com/ellioht/btc-miner/common"
)

type GetRawTransactionResult struct {
	Txid          common.Hash `json:"txid"`
	Hash          common.Hash `json:"hash"`
	Version       int32       `json:"version"`
	Size          int32       `json:"size"`
	VSize         int32       `json:"vsize"`
	Weight        int32       `json:"weight"`
	LockTime      int64       `json:"locktime"`
	VIn           []VIn       `json:"vin"`
	VOut          []VOut      `json:"vout"`
	Hex           string      `json:"hex"`
	BlockHash     common.Hash `json:"blockhash"`
	Confirmations int64       `json:"confirmations"`
	Time          int64       `json:"time"`
	BlockTime     int64       `json:"blocktime"`
}

type TxFullOrHex struct {
	TxFull *GetRawTransactionResult
	TxHex  *string
}

func UnmarshalGetRawTransactionResult(data []byte) (*GetRawTransactionResult, error) {
	var r GetRawTransactionResult
	err := json.Unmarshal(data, &r)
	return &r, err
}
