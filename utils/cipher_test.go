package utils

import (
	"github.com/malahngoding/filaments/config"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Cipher Test", func() {
	It("return encrypted value and then decrypt the value", func() {
		key := []byte(config.InsteadToken())
		encrypted, err := Encrypt(key, "combined")
		if err != nil {
			Expect(err).NotTo(HaveOccurred())
		}
		Expect(encrypted).To(Not(Equal("combined")))
		decrypted, err := Decrypt(key, encrypted)
		if err != nil {
			Expect(err).NotTo(HaveOccurred())
		}
		Expect(decrypted).To(Equal("combined"))
	})
})
