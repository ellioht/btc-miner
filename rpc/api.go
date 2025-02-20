package rpc

import (
	"encoding/json"
	"github.com/ellioht/btc-miner/common"
)

type ApiImpl interface {
	GetBlock(blockHash common.Hash, verbosity ...Verbosity) (*GetBlockResult, error)
	GetBlockHash(blockNumber uint64) (*common.Hash, error)
	GetBlockTemplate(params *GetBlockTemplateParams) (*GetBlockTemplateResult, error)
	GetRawTransaction(txId common.Hash, verbose bool) (*TxFullOrHex, error)
	Help(method Method) (string, error)
}

type ApiClient struct {
	client *Client
}

func NewApiClient(url string) *ApiClient {
	return &ApiClient{
		client: NewClient(url),
	}
}

func (a *ApiClient) GetBlock(blockHash common.Hash, verbosity ...Verbosity) (*GetBlockResult, error) {
	v := SerializedBlock
	if len(verbosity) > 0 {
		if len(verbosity) > 1 {
			return nil, ErrInvalidVerbosity
		}
		if verbosity[0] > 2 {
			return nil, ErrInvalidVerbosity
		}
		v = verbosity[0]
	}

	req, err := a.client.CreateRequestPayload(MethodGetBlock, blockHash.Str(), v)
	if err != nil {
		return nil, err
	}

	res, err := a.client.HandleRequest(req)
	if err != nil {
		return nil, err
	}

	return UnmarshalGetBlockResult(res.Result)
}

func (a *ApiClient) GetBlockHash(blockNumber uint64) (common.Hash, error) {
	req, err := a.client.CreateRequestPayload(MethodGetBlockHash, blockNumber)
	if err != nil {
		return common.Hash{}, err
	}

	res, err := a.client.HandleRequest(req)
	if err != nil {
		return common.Hash{}, err
	}

	var hashStr string
	if err = json.Unmarshal(res.Result, &hashStr); err != nil {
		return common.Hash{}, err
	}

	return common.StringToHash(hashStr), nil
}

func (a *ApiClient) GetBlockTemplate(params *GetBlockTemplateParams) (*GetBlockTemplateResult, error) {
	req, err := a.client.CreateRequestPayload(MethodGetBlockTemplate, params)
	if err != nil {
		return nil, err
	}

	res, err := a.client.HandleRequest(req)
	if err != nil {
		return nil, err
	}

	return UnmarshalGetBlockTemplateResult(res.Result)
}

func (a *ApiClient) Help(method Method) (string, error) {
	req, err := a.client.CreateRequestPayload(MethodHelp, method)
	if err != nil {
		return "", err
	}

	res, err := a.client.HandleRequest(req)
	if err != nil {
		return "", err
	}

	var help string
	if err = json.Unmarshal(res.Result, &help); err != nil {
		return "", err
	}

	return help, nil
}

func (a *ApiClient) GetRawTransaction(txId common.Hash, verbose bool) (*TxFullOrHex, error) {
	if !verbose {
		req, err := a.client.CreateRequestPayload(MethodGetRawTransaction, txId.Str(), false)
		if err != nil {
			return nil, err
		}

		res, err := a.client.HandleRequest(req)
		if err != nil {
			return nil, err
		}

		var rawTx string
		if err = json.Unmarshal(res.Result, &rawTx); err != nil {
			return nil, err
		}

		return &TxFullOrHex{TxHex: &rawTx}, nil
	} else {
		req, err := a.client.CreateRequestPayload(MethodGetRawTransaction, txId.Str(), true)
		if err != nil {
			return nil, err
		}

		res, err := a.client.HandleRequest(req)
		if err != nil {
			return nil, err
		}

		full, err := UnmarshalGetRawTransactionResult(res.Result)
		if err != nil {
			return nil, err
		}

		return &TxFullOrHex{TxFull: full}, nil
	}
}
