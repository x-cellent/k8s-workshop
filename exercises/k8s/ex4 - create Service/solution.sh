#!/usr/bin/env bash

cat <<EOF > svc.yaml
apiVersion: v1
kind: Service
metadata:
  name: web
  namespace: ex4
spec:
  selector:
    app: frontend
  ports:
    - protocol: TCP
      port: 80
      targetPort: 80
EOF
