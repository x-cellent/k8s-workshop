10m
Erst die ConfigMap anschauen:

```sh
kubectl get cm -n ex9 nginx-configmap
```

und als yaml file ausgeben:

```sh
kubectl get cm -n ex9 nginx-configmap -o yaml
```

und Deployment anpassen mit volumeMounts:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    app: frontend
  name: web
  namespace: default
spec:
  replicas: 3
  selector:
    matchLabels:
      app: frontend
  template:
    metadata:
      labels:
        app: frontend
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
        volumeMounts:
        - name: nginx-configmap
          mountPath: /etc/nginx 
          readOnly: true
      volumes:
        - name: nginx-configmap
          configMap:
            name: nginx-configmap
```
