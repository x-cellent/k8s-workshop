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

3. Manifest des letsencrypt-staging cluster issuers:
```yaml
apiVersion: cert-manager.io/v1
kind: ClusterIssuer
metadata:
  name: letsencrypt-staging
spec:
  acme:
    server: https://acme-v02.api.letsencrypt.org/directory
    email: sandro.koll@x-cellent.com
    privateKeySecretRef:
      name: letsencrypt-staging
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
```sh
curl -k -H"Host: nginx.nip.io" https://$(docker inspect --format='{{.NetworkSettings.Networks.kind.IPAddress}}' k8s-workshop-cluster-control-plane):$(kubectl -n ingress-nginx get svc ingress-nginx-controller -o custom-columns=NODEPORT:.spec.ports[1].nodePort --no-headers)
```

Zum Beispiel:
```sh
curl -k -H"Host: nginx.nip.io" https://172.22.0.2:32318
```