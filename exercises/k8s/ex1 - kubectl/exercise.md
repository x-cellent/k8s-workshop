Öffne https://x-cellent.github.io/k8s-workshop/

1. Verbinde dich mit kubectl gegen das Workshop cluster, indem du
   dir alle Pods im Namespace `kube-system` anzeigen lässt.

2. Erstelle lokal ein Pod Manifest `nginx-pod.yaml`, das vom Image `nginx` erbt
   und in den Namespace `my-web` deployed werden soll.
   Hinweis: Verwende dazu die Umgebungsvariable $do

3. Deploye nun das `nginx-pod.yaml` Manifest.
   Hinweis: Löse das Problem, auf das du stoßen wirst.

4. Zeige, dass der Pod läuft.

5. Lösche den Pod.

6. Erstelle nun ein Deployment, dass den obigen Pod deployed.
   Hinweis: Verwende dazu die Umgebungsvariable $do
