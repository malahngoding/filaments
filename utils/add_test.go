package utils

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Add Test", func() {
	It("return correct value", func() {
		commands := Add(1, 2)
		Expect(commands).To(Equal(3))
	})
})
