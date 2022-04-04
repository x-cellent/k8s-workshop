Aufgabe 2:

in dieser Mysql Dockerfile ist ein Fehler, dieser muss gefixt werden.

anschließend muss das Dockerfile gebuildet werden

```Dockerfile
FROM alpine:latest:3.9 
RUN apk add --no-cache mysql-client 
ENTRYPOINT ["/root/"]
ENTRYPOINT ["mysql"] 
```