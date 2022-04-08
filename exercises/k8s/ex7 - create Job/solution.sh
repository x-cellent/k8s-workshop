#!/usr/bin/env bash

cat <<EOF > job.yaml
apiVersion: batch/v1
kind: Job
metadata:
  name: pi
  namespace: ex5
spec:
  template:
    metadata:
      labels:
        job: pi
    spec:
      containers:
      - name: pi
        image: perl
        command: ["perl",  "-Mbignum=bpi", "-wle", "print bpi(5000)"]
      restartPolicy: Never
  backoffLimit: 4
EOF
