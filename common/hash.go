package common

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
)

const (
	HashLength = 32
)

type Hash [HashLength]byte

func NilHash() Hash {
	return BytesToHash(make([]byte, 32))
}

func (h Hash) String() string {
	rev := ReverseBytes(h.Bytes())
	return hex.EncodeToString(rev)
}

// Str returns the hash as a string without reversing the byte order.
// Used for things like sending jsonrpc requests where the byte order matters.
func (h Hash) Str() string {
	return hex.EncodeToString(h[:])
}

func (h Hash) Bytes() []byte {
	return h[:]
}

func (h *Hash) CloneBytes() []byte {
	newHash := make([]byte, HashLength)
	copy(newHash, h[:])
	return newHash
}

func (h *Hash) SetBytes(newHash []byte) error {
	if len(newHash) != HashLength {
		return fmt.Errorf("invalid hash length: got %d, expected %d", len(newHash), HashLength)
	}
	copy(h[:], newHash)
	return nil
}

func (h *Hash) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	if len(s) != 2*HashLength {
		return fmt.Errorf("invalid hash length: got %d, expected %d", len(s), 2*HashLength)
	}
	b, err := hex.DecodeString(s)
	if err != nil {
		return err
	}
	rev := ReverseBytes(b)
	return h.SetBytes(rev)
}

func BytesToHash(b []byte) Hash {
	var h Hash
	copy(h[:], b)
	return h
}

func HashFromBytes(b []byte) Hash {
	var h Hash
	copy(h[:], b)
	return h
}

// StringToHash converts a hex string to a Hash. It should not reverse the byte order.
func StringToHash(s string) Hash {
	bytes, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return BytesToHash(bytes)
}
