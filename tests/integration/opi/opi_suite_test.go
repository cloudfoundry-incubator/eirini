package opi_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	"code.cloudfoundry.org/eirini"
	"code.cloudfoundry.org/eirini/models/cf"
	"code.cloudfoundry.org/eirini/tests"
	"code.cloudfoundry.org/eirini/tests/integration"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/client-go/rest"
)

func TestOpi(t *testing.T) {
	SetDefaultEventuallyTimeout(time.Minute)
	RegisterFailHandler(Fail)
	RunSpecs(t, "Opi Suite")
}

const secretName = "certs-secret"

var (
	eiriniBins           integration.EiriniBinaries
	fixture              *tests.Fixture
	httpClient           *http.Client
	eiriniConfigFilePath string
	session              *gexec.Session
	url                  string
	apiConfig            *eirini.APIConfig
	opiEnvOverride       []string
)

var _ = SynchronizedBeforeSuite(func() []byte {
	eiriniBins = integration.NewEiriniBinaries()
	eiriniBins.OPI.Build()

	data, err := json.Marshal(eiriniBins)
	Expect(err).NotTo(HaveOccurred())

	return data
}, func(data []byte) {
	err := json.Unmarshal(data, &eiriniBins)
	Expect(err).NotTo(HaveOccurred())

	fixture = tests.NewFixture(GinkgoWriter)
})

var _ = SynchronizedAfterSuite(func() {
	fixture.Destroy()
}, func() {
	eiriniBins.TearDown()
})

var _ = BeforeEach(func() {
	fixture.SetUp()

	Expect(integration.CreateEmptySecret(fixture.Namespace, secretName, fixture.Clientset)).To(Succeed())

	var err error
	httpClient, err = integration.MakeTestHTTPClient()
	Expect(err).ToNot(HaveOccurred())

	apiConfig = integration.DefaultAPIConfig(fixture.Namespace, fixture.NextAvailablePort())
	opiEnvOverride = []string{}
})

var _ = JustBeforeEach(func() {
	session, eiriniConfigFilePath = eiriniBins.OPI.Run(apiConfig, opiEnvOverride...)

	url = fmt.Sprintf("https://localhost:%d/", apiConfig.TLSPort)
	Eventually(func() error {
		_, getErr := httpClient.Get(url)

		return getErr
	}).Should(Succeed())
})

var _ = AfterEach(func() {
	Expect(os.Remove(eiriniConfigFilePath)).To(Succeed())
	session.Kill()

	fixture.TearDown()
})

func desireLRP(httpClient rest.HTTPClient, url string, lrpRequest cf.DesireLRPRequest) *http.Response {
	body, err := json.Marshal(lrpRequest)
	Expect(err).NotTo(HaveOccurred())
	desireLrpReq, err := http.NewRequest("PUT", fmt.Sprintf("%s/apps/%s", url, lrpRequest.GUID), bytes.NewReader(body))
	Expect(err).NotTo(HaveOccurred())
	response, err := httpClient.Do(desireLrpReq)
	Expect(err).NotTo(HaveOccurred())

	return response
}
