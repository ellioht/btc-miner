package common

import (
	"encoding/json"
	"testing"
)

// StringToHash converts a big-endian hex string to a Hash.
// We keep it in big-endian format because that's how Bitcoin displays hashes.
func TestEndianStringToHash(t *testing.T) {
	testHash := "0000000000000000000028f749b7b00840f16f9179d4398b57560a4938aa6159"
	hash := StringToHash(testHash)

	expected := Hash{
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x28, 0xF7,
		0x49, 0xB7, 0xB0, 0x08,
		0x40, 0xF1, 0x6F, 0x91,
		0x79, 0xD4, 0x39, 0x8B,
		0x57, 0x56, 0x0A, 0x49,
		0x38, 0xAA, 0x61, 0x59,
	}

	if hash != expected {
		t.Errorf("expected %v, got %v", expected, hash)
	}
}

// TestEndianMarshalJSON tests the MarshalJSON method of the Hash type.
// We expect the Json Hash to be in big-endian format, so we reverse the bytes.
func TestEndianUnmarshalJSON(t *testing.T) {
	testJson := `{
  "hash": "0000000000000000000028f749b7b00840f16f9179d4398b57560a4938aa6159"
}`
	var testStruct struct {
		Hash Hash `json:"hash"`
	}

	if err := json.Unmarshal([]byte(testJson), &testStruct); err != nil {
		t.Errorf("expected nil, got %v", err)
	}

	expected := Hash{
		0x59, 0x61, 0xAA, 0x38,
		0x49, 0x0A, 0x56, 0x57,
		0x8B, 0x39, 0xD4, 0x79,
		0x91, 0x6F, 0xF1, 0x40,
		0x08, 0xB0, 0xB7, 0x49,
		0xF7, 0x28, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
	}

	if testStruct.Hash != expected {
		t.Errorf("expected %v, got %v", expected, testStruct.Hash)
	}
}

// TestEndianString tests the String method of the Hash type.
// Bitcoin displays hashes in big-endian format, but stores them in little-endian format.
func TestEndianString(t *testing.T) {
	testHashLittleEndian := "5961aa38490a56578b39d479916ff14008b0b749f728000000000000000000"
	hash := StringToHash(testHashLittleEndian)
	hashString := hash.String()
	expected := "0000000000000000000028f749b7b00840f16f9179d4398b57560a4938aa6159"
	if hashString != expected {
		t.Errorf("expected %v, got %v", expected, hashString)
	}
}
