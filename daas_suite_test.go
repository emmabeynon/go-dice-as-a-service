package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestDaas(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Daas Suite")
}
