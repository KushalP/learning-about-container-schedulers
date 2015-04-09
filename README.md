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
dependencies.

To download Godep run the following:

`go install github.com/tools/godep`.

### Fetching dependencies

To check out listed dependency versions in your `$GOPATH`:

```bash
godep restore
```

### Running the tests

To run the tests you'll need to run the go tool with saved
dependencies under Godep:

```bash
godep go test -v ./...
```
