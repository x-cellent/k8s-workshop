20m
1.
```sh
helm repo add jetstack https://charts.jetstack.io
helm repo update
helm upgrade --install cert-manager jetstack/cert-manager \
  --namespace cert-manager --create-namespace \
  --version v1.5.3 \
  --set installCRDs=true \
  --timeout 5m
```

2.
```sh
helm upgrade --install ingress-nginx ingress-nginx \
  --namespace ingress-nginx --create-namespace \
  --repo https://kubernetes.github.io/ingress-nginx \
  --timeout 5m
```

3. Manifest des letsencrypt-prod cluster issuers:
```yaml
apiVersion: cert-manager.io/v1
kind: ClusterIssuer
metadata:
  name: letsencrypt-prod
spec:
  acme:
    server: https://acme-v02.api.letsencrypt.org/directory
    email: sandro.koll@x-cellent.com
    privateKeySecretRef:
      name: letsencrypt-prod
    solvers:
      - http01:
          ingress:
            class: nginx
```

```sh
kubectl apply -f cluster-issuer.yaml
```

4.
```sh
kubectl -n ingress-nginx patch svc ingress-nginx-controller -p '{"spec": {"type": "NodePort"}}'
```

5.
```sh
kubectl apply -f dep.yaml
kubectl apply -f svc.yaml
kubectl apply -f ingress.yaml
```

6.
Zum Beispiel:
```sh
curl -k -H"Host: 10.10.11.138.nip.io" https://172.22.0.2:32318
```