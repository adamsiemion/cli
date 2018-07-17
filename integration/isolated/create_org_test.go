package isolated

import (
	"code.cloudfoundry.org/cli/integration/helpers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	. "github.com/onsi/gomega/gexec"
)

var _ = Describe("create-org", func() {
	Context("when the org does not exist yet", func() {
		var user string

		BeforeEach(func() {
			user = helpers.LoginCF()
		})

		It("creates the org", func() {
			orgName := helpers.NewOrgName()
			session := helpers.CF("create-org", orgName)
			Eventually(session).Should(Exit(0))

			Expect(session.Out).To(Say("Creating org %s as %s...", orgName, user))
			Expect(session.Out).To(Say("Assigning role OrgManager to user %s in org %s...", user, orgName))
			Expect(session.Out).To(Say(`TIP: Use 'cf target -o "%s"' to target new org`, orgName))
		})
	})
})
