10m
```yaml
---
apiVersion: v1
kind: Service
metadata:
  name: web
spec:
  selector:
    app: frontend
  ports:
    - protocol: TCP
      port: 80
      targetPort: 80
```
