package e2e_test

import (
	"os"
	"path/filepath"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestE2E(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AI Explorer E2E Suite")
}

func createDir(path string) {
	Expect(os.MkdirAll(path, os.ModePerm)).
		To(Succeed(), "Failed to create directory: %s", path)
	GinkgoWriter.Printf("Created directory: %s\n", path)
}

var _ = BeforeSuite(func() {

	By("Creating output directory")
	createDir(filepath.Join("..", ".build", "output"))
})
