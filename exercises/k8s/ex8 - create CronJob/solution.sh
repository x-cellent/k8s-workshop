#!/usr/bin/env bash

cat <<EOF > cronjob.yaml
apiVersion: batch/v1
kind: CronJob
metadata:
  name: hello
  namespace: ex8
spec:
  successfulJobsHistoryLimit: 5
  failedJobsHistoryLimit: 8
  schedule: "*/1 * * * *"
  jobTemplate:
    spec:
      template:
        metadata:
          labels:
            cronjob: hello
        spec:
          containers:
          - name: hello
            image: busybox:1.28
            imagePullPolicy: IfNotPresent
            command:
            - /bin/sh
            - -c
            - date; echo Pascal
          restartPolicy: OnFailure
EOF
