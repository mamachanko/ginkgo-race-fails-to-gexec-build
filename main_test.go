package main_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
	"os/exec"
)

var _ = Describe("Main", func() {

	It("says Orkbork", func() {
		command := exec.Command(pathToCLI)
		session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
		Expect(err).ShouldNot(HaveOccurred())
		Eventually(session).Should(gexec.Exit(0))
		Eventually(session).Should(gbytes.Say("Orkbork"))
	})

})
