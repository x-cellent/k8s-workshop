#!/usr/bin/env bash

cat <<EOF > app.go
package main

import "fmt"

func main() {
    fmt.Println("Hello Go!")
}
EOF

cat <<EOF > Dockerfile
FROM golang:1.17 AS builder

WORKDIR /work

COPY app.go .

ENV GO111MODULE=off

RUN go build -o bin/my-app

FROM scratch

COPY --from=builder /work/bin/my-app /

ENTRYPOINT ["/my-app"]
EOF

cat <<EOF > build-and-run.sh
#!/usr/bin/env bash

docker build -t hello .
docker run --rm -t hello
EOF

chmod +x build-and-run.sh