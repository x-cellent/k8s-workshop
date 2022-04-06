15m

in der Kubernetes Dokumentation von daemonsets https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/ ist ein Beispiel für ein fluentd daemonset drinnen. 

Dies als yaml deployen

nur ein pod, da wir nur eine Node haben.

sieht man mit 

```sh
kubectl get nodes
```