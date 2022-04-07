#!/usr/bin/env bash

cat <<EOF > Dockerfile
FROM golang:1.17

WORKDIR /work

ENV GO111MODULE=off

RUN go build -o bin/my-app

FROM scratch

ENTRYPOINT ["/my-app"]
EOF
