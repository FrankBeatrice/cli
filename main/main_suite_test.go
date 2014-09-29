package main_test

import (
	"os"
	"os/exec"
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestMain(t *testing.T) {
	RegisterFailHandler(Fail)
	dir, err := os.Getwd()
	Expect(err).NotTo(HaveOccurred())

	exPath := filepath.Join(dir, "..", "fixtures", "plugins", "test")
	srcPath := filepath.Join(dir, "..", "fixtures", "plugins", "test.go")
	cmd := exec.Command("go", "build", "-o", exPath, srcPath)
	println("Working dir: ", dir)
	println("Executable Path: ", exPath)
	println("SrcPath: ", srcPath)
	err = cmd.Run()
	if err != nil {
		println(err.Error())
	}
	Expect(err).NotTo(HaveOccurred())

	cmd = exec.Command("go", "build", "-o", filepath.Join(dir, "..", "fixtures", "plugins", "plugin2"), filepath.Join(dir, "..", "fixtures", "plugins", "plugin2.go"))
	err = cmd.Run()
	Expect(err).NotTo(HaveOccurred())

	RunSpecs(t, "Main Suite")
}