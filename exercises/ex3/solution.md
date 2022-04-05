5s
```yaml
apiVersion: v1
kind: Pod
metadata:
  labels:
    app: frontend
  name: web
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