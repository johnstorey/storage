package storage_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// PREAMBLE.

func TestCart(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Storage Test Suite")
}

// Tests.
var _ = Describe("Storage tests", func() {

	Context("Initially", func() {
		Specify("has 0 Environments", func() {

		})
	})

	Context("when a new Environment is added", func() {
		Specify("there are no Android nodes", func() {

		})

		Specify("there are no iOS nodes", func() {

		})

		Specify("getting the Android nodes returns a string 'UNKNOWN'", func() {

		})

		Specify("getting the iOS nodes returns a string 'UNKNOWN'", func() {

		})

		Specify("converting the Environment to string creates a table format string", func() {

		})

	})

	Context("when a node is added for a nonexistent environment", func() {
		Specify("an error is returned", func() {

		})

	})

	Context("when as Android node is added for an existing environment with no nodes", func() {
		Specify("that Environment has the added node", func() {

		})

		Specify("the Node shows up when the Environments is converted to a table", func() {

		})
	})

	Context("when an Android node is added for an existing Environment", func() {
		Specify("that Environment has the added node and the original node", func() {

		})

		Specify("both nodes show up when the Environment is converted to a table", func() {

		})

	})

	Context("when an Environment had an Android and and 2 iOS nodes", func() {
		Specify("the Environment table is laid our correctly when converted to a table", func() {

		})

	})

	Context("when a Node is removed for a nonexistent Environment", func() {
		It("an error is returned", func() {

		})
	})

	Context("when a Environment has 2 nodes and the 1st is removed", func() {
		Specify("the first is gone", func() {

		})

		Specify("the second is there", func() {

		})

		Specify("the table correct when the Environment is formatted as a table", func() {

		})

	})

	Context("when a Environment has 2 nodes and the 2nd is removed", func() {
		Specify("the first is there", func() {

		})

		Specify("the second is gone", func() {

		})

		Specify("the table correct when the Environment is formatted as a table", func() {

		})

	})

	Context("when a Environment has 2 nodes and both are removed", func() {
		Specify("both nodes are gone", func() {

		})

		Specify("attempting to retrieve a node returns a string 'UNKNOWN'", func() {

		})

		Specify("the table correct when the Environment is formatted as a table", func() {

		})

	})

})
