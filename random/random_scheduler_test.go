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
		pod := api.Pod{ObjectMeta: api.ObjectMeta{Name: "3"}}

		_, err := sched.Schedule(pod, minionLister)

		Expect(err).ToNot(BeNil())
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
