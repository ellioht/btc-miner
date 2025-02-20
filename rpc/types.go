package rpc

import (
	"encoding/json"
	"errors"
	"github.com/ellioht/btc-miner/common"
)

var (
	ErrInvalidBlockHashOrNumber = errors.New("invalid core hash or number")
	ErrInvalidVerbosity         = errors.New("invalid verbosity")
)

type Method string

const (
	MethodGetBlock          Method = "getblock"
	MethodGetBlockHash      Method = "getblockhash"
	MethodGetBlockTemplate  Method = "getblocktemplate"
	MethodGetRawTransaction Method = "getrawtransaction"
	MethodHelp              Method = "help"
)

type Request struct {
	JsonRpc string          `json:"jsonrpc"`
	Method  Method          `json:"method"`
	Params  json.RawMessage `json:"params,omitempty"`
	Id      int             `json:"id"`
}

func EmptyRequest(method Method) *Request {
	return &Request{
		JsonRpc: "1.0",
		Method:  method,
		Id:      1,
	}
}

type Response struct {
	Result json.RawMessage `json:"result"`
	Error  *Error          `json:"error"`
	Id     int             `json:"id"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type BlockHashOrNumber struct {
	Hash   *common.Hash `json:"hash,omitempty"`
	Number *uint64      `json:"number,omitempty"`
}

func BlockNumber(number uint64) *BlockHashOrNumber {
	return &BlockHashOrNumber{Number: &number}
}

func BlockHash(hash common.Hash) *BlockHashOrNumber {
	return &BlockHashOrNumber{Hash: &hash}
}
