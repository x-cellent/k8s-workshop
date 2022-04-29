#!/usr/bin/env bash

set -e

export KUBECONFIG=$( dirname "$(readlink -f "${BASH_SOURCE[0]}" )")/../k8s-workshop.kubeconfig

kubectl create ns ex13 || true

helm upgrade --install ingress-nginx ingress-nginx \
  --repo https://kubernetes.github.io/ingress-nginx \
  --namespace ingress-nginx --create-namespace \
  --timeout 5m

kubectl -n ingress-nginx patch svc ingress-nginx-controller -p '{"spec": {"type": "NodePort"}}'
