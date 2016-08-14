package grapes_test

import (
	"bytes"

	. "github.com/zaninime/go-grapes"
)

var _ = Describe("Chunk", func() {
	Describe("parsing", func() {
		It("should parse a valid packet successfully", func() {
			pkt := []byte{0xba, 0xdd, 0xca, 0xfe, 0, 1, 2, 3, 4, 5, 6, 7, 0, 0, 0, 2, 0, 0, 0, 2, 0xde, 0xad, 0xc0, 0xde}
			parsed, consumed, err := ParseChunk(pkt)
			Expect(err).NotTo(HaveOccurred())
			Expect(consumed).To(BeEquivalentTo(len(pkt)), "Consumed bytes")

			Expect(parsed.ID).To(BeEquivalentTo(0xbaddcafe), "Chunk ID")
			Expect(parsed.Timestamp).To(BeEquivalentTo(0x01020304050607), "Timestamp")
			Expect(bytes.Equal(parsed.Content, []byte{0xde, 0xad})).To(BeTrue(), "Content")
			Expect(bytes.Equal(parsed.Attributes, []byte{0xc0, 0xde})).To(BeTrue(), "Attributes")
		})
		It("should return an error when packet is too short", func() {
			_, _, err := ParseChunk([]byte{0x00})
			Expect(err).To(MatchError(ErrChunkTooShort))

			_, _, err = ParseChunk([]byte{0xca, 0xfe, 0, 1, 2, 3, 4, 5, 6, 7, 0, 0, 0, 5, 0, 0, 0, 5, 0xde, 0xad, 0xc0})
			Expect(err).To(MatchError(ErrChunkTooShort))
		})
	})
})
