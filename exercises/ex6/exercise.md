Aufgabe 6:

damit wir Konfigurationsdateien in Kubernetes einbinden können, gibt es configmaps, im Cluster ex6 liegt eine configmap namens `nginx-configmap`

schaue diese an, exportiere diese in eine yaml und update die `worker-connection`

Anschließend baue sie in das Deployment `web` als ReadOnly VolumeMount ein.