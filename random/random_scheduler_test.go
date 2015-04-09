package random_test

import (
	. "github.com/kushalp/learning-about-container-schedulers/random"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/GoogleCloudPlatform/kubernetes/pkg/api"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/scheduler"
)

var _ = Describe("RandomScheduler", func() {
	var sched scheduler.Scheduler

	BeforeEach(func() {
		sched = NewRandomScheduler(scheduler.FakePodLister([]api.Pod{}))
	})

	It("returns an error if no minions present", func() {
		minionLister := scheduler.FakeMinionLister(makeNodeList([]string{}))
		pod := api.Pod{ObjectMeta: api.ObjectMeta{Name: "never touched"}}

		_, err := sched.Schedule(pod, minionLister)

		Expect(err).ToNot(BeNil())
	})

	It("returns the only minion in a minion set of one", func() {
		minionLister := scheduler.FakeMinionLister(
			makeNodeList([]string{"1"}))
		pod := api.Pod{ObjectMeta: api.ObjectMeta{Name: "1"}}

		minion, err := sched.Schedule(pod, minionLister)

		Expect(err).To(BeNil())
		Expect(minion).To(Equal("1"))
	})

	It("eventually selects all minions", func() {
		nodes := []string{"3", "2", "1"}
		minionLister := scheduler.FakeMinionLister(
			makeNodeList(nodes))
		pod := api.Pod{ObjectMeta: api.ObjectMeta{Name: "3"}}

		for _, node := range nodes {
			Eventually(func() (string, error) {
				return sched.Schedule(pod, minionLister)
			}).Should(Equal(node))
		}
	})
})

func makeNodeList(nodeNames []string) api.NodeList {
	result := api.NodeList{
		Items: make([]api.Node, len(nodeNames)),
	}

	for ix := range nodeNames {
		result.Items[ix].Name = nodeNames[ix]
	}

	return result
}
