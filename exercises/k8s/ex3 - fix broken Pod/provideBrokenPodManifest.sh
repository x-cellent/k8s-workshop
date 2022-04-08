#!/usr/bin/env bash

cat <<EOF > broken-pod.yaml
apiVersion: v1
Kind: pod
metadata:
  labels:
    app: frontend
  name: web
spec:
containers:
- name: web
    image: nginx
    tag: latest
    ports:
  - containerPort: 80
    resources:
      requests:
       cpu: "1.0"
        memory: "1G"
      limits:
       cpu: "1.0"
        memory: "1G"
EOF
