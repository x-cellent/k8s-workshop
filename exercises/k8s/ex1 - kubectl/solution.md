20m
1. Liste alle Pods im Namespace `kube-system`:

```sh
k -n kube-system get pod
```

2. Erstelle lokal ein Pod Manifest `nginx-pod.yaml`, das vom Image `nginx` erbt
   und in den Namespace `my-web` deployed werden soll.
   Hinweis: Verwende dazu die Umgebungsvariable $do

```sh
k -n my-web run webserver --image nginx $do > nginx-pod.yaml
```

3. Deploye nun das `nginx-pod.yaml` Manifest.
   Hinweis: Löse das Problem, auf das du stoßen wirst.

```sh
k apply -f nginx-pod.yaml
# Error from server (NotFound): error when creating "nginx-pod.yaml": namespaces "my-deploy" not found
k create ns my-web
k apply -f nginx-pod.yaml
```

4. Zeige, dass der Pod läuft.

```sh
k -n my-web get pod
```

5. Lösche den Pod.

```sh
k -n my-web delete pod webserver
```

6. Erstelle nun ein Deployment, dass den obigen Pod deployed.
   Hinweis: Verwende dazu die Umgebungsvariable $do

```sh
k -n my-web create deploy --image nginx webserver $do > nginx-pod.yaml
```

oder