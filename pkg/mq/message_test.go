package mq

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMessage(t *testing.T) {
	type msg struct {
		IdA int
		IdB int
	}

	encodeFunc := func(source interface{}) ([]byte, error) {
		data := source.(*msg)
		return []byte{byte(data.IdA), byte(data.IdB)}, nil
	}
	decodeFunc := func(target interface{}, message *Message) error {
		data := target.(*msg)
		data.IdA = int(message.value[0])
		data.IdB = int(message.value[1])
		return nil
	}

	parser := NewMessageParser(encodeFunc, decodeFunc)
	msg1 := msg{IdA: 1, IdB: 2}
	mqMsg1, err := parser.Encode(&msg1)
	require.NoError(t, err)
	require.Equal(t, []byte{1, 2}, mqMsg1.value)

	decodedMsg := &msg{}
	err = parser.Decode(decodedMsg, mqMsg1)
	require.NoError(t, err)
	require.Equal(t, msg1.IdA, decodedMsg.IdA)
	require.Equal(t, msg1.IdB, decodedMsg.IdB)
}
