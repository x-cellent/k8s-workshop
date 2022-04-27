#!/usr/bin/env bash

set -e

kubectl create ns ex13 || true

helm repo add jetstack https://charts.jetstack.io
helm repo update
helm upgrade --install cert-manager jetstack/cert-manager \
  --namespace cert-manager --create-namespace \
  --version v1.5.3 \
  --set installCRDs=true \
  --timeout 5m
