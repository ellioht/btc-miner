package cpu

import (
	"github.com/ellioht/btc-miner/common"
	"github.com/ellioht/btc-miner/core"
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestMiner(t *testing.T) {
	miner := NewCpuMiner()
	if miner.IsStarted() {
		t.Errorf("expected false, got true")
	}

	if miner.IsMining() {
		t.Errorf("expected false, got true")
	}

	miner.Start()

	if !miner.IsStarted() {
		t.Errorf("expected true, got false")
	}

	miner.Stop()

	if miner.IsStarted() {
		t.Errorf("expected false, got true")
	}
}

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
