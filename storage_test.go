package storage

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

	var ds DataStore

	BeforeEach(func() {
		ds = DataStore{}
	})

	AfterEach(func() {
		ds.emptyCollection("environments")
		ds.emptyCollection("nodes")

		ds.disconnect()
	})

	Context("Initially", func() {

		Specify("has 0 Environments", func() {
			numEnvironments, e := ds.numEnvironments()
			if e != nil {
				panic("Did not get expected number of environments.")
			}
			Expect(numEnvironments).To(Equal(0), "Expected 0 environments got %d", numEnvironments)
		})
	})

	Context("when a new Environment is added", func() {
		Specify("there are no nodes", func() {
			newEnvironment := newEnvironment()
			newEnvironment.setName("test-1")
			newEnvironment.save(&ds)
			nodes := newEnvironment.nodes()
			Expect(len(nodes)).Should(Equal(0))
		})

		Specify("there are no Android nodes", func() {
			n, err := ds.findEnvironmentByName("test-1")
			Expect(err).NotTo(HaveOccurred())
			nodes, err := n.androidNodes(&ds)
			Expect(err).NotTo(HaveOccurred())
			Expect(len(nodes)).Should(Equal(0))
		})

		Specify("there are no iOS nodes", func() {
			newEnvironment := newEnvironment()
			newEnvironment.setName("test")
			nodes, err := newEnvironment.iOSNodes(&ds)
			Expect(err).NotTo(HaveOccurred())
			Expect(len(nodes)).Should(Equal(0))
		})

		Specify("can save, retrieve, alter, save, and retrieve and environemt", func() {
			createdEnv := newEnvironment()
			createdEnv.setName("Stephen")
			createdEnv.save(&ds)

			// Retrieve environment from datastore.
			foundEnv, err := findEnvironmentByName(&ds, createdEnv.Name)
			//foundEnv, err := findEnvironmentById(&ds, createdEnv.ID)
			Expect(err).NotTo(HaveOccurred())
			Expect(foundEnv.Name).To(Equal(createdEnv.Name))
			Expect(len(foundEnv.MesmerNodes)).To(Equal(0))

			// Update the host name.
			foundEnv.Name = "Nisha"
			foundEnv.save(&ds)
			foundEnv2, err := findEnvironmentByName(&ds, foundEnv.Name)
			Expect(foundEnv.Name).To(Equal(foundEnv2.Name))
		})

		Specify("can add and remove nodes", func() {
			newEnvironment := newEnvironment()
			newEnvironment.setName("test environment Bobbo")
			iOSNode := MesmerNode{Host: "OSX0023", NodeType: IOS, IP: "1.1.1.1"}
			androidNode := MesmerNode{Host: "UBR0001", NodeType: ANDROID, IP: "1.1.1.21"}

			newEnvironment.addNode(&ds, iOSNode)
			newEnvironment.addNode(&ds, androidNode)

			// Convert to a string.
			envString := newEnvironment.toStrings()
			Expect(envString[0]).To(Equal("ios OSX0023 1.1.1.1"))
			Expect(envString[1]).To(Equal("android UBR0001 1.1.1.21"))

			// Removing the first node, saving, and retrieving removes it from the db.
			newEnvironment.save(&ds)
		})

	})

	Context("when a Environment has 2 nodes and the 1st is removed", func() {
		Specify("you can remove them one by one and the database reflects it", func() {
			// Set up test environment.
			newEnvironment := newEnvironment()
			newEnvironment.setName("test environment Cake")
			iOSNode := MesmerNode{Host: "OSX0023", NodeType: IOS, IP: "1.1.1.1"}
			androidNode := MesmerNode{Host: "UBR0001", NodeType: ANDROID, IP: "1.1.1.21"}

			newEnvironment.addNode(&ds, iOSNode)
			newEnvironment.addNode(&ds, androidNode)

			newEnvironment.save(&ds)

			// Retrieve envrionment from datastore and ensure nodes are there.
			e1, err := findEnvironmentByID(&ds, newEnvironment.ID)
			if err != nil {
				panic(err)
			}
			Expect(len(e1.nodes())).To(Equal(2))

			// Remove a node, update the database, and verify it is gone but the other is there.
			e1.removeNodeByHost("OSX0023")
			e1.update(&ds)
			e2, _ := findEnvironmentByID(&ds, e1.ID)
			Expect(len(e2.nodes())).To(Equal(1))
			var e2Node MesmerNode
			e2Node = e2.nodes()[0]
			Expect(e2Node.Host).To(Equal("UBR0001"))
			Expect(e2Node.NodeType).To(Equal(ANDROID))
			Expect(e2Node.IP).To(Equal("1.1.1.21"))

			// Remove the last node and verify it is gone.
			e2.removeNodeByHost("UBR0001")
			e2.update(&ds)
			e3, _ := findEnvironmentByID(&ds, e1.ID)
			Expect(len(e3.nodes())).To(Equal(0))
		})

	})
})
