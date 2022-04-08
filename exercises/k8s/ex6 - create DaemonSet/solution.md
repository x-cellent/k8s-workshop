15m
In der Kubernetes Dokumentation von [DaemonSets](https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/)
ist ein Beispiel für ein fluentd enthalten.

In dieser YAML den Namespace anpassen und dann mit kubectl deployen.

Ergebnis ist nur Pod, da wir nur einen Node haben.

Sieht man mit

```sh
k get nodes
```
