#!/usr/bin/env bash

cat <<EOF > pod.yaml
apiVersion: v1
kind: Pod
metadata:
  labels:
    app: frontend
  name: web
  namespace: ex1 # Optional, kann auch kubectl auch via "-n ex1" mitgegeben werden
spec:
  containers:
  - name: web
    image: nginx:latest
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
