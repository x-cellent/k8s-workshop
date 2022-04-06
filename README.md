# k8s workshop

## Prerequisites

Only [Make](https://www.gnu.org/software/make/) and [Docker](https://docs.docker.com/get-docker/) is required to be installed.

## Build

Run `make` or `make build` to build the workshop CLI, which will then be compiled to `bin/w6p`.

## Slides

The workshop slides are available
- ...[online](https://x-cellent.github.io/k8s-workshop)
- ...at `http://localhost:8080/docs` after running `bin/w6p slides`
- ...at `http://localhost:8080` after running `make slides`

The last option can be used for development.
Simply press `F5` in your browser after any update to [slides.md](./docs/slides.md).

## Hands-On

### Start/Stop workshop cluster

```
bin/w6p cluster run
bin/w6p cluster shutdown
```

#### Interact with workshop cluster

The easiest way to interact with the workshop cluster can be accomplished by
```
export KUBECONFIG=$(pwd)/k8s-workshop.kubeconfig
k9s
```

### Exercises

After `make cluster run` run
```
bin/w6p exercise -n X
```
to start exercise number `X`.
Each exercise will be deployed to a fresh namespace `exX` including the exercise number `X`, e.g. `ex3`.
