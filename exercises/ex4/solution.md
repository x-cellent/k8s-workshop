10m
```yaml
---
apiVersion: v1
kind: Service
metadata:
  name: web
  namespace: ex2
spec:
  selector:
    app: frontend
  ports:
    - protocol: TCP
      port: 80
      targetPort: 80
```
