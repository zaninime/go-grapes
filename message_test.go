package grapes_test

import (
	. "github.com/zaninime/go-grapes"
)

var _ = Describe("Message", func() {
	Describe("parsing", func() {

		It("should return the right message type", func() {
			pkt1, _ := ParseMessage([]byte{0x10, 0x11, 0x12, 0x13})
			Expect(pkt1.Type).To(BeEquivalentTo(TypeTopology))

			pkt2, _ := ParseMessage([]byte{0x11, 0x11, 0x12, 0x13})
			Expect(pkt2.Type).To(BeEquivalentTo(TypeChunk))

			pkt3, _ := ParseMessage([]byte{0x12, 0x11, 0x12, 0x13})
			Expect(pkt3.Type).To(BeEquivalentTo(TypeSignaling))

			pkt4, _ := ParseMessage([]byte{0x13, 0x11, 0x12, 0x13})
			Expect(pkt4.Type).To(BeEquivalentTo(TypeTman))
		})

		It("should not return an error with a valid message", func() {
			_, err := ParseMessage([]byte{0x13, 0x11, 0x12, 0x13})
			Expect(err).NotTo(HaveOccurred())
		})

		It("should return an error with an invalid message type", func() {
			_, err := ParseMessage([]byte{0xff, 0x11, 0x12, 0x13})
			Expect(err).To(MatchError(ErrInvalidMessageType))
		})

		It("should return an error with any message shorter than the header length", func() {
			_, err := ParseMessage([]byte{0x00})
			Expect(err).To(MatchError(ErrMessageTooShort))
		})

		It("should return the right transaction ID", func() {
			pkt, _ := ParseMessage([]byte{0x10, 0x11, 0x12, 0x13})
			Expect(pkt.TransactionID).To(BeEquivalentTo(0x11<<8 | 0x12))
		})

		It("should return the content", func() {
			pkt, _ := ParseMessage([]byte{0x10, 0x11, 0x12, 0x13, 0x14, 0x15})
			Expect(len(pkt.Content)).To(Equal(3))
			Expect(pkt.Content[0]).To(BeEquivalentTo(0x13))
		})
	})
})
