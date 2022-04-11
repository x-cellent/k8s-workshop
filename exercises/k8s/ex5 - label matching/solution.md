15m
1. Erstelle das Deployment im Aufgaben-Unterverzeichnis.

```sh
k apply -f deploy.yaml
```

2. Nehme den laufenden Pod aus der Verantwortung des Deployments,
   um ihn für Analysezwecke nutzen zu können.
   Es gibt drei Wege. Welche?

Ändere oder lösche das Label `app: frontend` des laufenden Pods.

3. Lösche nun diesen Pod.

Via k9s mit `Ctrl^d`

4. Stelle einen Update-Konflikt nach:
   Das Scaling des Deployments soll erhöht werden.
   Alice und Bob gehen das Problem zeitgleich an, ohne voneinander
   zu wissen. Alice empfindet 3 Replicas als ausreichend und updated das
   Deployment vor Bob, der sich sogar mit 2 zufrieden gibt.
   Was ist das Ergebnis und warum?

Starte k9s in einem weiteren Terminal.
Öffne nun in beiden Terminals mit `e` das Deployment.
Verifiziere, dass unter `.metadata.annotations.kubectl.kubernetes.io/last-applied-configuration`
die Replicas mit 1 angegeben sind.
Ändere in Terminal 1 (Alice) den Replica-Count auf 3 und speichere.
Beobachte, wie zwei weitere Pods entstehen.
Ändere in Terminal 2 (Bob) den Replica-Count auf 2 und speichere.
Beobachte, wie ein Pod wieder gelöscht wird.

Es entsteht kein Konflikt! Kubernetes setzt immer den Request um, der gerade reinkommt,
ohne auf etwaige Konflikte zu prüfen, selbst wenn ein solcher anhand der Metadaten festgestellt
werden könnte. Für solche Zwecke ist ein vorgeschaltetes Version-Control System wie Git unabdingbar.
