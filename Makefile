.PHONY .SILENT: default
default: build

.PHONY .SILENT: build
build:
	docker run --rm -v $(shell pwd):/src golang:1.18 sh -c 'cd /src && go fmt ./... && go mod download && go mod tidy && go build -o bin/k8s-workshop'

.PHONY .SILENT: slides
slides: down
	docker run -d --name k8s-workshop-slides --net host -v $(shell pwd):/k8s-workshop golang:1.18 sh -c 'cd /k8s-workshop/cmd/slides && go run main.go' >/dev/null
	echo "Slides are ready at localhost:8080! Stop slides server with 'make down' or 'docker rm -f k8s-workshop-slides'"

.PHONY .SILENT: down
down:
	docker rm -f k8s-workshop-slides >/dev/null 2>&1 || true
