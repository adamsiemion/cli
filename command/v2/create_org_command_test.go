package v2_test

import (
	"code.cloudfoundry.org/cli/command/commandfakes"
	"code.cloudfoundry.org/cli/command/flag"
	. "code.cloudfoundry.org/cli/command/v2"
	"code.cloudfoundry.org/cli/command/v2/v2fakes"
	"code.cloudfoundry.org/cli/util/ui"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
)

var _ = Describe("CreateOrgCommand", func() {
	var (
		fakeConfig *commandfakes.FakeConfig
		fakeActor  *v2fakes.FakeCreateOrgActor
		testUI     *ui.UI

		cmd CreateOrgCommand

		executeErr error
	)

	BeforeEach(func() {
		testUI = ui.NewTestUI(nil, NewBuffer(), NewBuffer())
		fakeConfig = new(commandfakes.FakeConfig)
		fakeActor = new(v2fakes.FakeCreateOrgActor)

		cmd = CreateOrgCommand{
			UI:           testUI,
			Config:       fakeConfig,
			Actor:        fakeActor,
			RequiredArgs: flag.Organization{Organization: "some-org"},
		}
	})

	JustBeforeEach(func() {
		executeErr = cmd.Execute(nil)
	})

	It("creates the org", func() {
		Expect(executeErr).ToNot(HaveOccurred())
		Expect(fakeActor.CreateOrgCallCount()).To(Equal(1))
		Expect(fakeActor.CreateOrgArgsForCall(0)).To(Equal("some-org"))
	})
})
