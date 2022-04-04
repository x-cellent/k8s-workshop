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
1. kleinere Images <!-- .element: class="fragment" data-fragment-index="1" -->
1. Geringerer Ressourcenverbrauch <!-- .element: class="fragment" data-fragment-index="2" -->
1. Erhoehte Sicherheit <!-- .element: class="fragment" data-fragment-index="3" -->
1. Abhaengigkeiten mit im Image <!-- .element: class="fragment" data-fragment-index="4" -->

+++

## Das Dockerfile
1. Datei zum Image bauen

```Dockerfile
FROM alpine:3.9 #base Image
RUN apk add --no-cache mysql-client #Commands welche man ausfuehren möchte, in diesem fall mysql-client installieren
ENTRYPOINT ["mysql"] #Startcommand welcher der container ausfuehren soll
```

+++

## Das Dockerfile !TO_DO!
1. Multi-Stage Dockerfiles auch Moeglich <!-- .element: class="fragment" data-fragment-index="1" -->
1. Vorteile des Multi-Stage Dockerfiles <!-- .element: class="fragment" data-fragment-index="2" -->
    1. Vorteil 1 <!-- .element: class="fragment" data-fragment-index="3" -->
    1. Vorteil 2 <!-- .element: class="fragment" data-fragment-index="4" -->

+++

## Wichtige Docker befehle
1. docker run
1. docker ps 
1. docker logs
1. docker build
1. docker rm
1. docker exec
1. docker --help

+++

## Aufgabe
1. Bitte aufgabe ex1 starten
```sh
bin/k8s-workshop cluster exercise -n 1
```
Zeit: ca 15 min

+++

## Image Builden

```sh
docker build -t IMAGENAME:IMAGETAG ./location/of/docker-file
```

+++

## Aufgabe
1. Bitte aufgabe ex2 starten
```sh
bin/k8s-workshop cluster exercise -n 2
```
Zeit: ca 10 min

+++

## Nachteile von Containarisierung !TO-DO!
1. Fehlende Orchestrierung
1. Fehlende Ausfallsicherheit

---

# Monolith vs Microservices !TO-DO!

---

# Kubernetes
*Kubernetes ist ein Open-Source-System  zur Automatisierung der Bereitstellung, Skalierung und Verwaltung von Container-Anwendungen*

+++

### Kubernetes

1. Urspruenglich 2014 entwickelt von Google
1. Abgegeben 2015 an die Cloud Native Compute Fondation (CNCF)

+++

## Warum Kubernetes? !TO_DO!
1. 

+++

## Architektur von Kubernetes !TO-DO GROESSE ANPASSEN!
![image](https://raw.githubusercontent.com/kubernetes/kubernetes/release-1.3/docs/design/architecture.png)

+++

### Architektur
1. Einzelne Services sind Modular aufgebaut und austauschbar <!-- .element: class="fragment" data-fragment-index="1" -->
    1. API-Server <!-- .element: class="fragment" data-fragment-index="2" -->
    1. Scheudler <!-- .element: class="fragment" data-fragment-index="3" -->
    1. Kubelet <!-- .element: class="fragment" data-fragment-index="4" -->
    1. Kube-Controller-Manager <!-- .element: class="fragment" data-fragment-index="5" -->
    1. Kube-Proxy <!-- .element: class="fragment" data-fragment-index="6" -->
1. Core Services sind Open-Source von der CNCF bereitgestellt <!-- .element: class="fragment" data-fragment-index="7" -->

+++

## Ausfallsicherheit

1. Container Health Check
    1. readyness
    1. liveness
1. Hostsystemausfall
1. Update

+++

## Wichtige Ressourcen
1. kubectl cheat sheet: https://kubernetes.io/docs/reference/kubectl/cheatsheet
1. kubernetes docs: https://kubernetes.io/docs/concepts/overview

---

# Setup

+++

# Install Tools
Kubernetes Dokumentation:​
- kubectl <!-- .element: class="fragment" data-fragment-index="1" -->
- krew <!-- .element: class="fragment" data-fragment-index="2" -->
- helm <!-- .element: class="fragment" data-fragment-index="3" -->
- kind <!-- .element: class="fragment" data-fragment-index="4" -->
- k9s <!-- .element: class="fragment" data-fragment-index="5" -->
