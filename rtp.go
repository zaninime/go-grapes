package grapes

import "encoding/binary"

// RTPEnvelope defines the container of a RTP packet serialized in a chunk.
// A chunk should contain one or more of these structures, when created
// with the RTP chunkiser.
type RTPEnvelope struct {
	StreamID uint8  // ID of RTP stream the packet belongs to
	Content  []byte // Actual RTP/RTCP packet belonging to stream StreamID.
}

// ParseRTPEnvelope returns an RTPEnvelope parsed from a chunk content
// pointer. The returned int is the length of the data consumed inside
// dataPointer. ErrChunkTooShort is returned if the dataPointer doesn't contain
// enough data to successfully parse the envelope.
func ParseRTPEnvelope(data []byte) (*RTPEnvelope, uint, error) {
	const headerLength = 3
	if len(data) < headerLength {
		return nil, 0, ErrChunkTooShort
	}
	contentSize := binary.BigEndian.Uint16(data[0:2])
	if len(data) < int(headerLength+contentSize) {
		return nil, 0, ErrChunkTooShort
	}
	envelope := RTPEnvelope{}
	envelope.StreamID = data[2]
	envelope.Content = data[3 : 3+contentSize]
	return &envelope, uint(headerLength + contentSize), nil
}
