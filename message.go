package grapes

import (
	"encoding/binary"
	"errors"
)

// A MessageType is a GRAPES message (packet) type.
// It should be used to identify the type of a message content for decoding.
type MessageType int

// Message types
const (
	TypeTopology  MessageType = 0x10
	TypeChunk                 = 0x11
	TypeSignaling             = 0x12
	TypeTman                  = 0x13
)

// ErrMessageTooShort is returned when ParseMessage reads a byte slice shorter
// than the message headers.
var ErrMessageTooShort = errors.New("grapes: message too short")

// ErrInvalidMessageType is returned when ParseMessage finds a bad message type
// in the headers.
var ErrInvalidMessageType = errors.New("grapes: invalid message type")

// Message is the structure enveloping GRAPES messages (chunks, topology, etc.)
type Message struct {
	Type          MessageType
	TransactionID uint16
	Content       []byte
}

// ParseMessage returns a GRAPES message parsed from byte slice.
// The slice is consumed entirely, therefore is must contain only one message.
//
func ParseMessage(data []byte) (*Message, error) {
	const headerLength = 3
	if len(data) < headerLength {
		return nil, ErrMessageTooShort
	}
	msg := Message{}
	switch data[0] {
	case 0x10:
		msg.Type = TypeTopology
	case 0x11:
		msg.Type = TypeChunk
	case 0x12:
		msg.Type = TypeSignaling
	case 0x13:
		msg.Type = TypeTman
	default:
		return nil, ErrInvalidMessageType
	}
	msg.TransactionID = binary.BigEndian.Uint16(data[1:3])
	return &msg, nil
}
