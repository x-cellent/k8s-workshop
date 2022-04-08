Folgende YAML Datei definiert einen Pod, ist aber fehlerhaft.
Korrigiere alle Fehler und deploye den Pod dann in den Namespace `ex1`.

```yaml
apiversion: v1
Kind: pod
metadata:
  labels:
    app: frontend
  name: web
spec:
containers:
  name: web
    image: nginx
    tag: latest
    ports:
  - containerPort: 80
    resources:
      requests:
        cpu: "1.0"
        memory:"1G"
      limits:
       cpu: "1.0"
        memory: 1G
```

Hinweis:
Suche nach `Pod` auf "https://kubernetes.io/docs" und achte auf die YAML Syntax.
