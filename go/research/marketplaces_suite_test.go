// +build spec

package research

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestMarketplaces(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Marketplaces Suite")
}
