```Dockerfile
FROM alpine:3.9 
RUN apk add --no-cache mysql-client 
ENTRYPOINT ["mysql"] 
```

```sh
docker build -t mysql:5.6 .
```