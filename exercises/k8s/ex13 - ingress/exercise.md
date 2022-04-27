1. Installiere den letsencrypt certmanager via Helm.

2. Installiere den ingress-nginx controller via Helm.

3. Installiere den letsencrypt-prod cluster issuer via kubectl.
   Finde dafür den Namen der Ingress Klasse heraus und verwende deine Firmen-Email Adresse.

4. Da wir uns in einem unmanaged Kubernetes-Cluster befinden, gibt es keine CRD Admission Controller, die Services
   vom Typ `LoadBalancer` abfangen und uns automatisch einen solchen aufbaut.
   In dieser Aufgabe geht es jetzt nicht darum, einen geeigneten LoadBalancer manuell aufzusetzen,
   sondern den Service vom Typ `LoadBalancer` aus Aufgabe 3 in einen Service vom Typ `NodePort` umzuwandeln.

5. Erstelle ein nginx Deployment mit Service und Ingress.
   Verwende für den Ingress die Ingress-Klasse aus Aufgabe 2 und als Host die NodePort IP des Services aus Aufgabe 4, gefolgt vom Suffix `.nip.io`.

6. Teste, ob du die Landing Page vom nginx via `curl -k -H"Host: <node-port-service-ip>.nip.io" https://<cluster-container-ip>:<https-node-port>`.
