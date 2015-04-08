package learning_about_container_schedulers_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestLearningAboutContainerSchedulers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "LearningAboutContainerSchedulers Suite")
}
