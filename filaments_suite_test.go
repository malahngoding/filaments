package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestFilaments(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Filaments Suite")
}
