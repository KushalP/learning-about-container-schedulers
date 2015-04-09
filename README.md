# Learning about container schedulers

This is a repository of artefacts I'm using to learn more about how
container schedulers work. The common platform used is the
[Kubernetes](https://github.com/GoogleCloudPlatform/kubernetes)
project and the
[scheduler interface](https://github.com/GoogleCloudPlatform/kubernetes/tree/master/pkg/scheduler)
it exposes.

## Schedulers

Currently there's only a single scheduler that's been implemented: a
random scheduler.
