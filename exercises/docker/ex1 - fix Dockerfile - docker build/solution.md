5m
1.

```Dockerfile
FROM ubuntu:20.04
RUN apt update \
 && apt install -y netcat
ENTRYPOINT ["nc"]
```

2.

```sh
docker build -t nc .
```
