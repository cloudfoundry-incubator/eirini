package cmd_test

import (
	"os"

	"code.cloudfoundry.org/eirini"
	"code.cloudfoundry.org/eirini/tests/integration"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("EiriniController", func() {
	var (
		config         *eirini.ControllerConfig
		configFilePath string
		session        *gexec.Session
	)

	BeforeEach(func() {
		config = integration.DefaultControllerConfig(fixture.Namespace)
	})

	JustBeforeEach(func() {
		session, configFilePath = eiriniBins.EiriniController.Run(config)
	})

	AfterEach(func() {
		if configFilePath != "" {
			Expect(os.Remove(configFilePath)).To(Succeed())
		}
		if session != nil {
			Eventually(session.Kill()).Should(gexec.Exit())
		}
	})

	It("should be able to start properly", func() {
		Consistently(session).ShouldNot(gexec.Exit())
	})
})
