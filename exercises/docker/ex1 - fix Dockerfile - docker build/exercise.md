1. Korrigiere folgendes Dockerfile:

```Dockerfile
FROM ubuntu:latest:20.04
RUN apt update
 && apt install netcat
ENTRYPOINT ["nc"]
CMD ["nc"]
ENTRYPOINT ["netcat"]
```

2. Baue dann daraus ein Image und benenne es `nc`.
