20m
1. Verbinde dich über k9s mit dem Cluster.
   Switche in den Namespace `my-web`.

```sh
k9s
:ns
```

Selektiere 'my-web' und drücke ENTER.

2. Skaliere das Deployment `werbserver` aus Aufgabe 1 hoch auf 2 Replicas.

```sh
:dp
```

Selektiere das webserver Deployment und drücke 's'.
Trage nun unter Replicas 2 ein und drücke OK.

3. Lasse dir alle ReplicaSets im Namespace `my-web` anzeigen.
   Es gibt nur eins. Skaliere dieses runter auf 1.
   Warum hat das keinen Effekt?

```sh
:rs
```

Selektiere das webserver ReplicaSets und drücke 's'.
Trage nun unter Replicas 1 ein und drücke OK.

Die Änderung wird sofort wieder vom übergeordneten Deployment rückgängig gemacht,
denn dort sind nach wie vor 2 Replicas definiert.

4. Wechlse zur Podansicht. Lösche einen der beiden webserver Pods.
   Warum wird sofort wieder ein neuer erstellt?

```sh
:po
```

Selektiere einen der beiden webserver Pods und drücke 'Ctrl^d'.
Bestätige das Löschen mit OK.

Du kannst sehen, wie der Pod terminiert wird und ein neuer Pod gestartet wird.
Das hat das übergeordnete ReplicaSet veranlasst.

5. Wechsle wieder zur Deploymentansicht.
   Ändere den Image-Tag des nginx Images zu `alpine`.
   Was erwartest du?

```sh
:dp
```

Selektiere das webserver Deployment und drücke 'e'.
Es öffnet sich dein Default Editor.
Suche nach dem nginx Image und schreibe ':alpine' dahinter.
Speichere deine Änderung und verlasse den Editor.

6. Beweise, dass die zugrunde liegenden Pods neu gestartet wurden und nun
   auf dem neuen Image basieren.

Wechlse in die Podansicht.

```sh
:po
```

Selektiere einen der Pods und drücke 'y'.
Suche in der YAML Ausgabe nach dem zugrunde liegenden Image.
