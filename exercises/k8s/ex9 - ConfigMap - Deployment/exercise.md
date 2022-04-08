Damit wir Konfigurationsdateien in Kubernetes einbinden können,
gibt es ConfigMaps.

Im Cluster ex9 liegt eine ConfigMap namens `nginx-configmap`.

Schaue diese an, exportiere sie in eine YAML Datei und update die `worker-connection`.
Anschließend baue sie in das Deployment `web` als ReadOnly VolumeMount ein.
