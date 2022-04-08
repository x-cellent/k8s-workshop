.PHONY .SILENT: default
default: build

.PHONY .SILENT: build
build:
	git pull >/dev/null 2>&1 || true
	docker run --rm -v $(shell pwd):/src golang:1.17 sh -c 'cd /src && GO111MODULE=off go fmt ./... && go build -o bin/w6p'
	mkdir -p $(HOME)/bin
	cp bin/w6p $(HOME)/bin/

.PHONY .SILENT: up
up: down
	docker run -d --name k8s-workshop-slides --net host -v $(shell pwd):/k8s-workshop golang:1.17 sh -c 'cd /k8s-workshop && go run static.go' >/dev/null

.PHONY .SILENT: down
down:
	docker rm -f k8s-workshop-slides >/dev/null 2>&1 || true

.PHONY .SILENT: slides
slides: up
	echo "Will browse slides at http://localhost:8080 in 3 seconds..."
	sleep 3
	xdg-open http://localhost:8080
	echo "Stop slides server with 'make down' or 'docker rm -f k8s-workshop-slides'"

.PHONY .SILENT: pdf
pdf: up
	sleep 3
	docker run --rm -t --net host -v $(shell pwd):/docs astefanutti/decktape http://localhost:8080 /docs/slides.pdf
	make --no-print-directory down
