package mq

type Message struct {
	value []byte
}

type encodeMessageFunc func(source interface{}) ([]byte, error)
type decodeMessageFunc func(target interface{}, message *Message) error

type MessageParser struct {
	encode encodeMessageFunc
	decode decodeMessageFunc
}

func NewMessageParser(encode encodeMessageFunc, decode decodeMessageFunc) *MessageParser {
	return &MessageParser{
		encode: encode,
		decode: decode,
	}
}

func (p *MessageParser) Encode(source interface{}) (*Message, error) {
	value, err := p.encode(source)
	if err != nil {
		return nil, err
	}
	return &Message{
		value: value,
	}, nil
}

func (p *MessageParser) Decode(target interface{}, message *Message) error {
	return p.decode(target, message)
}
