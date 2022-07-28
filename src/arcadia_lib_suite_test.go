package arcadia_lib_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestArcadiaLib(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ArcadiaLib Suite")
}
