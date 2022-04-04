Erst die configmap anschauen:

```sh
k get cm -n ex6 nginx-configmap
```

und als yaml file ausgeben:
```sh
k get cm -n ex6 nginx-configmap -o yaml
```

und deployment yaml anpassen mit volumemount:
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
        volumeMounts:
        - name: nginx-configmap
          mountPath: /etc/nginx 
          readOnly: true
      volumes:
        - name: nginx-configmap
          configMap:
            name: nginx-configmap
```