10m
```yaml
apiVersion: v1 #Typo in apiVersion, V von Version muss groß sein
kind: Pod #Typo, kind muss klein sein
metadata:
  labels:
    app: frontend
  name: web
  namespace: ex1 # Optional, kann auch kubectl auch via "-n ex1" mitgegeben werden
spec:
  containers: #ab hier muss alles eingeruckt sein
  - name: web #listen in yaml werden beim ersten punkt mit `-`angegeben
    image: nginx:latest #Image und Tag definiert man in einer zeile mit `:` dazwischen
    ports:
    - containerPort: 80 #hier auch falsch eingerückt
    resources:
      requests:
        cpu: "1.0"
        memory: "1G" # zwischen memory und den 1G muss ein Leerzeichen sein
      limits:
        cpu: "1.0"
        memory: "1G" #1G muss in anführungszeichen sein
```
