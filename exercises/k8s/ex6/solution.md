15m
```yaml
apiVersion: batch/v1
kind: CronJob
metadata:
  name: hello
  namespace: ex6
spec:
  successfulJobsHistoryLimit: 5
  failedJobsHistoryLimit: 8
  schedule: "*/10 * * * *"
  jobTemplate:
    spec:
      template:
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
```

```sh
kubectl create job -n ex6 --from=cronjob/hello hello-test
```