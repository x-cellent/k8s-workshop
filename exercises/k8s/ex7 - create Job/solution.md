5m
Von der Kubernetes Dokumentation bzgl. Jobs das Manifest ├╝bernehmen und anpassen:

```yaml
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
```

Test via:

```sh
k -n ex7 logs -l job=pi
```
