1. Installiere den letsencrypt certmanager via Helm.

2. Installiere den ingress-nginx controller via Helm.

3. Installiere den letsencrypt-staging cluster issuer via kubectl.
   Finde dafür den Namen der Ingress Klasse heraus.

4. Da wir uns in einem unmanaged Kubernetes-Cluster befinden, gibt es keine CRD Admission Controller, die Services
   vom Typ `LoadBalancer` abfangen und uns automatisch einen solchen aufbaut.
   In dieser Aufgabe geht es jetzt nicht darum, einen geeigneten LoadBalancer manuell aufzusetzen,
   sondern den Service vom Typ `LoadBalancer` aus Aufgabe 2 in einen Service vom Typ `NodePort` umzuwandeln.

5. Erstelle ein default nginx Deployment mit Service und Ingress.
   Verwende für den Ingress die Ingress-Klasse aus Aufgabe 3 und als Host `nginx.nip.io`.

6. Teste, ob du die Landing Page vom nginx via `curl -k -H"Host: nginx.nip.io" https://<cluster-container-ip>:<https-node-port>`.
   Tipps:
   Die <cluster-container-ip> kann mittels `docker inspect --format='{{.NetworkSettings.Networks.kind.IPAddress}}' k8s-workshop-cluster-control-plane` ermittelt werden.
   Der <https-node-port> kann mittels `kubectl -n ingress-nginx get svc ingress-nginx-controller -o custom-columns=NODEPORT:.spec.ports[1].nodePort` ermittelt werden.
