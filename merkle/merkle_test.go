package merkle

import (
	"encoding/json"
	"github.com/ellioht/btc-miner/common"
	"github.com/ellioht/btc-miner/rpc"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"testing"
)

func TestComputeMerkleRoot(t *testing.T) {
	file, err := os.Open("../rpc/test_data/getblock-884103.json")
	assert.NoError(t, err)
	defer file.Close()

	data, err := io.ReadAll(file)
	assert.NoError(t, err)

	var block rpc.GetBlockResult
	assert.NoError(t, json.Unmarshal(data, &block))

	txIds := make([]common.Hash, len(block.Tx))
	for i, tx := range block.Tx {
		txIds[i] = tx.TxId
	}

	tree := NewTreeFromHashes(txIds)
	calculatedMerkleRoot := tree.ComputeMerkleRoot().String()

	expectedMerkleRoot := block.MerkleRoot

	assert.Equal(t, expectedMerkleRoot, calculatedMerkleRoot)
}
