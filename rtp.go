package grapes

import "errors"

// RTPEnvelope defines the container of a RTP packet serialized in a chunk.
// A chunk should contain one or more of these structures, when created
// with the RTP chunkiser.
type RTPEnvelope struct {
	// DataLength is the length of the field content, declared in the header.
	DataLength uint16

	// StreamID identifies to which stream the following data belongs to.
	// The RTP chunkiser may create one chunk with RTP packets coming from
	// different streams.
	StreamID uint8

	// Content is the actual RTP/RTCP packet belonging to stream StreamID.
	Content []byte
}

// ErrChunkTooShort is returned when an envelope parser wants to consume
// more bytes than the chunk contains.
var ErrChunkTooShort = errors.New("grapes: chunk too short")

// ParseChunkRTPEnvelope returns an RTPEnvelope parsed from a chunk content
// pointer. The returned int is the length of the data consumed inside
// dataPointer. ErrChunkTooShort is returned if the dataPointer doesn't contain
// enough data to successfully parse the envelope.
func ParseChunkRTPEnvelope(dataPointer []byte) (*RTPEnvelope, uint, error) {
	return nil, 0, nil
}
