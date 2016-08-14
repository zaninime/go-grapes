package grapes_test

import (
	"bytes"

	. "github.com/zaninime/go-grapes"
)

var _ = Describe("RTP", func() {
	Describe("parsing", func() {
		It("should parse a valid envelope successfully", func() {
			pkt := []byte{0, 4, 128, 0xde, 0xad, 0xbe, 0xef}
			parsed, consumed, err := ParseRTPEnvelope(pkt)

			Expect(err).NotTo(HaveOccurred())
			Expect(consumed).To(BeEquivalentTo(len(pkt)), "Consumed bytes")

			Expect(parsed.StreamID).To(BeEquivalentTo(128), "Stream ID")
			Expect(bytes.Equal(parsed.Content, []byte{0xde, 0xad, 0xbe, 0xef})).To(BeTrue(), "Content")
		})

		It("should return an error when there aren't enough bytes", func() {
			_, _, err := ParseRTPEnvelope([]byte{0})
			Expect(err).To(MatchError(ErrChunkTooShort))

			_, _, err = ParseRTPEnvelope([]byte{0, 4, 128, 1, 2, 3})
			Expect(err).To(MatchError(ErrChunkTooShort))
		})
	})
})
