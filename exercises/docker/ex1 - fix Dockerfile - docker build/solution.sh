#!/usr/bin/env bash

cat <<EOF > Dockerfile
FROM ubuntu:20.04
RUN apt update \\
 && apt install -y netcat
ENTRYPOINT ["nc"]
EOF

echo "docker build -t nc ." > build.sh
chmod +x build.sh
