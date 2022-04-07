15m
Pods werden von ReplicaSets gescaled.
ReplicaSets wiederum werden üblicherweise von Deployments gemanaged,
die außerdem noch Versionierung und zero-downtime Rollouts bieten:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    app: frontend
  name: web
  namespace: ex3
spec:
  replicas: 3
  selector:
    matchLables:
      app: frontend
  template:
    metadata:
      labels:
        app: web
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
```
