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

## Development

The project uses [Godep](https://github.com/tools/godep) to manage
dependencies. To download Godep do the following:
`go install github.com/tools/godep`.

You can then load all of the dependencies: `godep restore`. And then
run all of the tests: `godep go test -v ./...`.
