package main_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
	"testing"
)

func TestGinkgoRaceFailsToGexecBuild(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GinkgoRaceFailsToGexecBuild Suite")
}

var pathToCLI string

var _ = BeforeSuite(func() {
	var err error
	pathToCLI, err = gexec.Build("github.com/mamachanko/ginkgo-race-fails-to-gexec-build")
	Expect(err).ShouldNot(HaveOccurred())
})

var _ = AfterSuite(func() {
	gexec.KillAndWait()
	gexec.CleanupBuildArtifacts()
})
