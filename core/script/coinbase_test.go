package script

import (
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecodeCoinbaseScript(t *testing.T) {
	coinbaseScript, err := NewCoinbaseScriptBytes("03877d0d122f5669614254432f5472757374706f6f6c2f2cfabe6d6d3fdbecfd646e97d92702be5ea9d2453bb723a80e771d2d03fa8008e10b06db50100000000000000010301b1b05a59f6c9c96f415fa0fbc3a0000000000")
	assert.NoError(t, err)

	height, messages, err := DecodeCoinbaseScript(coinbaseScript)
	assert.NoError(t, err)

	assert.Equal(t, int64(884103), height)
	assert.Equal(t, len(messages), 7)
	assert.Equal(t, "/ViaBTC/Trustpool/", messages[0].String())
	assert.Equal(t, "fabe6d6d3fdbecfd646e97d92702be5ea9d2453bb723a80e771d2d03fa8008e10b06db501000000000000000", messages[1].Hex())
	assert.Equal(t, "301b1b05a59f6c9c96f415fa0fbc3a00", messages[2].Hex())
	assert.Equal(t, "", messages[3].Hex())
	assert.Equal(t, "", messages[4].Hex())
	assert.Equal(t, "", messages[5].Hex())
	assert.Equal(t, "", messages[6].Hex())
}

func TestEncodeCoinbaseScript(t *testing.T) {
	coinbaseScriptString := "03877d0d122f5669614254432f5472757374706f6f6c2f2cfabe6d6d3fdbecfd646e97d92702be5ea9d2453bb723a80e771d2d03fa8008e10b06db50100000000000000010301b1b05a59f6c9c96f415fa0fbc3a0000000000"

	height := int64(884103)

	var messages [][]byte
	message1 := "/ViaBTC/Trustpool/"
	messages = append(messages, []byte(message1))

	message2, err := hex.DecodeString("fabe6d6d3fdbecfd646e97d92702be5ea9d2453bb723a80e771d2d03fa8008e10b06db501000000000000000")
	assert.NoError(t, err)
	messages = append(messages, message2)

	message3, err := hex.DecodeString("301b1b05a59f6c9c96f415fa0fbc3a00")
	assert.NoError(t, err)
	messages = append(messages, message3)

	messages = append(messages, []byte(""))
	messages = append(messages, []byte(""))
	messages = append(messages, []byte(""))
	messages = append(messages, []byte(""))

	coinbaseScript, err := EncodeCoinbaseScript(height, messages...)
	assert.NoError(t, err)

	assert.Equal(t, coinbaseScriptString, CoinbaseScriptToString(coinbaseScript))

	decodedHeight, decodedMessages, err := DecodeCoinbaseScript(coinbaseScript)
	assert.NoError(t, err)

	assert.Equal(t, height, decodedHeight)

	assert.Equal(t, len(messages), len(decodedMessages))
	assert.Equal(t, message1, decodedMessages[0].String())
	assert.Equal(t, hex.EncodeToString(message2), decodedMessages[1].Hex())
	assert.Equal(t, hex.EncodeToString(message3), decodedMessages[2].Hex())
	assert.Equal(t, "", decodedMessages[3].Hex())
	assert.Equal(t, "", decodedMessages[4].Hex())
	assert.Equal(t, "", decodedMessages[5].Hex())
	assert.Equal(t, "", decodedMessages[6].Hex())
}
