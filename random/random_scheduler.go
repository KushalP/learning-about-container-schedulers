package random

import (
	"fmt"

	"github.com/GoogleCloudPlatform/kubernetes/pkg/api"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/scheduler"
)

type randomScheduler struct {
	pods scheduler.PodLister
}

func (r *randomScheduler) Schedule(pod api.Pod, minionLister scheduler.MinionLister) (string, error) {
	minions, err := minionLister.List()
	if err != nil {
		return "", err
	}
	if len(minions.Items) == 0 {
		return "", fmt.Errorf("no minions available to schedule pods")
	}

	return "", nil
}

func NewRandomScheduler(pods scheduler.PodLister) scheduler.Scheduler {
	return &randomScheduler{pods: pods}
}
