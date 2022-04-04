# Kubernetes Workshop

---

# Agenda

---

<!-- .slide: style="text-align: left;"> -->
## Tag 1
1. Container
1. Kubernetes
1. Setup
    1. Install Tools
    1. Aliases
        - k=kubectl
        - dy='--dry-run=client -o yaml'
    1. Cluster
```sh
git clone https://github.com/x-cellent/k8s-workshop.git
cd k8s-workshop
make
```

---

# Container

+++

## Vorteile Containarisierung
1. kleinere Images
1. Geringerer Ressourcenverbrauch
1. Erhoehte Sicherheit
1. Abhaengigkeiten mit im Image

+++

## Das Dockerfile
1. Datei zum Image bauen

```Dockerfile
FROM alpine:3.9 #base Image
RUN apk add --no-cache mysql-client #Commands welche man ausfuehren möchte, in diesem fall mysql-client installieren
ENTRYPOINT ["mysql"] #Startcommand welcher der container ausfuehren soll
```

---

# Kubernetes

---

# Setup

+++

# Install Tools

- kubectl <!-- .element: class="fragment" data-fragment-index="1" -->
- krew <!-- .element: class="fragment" data-fragment-index="2" -->
- helm <!-- .element: class="fragment" data-fragment-index="3" -->
- kind <!-- .element: class="fragment" data-fragment-index="4" -->
- k9s <!-- .element: class="fragment" data-fragment-index="5" -->
