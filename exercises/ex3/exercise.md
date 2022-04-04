Diese yaml für einen Pod ist kaputt und muss gefixt werden, ob alles passt, sieht man wenn man den pod in den Namepsace ex3 deployen kann.

```yaml
apiVersion: v1
Kind: pod
metadata:
  labels:
    app: frontend
  name: web
  namespace: default
spec:
containers:
- name: web
    image: nginx
    tag: latest
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