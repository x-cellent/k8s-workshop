#!/usr/bin/env bash

cat <<EOF > Dockerfile
# inherit from python:3.6
FROM python:3.6

# add a "author" label
LABEL authors="Jon Doe <jon.doe@acme.com>, Max Mustermann <max.mustermann@acme.com>"

# install the following packages with pip3: flask, redis, requests, uwsgi
RUN pip3 install --no-cache-dir \
    flask==1.0.2 \
    redis==3.2.1 \
    requests==2.21.0 \
    uwsgi==2.0.18

WORKDIR /app

# copy the startscript (cmd.sh) into the image
COPY cmd.sh /

# copy the app folder (./app) into the image
COPY app/* /app/

# Expose the ports 9090 9091
EXPOSE 9090 9091

# Start the application script
CMD ["/cmd.sh"]
EOF
