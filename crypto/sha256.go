package crypto

import (
	"crypto/sha256"
	"github.com/ellioht/btc-miner/common"
)

// HashData takes a byte slice and returns a double sha256 hash of the data.
func HashData(data []byte) []byte {
	first := sha256.Sum256(data)
	second := sha256.Sum256(first[:])
	return second[:]
}

// HashPair takes two byte slices and returns a double sha256 hash of the two slices concatenated.
func HashPair(left, right []byte) []byte {
	var both [common.HashLength * 2]byte
	copy(both[:common.HashLength], left[:])
	copy(both[common.HashLength:], right[:])
	return HashData(both[:])
}
