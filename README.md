# k8s workshop

## Prerequisites

Only [Git](https://www.atlassian.com/git/tutorials/install-git), [Make](https://www.gnu.org/software/make/) and [Docker](https://docs.docker.com/get-docker/) are required to be installed.

## Build

Run `make` or `make build` to build the workshop CLI, which will then be compiled to `bin/w6p`.

## Slides

The workshop slides are available
- ...[online](https://x-cellent.github.io/k8s-workshop)
- ...at `http://localhost:8080/docs` after running `bin/w6p slides show`
- ...at `http://localhost:8080` after running `make slides`

The last option can be used for development.
Simply press `F5` in your browser after any update to [slides.md](./docs/slides.md).

### Export to PDF

```
bin/w6p slides export
```

## Hands-On

### Docker exercises
Run
```
bin/w6p exercise docker -nX
´``
to start the Docker exercise number `X`.

### Kubernetes exercises

First you need to start the local Kubernetes workshop cluster:
```
bin/w6p cluster run
```

Note:
This cluster runs within the Docker container `k8s-workshop-cluster-control-plane`.
Cou can shutdown the workshop cluster by either of the following commands:
```
bin/w6p cluster shutdown
docker rm -f k8s-workshop-cluster-control-plane
```

#### Interact with workshop cluster

The easiest way to interact with the workshop cluster can be accomplished by
```
export KUBECONFIG=$(pwd)/k8s-workshop.kubeconfig
k9s
```

### Kubernetes exercises

First start the workshop cluster through `make cluster run`.
Then run
```
bin/w6p exercise k8s -nX
```
to start the Kubernetes exercise number `X`.
Each exercise will be deployed into a fresh namespace `exX` including the exercise number `X`, for example `ex1`.
