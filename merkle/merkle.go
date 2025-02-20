package merkle

import (
	"github.com/ellioht/btc-miner/common"
	"github.com/ellioht/btc-miner/crypto"
)

type Tree struct {
	roots  []common.Hash
	leaves uint64
}

func NewTreeFromHashes(hashes []common.Hash) *Tree {
	roots := make([]common.Hash, 0)
	for _, hash := range hashes {
		roots = append(roots, hash)
	}
	return &Tree{
		roots:  roots,
		leaves: uint64(len(hashes)),
	}
}

func (t *Tree) ComputeMerkleRoot() common.Hash {
	if t.leaves == 0 {
		return common.Hash{}
	}

	for len(t.roots) > 1 {
		if len(t.roots)%2 != 0 {
			t.roots = append(t.roots, t.roots[len(t.roots)-1])
		}

		var newLevel []common.Hash
		for i := 0; i < len(t.roots); i += 2 {
			newHash := common.BytesToHash(crypto.HashPair(t.roots[i].Bytes(), t.roots[i+1].Bytes()))
			newLevel = append(newLevel, newHash)
		}
		t.roots = newLevel
	}

	return t.roots[0]
}
