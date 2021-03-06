15m
```yaml
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
```

```sh
k apply -f cronjob.yaml
k create job -n ex8 --from=cronjob/hello hello-test
```

Test via:

```sh
k logs -n ex8 -l cronjob=hello
```
