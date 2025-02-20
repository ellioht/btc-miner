package rpc

import (
	"encoding/json"
	"github.com/ellioht/btc-miner/common"
)

type GetBlockTemplateParams struct {
	Mode         Mode           `json:"mode,omitempty"`
	Capabilities []Capabilities `json:"capabilities,omitempty"`
	Rules        []Rules        `json:"rules,omitempty"`
}

type Mode string

const (
	ModeTemplate Mode = "template"
	ModeProposal Mode = "proposal"
)

type Capabilities string

const (
	CapabilitiesCoinbaseTxn    Capabilities = "coinbasetxn"
	CapabilitiesCoinbaseAux    Capabilities = "coinbaseaux"
	CapabilitiesWorkID         Capabilities = "workid"
	CapabilitiesCoinbaseAppend Capabilities = "coinbase/append"
)

type Rules string

const (
	RulesSegWit  Rules = "segwit"
	RulesSegWitN Rules = "!segwit"
	RulesCSV     Rules = "csv"
	RulesTaproot Rules = "taproot"
)

type GetBlockTemplateResult struct {
	Capabilities             []Capabilities `json:"capabilities"`
	Version                  uint32         `json:"version"`
	Rules                    []Rules        `json:"rules"`
	VBavailable              interface{}    `json:"vbavailable"`
	VBrequired               int            `json:"vbrequired"`
	PreviousBlockHash        common.Hash    `json:"previousblockhash"`
	Transactions             []Transaction  `json:"transactions"`
	CoinbaseAux              interface{}    `json:"coinbaseaux"`
	LongPollID               string         `json:"longpollid"`
	Target                   string         `json:"target"`
	MinTime                  int            `json:"mintime"`
	Mutable                  []string       `json:"mutable"`
	NonceRange               string         `json:"noncerange"`
	SigOpLimit               int            `json:"sigoplimit"`
	WeightLimit              int            `json:"weightlimit"`
	CurlTime                 int            `json:"curtime"`
	Bits                     string         `json:"bits"`
	Height                   int            `json:"height"`
	DefaultWitnessCommitment string         `json:"default_witness_commitment"`
}

type Transaction struct {
	Data    string      `json:"data"`
	TxID    common.Hash `json:"txid"`
	Hash    common.Hash `json:"hash"`
	Depends []int       `json:"depends"`
	Fee     int         `json:"fee"`
	SigOps  int         `json:"sigops"`
	Weight  int         `json:"weight"`
}

func CreateGetBlockTemplateParams() *GetBlockTemplateParams {
	return &GetBlockTemplateParams{
		Mode: ModeTemplate,
		Capabilities: []Capabilities{
			CapabilitiesCoinbaseTxn,
			CapabilitiesWorkID,
			CapabilitiesCoinbaseAppend,
		},
		Rules: []Rules{
			RulesSegWit,
		},
	}
}

func UnmarshalGetBlockTemplateResult(data []byte) (*GetBlockTemplateResult, error) {
	var result GetBlockTemplateResult
	err := json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
