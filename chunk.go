package grapes

import (
	"encoding/binary"
	"errors"
)

// ErrChunkTooShort is returned when a parser wants to consume
// more bytes than the chunk contains.
var ErrChunkTooShort = errors.New("grapes: chunk too short")

// Chunk is the data structure of the packets traveling between peers
// in the P2P network containing audio/video data.
// Note that chunk may be incapsulated in other protocols (like ML)
// that multiplex the socket functionality (ie. signaling).
type Chunk struct {
	ID         uint32
	Timestamp  uint64
	Content    []byte
	Attributes []byte
}

// ParseChunk returns a chunk parsed from a data slice. It returns how many
// bytes of the slice have been consumed. If the byte slice is too short,
// shorter than the headers or the headers plus the declared data size,
// ErrChunkTooShort is returned.
func ParseChunk(data []byte) (*Chunk, uint, error) {
	const headerLength = 20
	if len(data) < headerLength {
		return nil, 0, ErrChunkTooShort
	}
	chunk := Chunk{}
	chunk.ID = binary.BigEndian.Uint32(data[0:4])
	chunk.Timestamp = binary.BigEndian.Uint64(data[4:12])
	contentSize := binary.BigEndian.Uint32(data[12:16])
	attributesSize := binary.BigEndian.Uint32(data[16:20])
	if len(data) < int(headerLength+contentSize+attributesSize) {
		return nil, 0, ErrChunkTooShort
	}
	chunk.Content = data[headerLength : headerLength+contentSize]
	chunk.Attributes = data[headerLength+contentSize : headerLength+contentSize+attributesSize]
	return &chunk, uint(headerLength + contentSize + attributesSize), nil
}
