package cpu

import (
	"encoding/json"
	"github.com/ellioht/btc-miner/common"
	"github.com/ellioht/btc-miner/core"
	"github.com/ellioht/btc-miner/merkle"
	"github.com/ellioht/btc-miner/rpc"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"testing"
)

func TestHashHeader(t *testing.T) {
	version := uint32(549453824)
	prevBlockHash := common.StringToHash("00000000000000000000da64a1be92cb508b71344ec8f08092fe19ff7501ea5e").Reverse()
	merkleRoot := common.StringToHash("593b3d441cfbfd23b8b9e02af83317a413433cd201c51264f3df33a6560f8962").Reverse()
	timestamp := uint32(1739743264)
	bits, err := common.CompactToUint32("17027726")
	assert.NoError(t, err)
	nonce := uint32(4122426899)
	header := core.NewHeader(version, prevBlockHash, merkleRoot, timestamp, bits, nonce)

	hash := header.Hash()
	assert.Equal(t, "00000000000000000000c548a1214727445027fff8fb5ef81a8746c3365e76b3", hash.String())
}

func TestIsSolved(t *testing.T) {
	version := uint32(549453824)
	prevBlockHash := common.StringToHash("00000000000000000000da64a1be92cb508b71344ec8f08092fe19ff7501ea5e").Reverse()
	merkleRoot := common.StringToHash("593b3d441cfbfd23b8b9e02af83317a413433cd201c51264f3df33a6560f8962").Reverse()
	timestamp := uint32(1739743264)
	bits, err := common.CompactToUint32("17027726")
	assert.NoError(t, err)
	invalidNonce := uint32(1)
	header := core.NewHeader(version, prevBlockHash, merkleRoot, timestamp, bits, invalidNonce)

	assert.True(t, !IsSolved(header))

	validNonce := uint32(4122426899)
	header.UpdateNonce(validNonce)

	assert.True(t, IsSolved(header))
}

func TestHashHeaderAndSolveBlock(t *testing.T) {
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

	tree := merkle.NewTreeFromHashes(txIds)
	calculatedMerkleRoot := tree.ComputeMerkleRoot()

	expectedMerkleRoot := block.MerkleRoot

	assert.Equal(t, expectedMerkleRoot, calculatedMerkleRoot.String())

	version := uint32(549453824)
	prevBlockHash := common.StringToHash("00000000000000000000da64a1be92cb508b71344ec8f08092fe19ff7501ea5e").Reverse()
	timestamp := uint32(1739743264)
	bits, err := common.CompactToUint32("17027726")
	assert.NoError(t, err)
	nonce := uint32(4122426899)
	header := core.NewHeader(version, prevBlockHash, calculatedMerkleRoot, timestamp, bits, nonce)

	hash := header.Hash()
	assert.Equal(t, "00000000000000000000c548a1214727445027fff8fb5ef81a8746c3365e76b3", hash.String())

	assert.True(t, IsSolved(header))
}
