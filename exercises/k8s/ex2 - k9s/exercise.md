1. Verbinde dich über k9s mit dem Cluster.
   Switche in den Namespace `my-web`.

2. Skaliere das Deployment `werbserver` aus Aufgabe 1 hoch auf 2 Replicas.

3. Lasse dir alle ReplicaSets im Namespace `my-web` anzeigen.
   Skaliere dieses runter auf 1.
   Warum hat das keinen Effekt?

4. Wechlse zur Podansicht. Lösche einen der beiden webserver Pods.
   Warum wird sofort wieder ein neuer erstellt?

5. Wechsle wieder zur Deploymentansicht.
   Ändere den Image-Tag des nginx Images zu `alpine`.
   Was erwartest du?

6. Beweise, dass die zugrunde liegenden Pods neu gestartet wurden und nun
   auf dem neuen Image basieren.
