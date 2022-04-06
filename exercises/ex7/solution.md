5m

Von der Kubernetes dokumentation zu Jobs das Manifest übernehmen und anpassen

```yaml
apiVersion: batch/v1
kind: Job
metadata:
  name: pi
  namespace: ex5
spec:
  template:
    spec:
      containers:
      - name: pi
        image: perl
        command: ["perl",  "-Mbignum=bpi", "-wle", "print bpi(5000)"]
      restartPolicy: Never
  backoffLimit: 4
```

Anschließend mit kubectl logs die logs ausgeben lassen