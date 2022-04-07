#!/usr/bin/env bash

cat <<EOF > Dockerfile
FROM ubuntu:latest:20.04
RUN apt update
 && apt install netcat
ENTRYPOINT ["nc"]
CMD ["nc"]
ENTRYPOINT ["netcat"]
EOF
