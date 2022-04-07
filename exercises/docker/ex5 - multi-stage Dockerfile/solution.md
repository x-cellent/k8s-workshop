15m
1.

```Dockerfile
FROM golang:1.17 AS builder

WORKDIR /work

COPY app.go .

RUN go build -o bin/my-app

FROM scratch

COPY --from=builder /work/bin/my-app /

ENTRYPOINT ["/my-app"]
```

2.

```sh
docker build -t hello .
```

3.

```sh
docker run --rm -t hello
# Hello Go!
```
