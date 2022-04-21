# Kubernetes Workshop

---

# Vortragende

+++

- Sandro Koll

+++

- Pascal Rimann

+++

## Teilnehmer

+++

- Kurze Vorstellung
- Erfahrungen?
    - Docker
    - Kubernetes
- Erwartungen?

---

# Tag 1

+++

<!-- .slide: style="text-align: left;"> -->
## Agenda
1. Setup
1. Container
1. Monolithen vs. Microservices
1. Container-Orchestrierung
1. Prinzipien hinter Kubernetes 

---

<!-- .slide: style="text-align: left;"> -->
## Setup

```sh
sudo usermod -aG docker ${USER}
git clone https://github.com/x-cellent/k8s-workshop.git
cd k8s-workshop
make
mkdir -p ~/bin
mv bin/w6p ~/bin/
echo "export PATH=$PATH:~/bin" >> ~/.bashrc
source ~/.bashrc
w6p
```

+++

<!-- .slide: style="text-align: left;"> -->
## Output

```sh
Usage:
  w6p [flags]
  w6p [command]

Available Commands:
  cluster     Runs the workshop cluster or exercises
  exercise    Runs the given exercise
  help        Help about any command
  install     Installs tools on local machine
  slides      Shows or exports workshop slides

Flags:
  -h, --help   help for w6p
```

+++

<!-- .slide: style="text-align: left;"> -->
## Was ist w6p?

Go CLI executable ausschließlich für diesen Workshop  <!-- .element: class="fragment" data-fragment-index="1" -->
- w6p install TOOL  <!-- .element: class="fragment" data-fragment-index="2" -->
    - lokale Installation von gebräuchlichen k8s Tools  <!-- .element: class="fragment" data-fragment-index="2" -->
- w6p exercise CONTEXT -n NUMBER  <!-- .element: class="fragment" data-fragment-index="3" -->
    - Startet Aufgaben aus dem jeweiligen Kontext (docker oder k8s)  <!-- .element: class="fragment" data-fragment-index="3" -->
- w6p cluster  <!-- .element: class="fragment" data-fragment-index="4" -->
    - Startet/stoppt Single-Node Kubernetes Cluster in Container  <!-- .element: class="fragment" data-fragment-index="4" -->

+++

<!-- .slide: style="text-align: left;"> -->
- w6p slides
    - Startet Webserver Container, der die Workshop Slides hosted
    - erreichbar unter localhost:8080

---

<!-- .slide: style="text-align: left;"> -->
## Container

Software wird schon seit Jahrzehnten in Archive oder Single-Binaries verpackt
- Einfache Auslieferung  <!-- .element: class="fragment" data-fragment-index="1" -->
- Einfache Verteilung  <!-- .element: class="fragment" data-fragment-index="2" -->

+++

<!-- .slide: style="text-align: left;"> -->
## Aber
- Installation notwendig  <!-- .element: class="fragment" data-fragment-index="1" -->
- Dependency Hell  <!-- .element: class="fragment" data-fragment-index="2" -->
- No cross platform functionality  <!-- .element: class="fragment" data-fragment-index="3" -->

+++

<!-- .slide: style="text-align: left;"> -->
## Lösung
- Verpacken der Software mitsamt aller Dependencies (Image)  <!-- .element: class="fragment" data-fragment-index="1" -->
    - Nichts darüber hinaus (Betriebssytem notwendig?)  <!-- .element: class="fragment" data-fragment-index="2" -->
- Container-Runtime für alle Plattformen  <!-- .element: class="fragment" data-fragment-index="3" -->

+++

<!-- .slide: style="text-align: left;"> -->
## Umsetzung
- Linux  <!-- .element: class="fragment" data-fragment-index="1" -->
- Idee: Container teilen sich Kernel  <!-- .element: class="fragment" data-fragment-index="2" -->
- LXC: basierend auf Kernel-Funktionalitäten  <!-- .element: class="fragment" data-fragment-index="3" -->
    - namespaces  <!-- .element: class="fragment" data-fragment-index="4" -->
    - cgroups  <!-- .element: class="fragment" data-fragment-index="5" -->
- Docker erweitert LXC um  <!-- .element: class="fragment" data-fragment-index="6" -->
    - ...CLI zum Starten und Verwalten von Containern  <!-- .element: class="fragment" data-fragment-index="6" -->
    - Image Registry  <!-- .element: class="fragment" data-fragment-index="7" -->
    - Networking  <!-- .element: class="fragment" data-fragment-index="8" -->
    - docker-compose  <!-- .element: class="fragment" data-fragment-index="9" -->

<aside class="notes">
  Kein komplettes OS installiert wird
  Kernel vom Host System wird geteilt
    - namespaces: Isolierung Prozesse (Ressourcen-Sichtbarkeit), Microservice-Architekturstil
    - cgroups: Resource-Limits für Prozesse (CPU, Memory); CPU, Disk und Netzwerk-Ressourcen Aufteilung
</aside>

---

<!-- .slide: style="text-align: left;"> -->
## Vorteile Container
1. Geringere Größe <!-- .element: class="fragment" data-fragment-index="1" -->
1. Erhöhte Sicherheit <!-- .element: class="fragment" data-fragment-index="2" -->
1. Funktional auf allen Systemen <!-- .element: class="fragment" data-fragment-index="3" -->
1. Immutable <!-- .element: class="fragment" data-fragment-index="4" -->
    - Damit Baukastenprinzip möglich (DRY) <!-- .element: class="fragment" data-fragment-index="5" -->

<aside class="notes">
  Beinhaltet kein komplettes OS =>
  Ubuntu Image ca. 72.8 MB
  Alpine ca. 5.57 MB
  Busybox ca. 1.24 MB
</aside>

+++

<!-- .slide: style="text-align: left;"> -->
## Vorteile gegenüber VMs
1. Geringere Größe <!-- .element: class="fragment" data-fragment-index="1" -->
1. Geringerer Ressourcenverbrauch <!-- .element: class="fragment" data-fragment-index="2" -->
1. Viel schnellere Startup-Zeiten <!-- .element: class="fragment" data-fragment-index="3" -->
1. Auch geeignet für Entwicklung und Test <!-- .element: class="fragment" data-fragment-index="4" -->

<aside class="notes">
  Ein Rechner kann deutlich mehr Anwendungen als VMs hosten
</aside>

+++

<!-- .slide: style="text-align: left;"> -->
## Nachteile gegenüber VMs
1. Geringere Sicherheit <!-- .element: class="fragment" data-fragment-index="1" -->
1. Keine echte Trennung <!-- .element: class="fragment" data-fragment-index="2" -->
    - z.B. kein Block-Storage möglich <!-- .element: class="fragment" data-fragment-index="2" -->

Container und VMs schließen sich aber nicht gegenseitig aus <!-- .element: class="fragment" data-fragment-index="3" -->

---

<!-- .slide: style="text-align: left;"> -->
## Docker Komponenten
1. Image <!-- .element: class="fragment" data-fragment-index="1" -->
    - Layer <!-- .element: class="fragment" data-fragment-index="2" -->
    - Dockerfile <!-- .element: class="fragment" data-fragment-index="3" -->
1. Container <!-- .element: class="fragment" data-fragment-index="4" -->
1. Image Registry <!-- .element: class="fragment" data-fragment-index="5" -->

+++

<!-- .slide: style="text-align: left;"> -->
## Dockerfile
- [Referenz](https://docs.docker.com/engine/reference/builder/)
- Image-*Rezept* mit u.a. folgenden *Zutaten*:
    - FROM <!-- .element: class="fragment" data-fragment-index="1" -->
    - COPY/ADD <!-- .element: class="fragment" data-fragment-index="2" -->
    - RUN <!-- .element: class="fragment" data-fragment-index="3" -->
    - USER <!-- .element: class="fragment" data-fragment-index="4" -->
    - WORKDIR <!-- .element: class="fragment" data-fragment-index="5" -->
    - ARG/ENV <!-- .element: class="fragment" data-fragment-index="6" -->
    - ENTRYPOINT <!-- .element: class="fragment" data-fragment-index="7" -->
    - CMD <!-- .element: class="fragment" data-fragment-index="8" -->

+++

<!-- .slide: style="text-align: left;"> -->
## Beispiel

```Dockerfile stretch
# Basis-Image
FROM alpine:3.15

# Installiert busybox-extras ins Basis-Image
RUN apk add --no-cache busybox-extras
# ...und committed den FS-Diff als neuen Layer

# Installiert auf dem obigen Layer mysql-client
RUN apk add --no-cache mysql-client
# ...und commited das FS-Diff als neuen Layer

# Default command
ENTRYPOINT ["mysql"]

# Default arg(s)
CMD ["--help"]
```

+++

<!-- .slide: style="text-align: left;"> -->
## Einige Docker Commands
- docker build <!-- .element: class="fragment" data-fragment-index="1" -->
    - Baut ein Image von Dockerfile <!-- .element: class="fragment" data-fragment-index="1" -->
- docker images / docker image ls <!-- .element: class="fragment" data-fragment-index="2" -->
    - Listet alle (lokalen) Images <!-- .element: class="fragment" data-fragment-index="2" -->
- docker tag <!-- .element: class="fragment" data-fragment-index="3" -->
    - Erstellt Image "Kopie" unter anderem Namen <!-- .element: class="fragment" data-fragment-index="3" -->
- docker rmi / docker image rm <!-- .element: class="fragment" data-fragment-index="4" -->
    - Löscht ein Image <!-- .element: class="fragment" data-fragment-index="4" -->
- docker login/logout <!-- .element: class="fragment" data-fragment-index="5" -->
- docker push/pull <!-- .element: class="fragment" data-fragment-index="6" -->

+++

<!-- .slide: style="text-align: left;"> -->
- docker [run](https://docs.docker.com/engine/reference/run/)
    - Startet ein Image -> Container
- docker ps [-q] <!-- .element: class="fragment" data-fragment-index="1" -->
    - Listet alle (laufenden) Container <!-- .element: class="fragment" data-fragment-index="1" -->
- docker rm <!-- .element: class="fragment" data-fragment-index="2" -->
    - Löscht einen Container <!-- .element: class="fragment" data-fragment-index="2" -->
- docker logs <!-- .element: class="fragment" data-fragment-index="3" -->
    - Zeigt Container Logs <!-- .element: class="fragment" data-fragment-index="3" -->
- docker exec <!-- .element: class="fragment" data-fragment-index="4" -->
    - Führt Befehl in laufendem Container aus <!-- .element: class="fragment" data-fragment-index="4" -->

+++

<!-- .slide: style="text-align: left;"> -->
- docker [image] inspect
    - Zeigt Metadaten von Container/Images
- docker cp <!-- .element: class="fragment" data-fragment-index="1" -->
    - Kopiert eine Datei aus Container ins Host-FS und umgekehrt <!-- .element: class="fragment" data-fragment-index="1" -->
- docker save/load <!-- .element: class="fragment" data-fragment-index="2" -->
    - Erzeugt Tarball aus Image und umgekehrt <!-- .element: class="fragment" data-fragment-index="2" -->
- docker network <!-- .element: class="fragment" data-fragment-index="3" -->
    - Netzwerk Management <!-- .element: class="fragment" data-fragment-index="3" -->

+++

<!-- .slide: style="text-align: left;"> -->
## Image bauen

```sh
docker build -t [REPOSITORY_HOST/]IMAGENAME:IMAGETAG \
    [-f path/to/Dockerfile] path/to/context-dir
```

- Kontext-Verzeichnis wird zum Docker Daemon hochgeladen <!-- .element: class="fragment" data-fragment-index="1" -->
    - lokal oder remote (via DOCKER_HOST) <!-- .element: class="fragment" data-fragment-index="1" -->
    - Nur darin enthaltene Dateien können im Dockerfile verwendet werden (COPY/ADD) <!-- .element: class="fragment" data-fragment-index="2" -->
    - Nach Möglichkeit keine ungenutzten Dateien hochladen <!-- .element: class="fragment" data-fragment-index="3" -->

+++

<!-- .slide: style="text-align: left;"> -->
## Aufgabe 1

```sh
w6p exercise docker -n 1
```
Lösung nach 5 min

+++

<!-- .slide: style="text-align: left;"> -->
## Container starten

```sh
docker run [--name NAME] [-i] [-t] [-d|--rm] [--net host|NETWORK] [-v HOST_PATH:CONTAINER_PATH] \
    [-p HOST_PORT:CONTAINER_PORT] [-u UID:GID] IMAGE [arg(s)]
```

- Mehr Optionen möglich
- [Referenz](https://docs.docker.com/engine/reference/run/)

+++

<!-- .slide: style="text-align: left;"> -->
## Execute CMD in Container

```sh
docker exec [-i] [-t] CONTAINER COMMAND
```

Via Bash in den Container "springen":

```sh
docker exec -it CONTAINER /bin/bash
```

+++

## Dateien kopieren

...vom Host in den Container:

```sh
docker cp HOST_FILE CONTAINER_NAME:CONTAINER_FILE
```

z.B.:

```sh stretch
docker cp ~/local-index.html my-server:/static/index.html
```

+++

...vom Container in das Host-FS:

```sh
docker cp CONTAINER_NAME:CONTAINER_FILE HOST_FILE
```

z.B.:

```sh stretch
docker cp my-server:/static/index.html ~/local-index.html
```

+++

<!-- .slide: style="text-align: left;"> -->
## Aufgabe 2

```sh
w6p exercise docker -n2
```
Lösung nach 20 min

---

<!-- .slide: style="text-align: left;"> -->
## Metadaten

```sh
docker inspect IMAGE|CONTAINER
```

- Image Metadaten
    - ID
    - Architecture
    - Layers
    - Env
    - ...

+++

<!-- .slide: style="text-align: left;"> -->
- Container Metadaten
    - ID
    - Image ID
    - NetworkSettings
    - Mounts
    - State
    - ...

+++

<!-- .slide: style="text-align: left;"> -->
## Aufgabe 3

```sh
w6p exercise docker -n3
```
Lösung nach 15 min

---

<!-- .slide: style="text-align: left;"> -->
## Linting

- [hadolint](https://github.com/hadolint/hadolint)
- Erhältlich als Docker Image:

```sh
docker run ... hadolint/hadolint hadolint path/to/Dockerfile
```

+++

<!-- .slide: style="text-align: left;"> -->
## Aufgabe 4

```sh
w6p exercise docker -n4
```
Lösung nach 5 min

---

<!-- .slide: style="text-align: left;"> -->
## Multi-Stage Dockerfile

```Dockerfile stretch
# Build as usual in the first stage
FROM golang:1.17 AS builder # named stage

WORKDIR /work
# Copy source code into image
COPY app.go .

# Compile source(s)
RUN go build -o bin/my-app

# Further builder images possible, e.g.
# FROM nginx AS webserver
# ...

# Final stage: all prior images will be discarded after build
FROM scratch

# ...but here we can copy files from builder image(s)
COPY --from=builder /work/bin/my-app /

ENTRYPOINT ["/my-app"]
```

+++

<!-- .slide: style="text-align: left;"> -->
## Vorteile
1. Kompakte Imagegröße <!-- .element: class="fragment" data-fragment-index="1" -->
1. Erhöhte Sicherheit <!-- .element: class="fragment" data-fragment-index="2" -->

+++

<!-- .slide: style="text-align: left;"> -->
## Aufgabe 5

```sh
w6p exercise docker -n5
```
Lösung nach 15 min

---

<!-- .slide: style="text-align: left;"> -->
## Docker Registry
- [Docker-Hub](https://hub.docker.com/) <!-- .element: class="fragment" data-fragment-index="1" -->
    - öffentlich <!-- .element: class="fragment" data-fragment-index="2" -->
- private Registries möglich <!-- .element: class="fragment" data-fragment-index="3" -->
    - Image [registry](https://hub.docker.com/_/registry) <!-- .element: class="fragment" data-fragment-index="4" -->
    - absicherbar <!-- .element: class="fragment" data-fragment-index="5" -->
    - praktisch in jeder Firma eingesetzt <!-- .element: class="fragment" data-fragment-index="6" -->

+++

<!-- .slide: style="text-align: left;"> -->
## Aufgabe 6

```sh
w6p exercise docker -n6
```
Lösung nach 15 min

---

<!-- .slide: style="text-align: left;"> -->
## Was Docker nicht bietet:
1. Orchestrierung  <!-- .element: class="fragment" data-fragment-index="1" -->
1. Ausfallsicherheit  <!-- .element: class="fragment" data-fragment-index="2" -->

=> Kubernetes <!-- .element: class="fragment" data-fragment-index="3" -->
    - bietet beides  <!-- .element: class="fragment" data-fragment-index="3" -->
    - ...und noch viel mehr  <!-- .element: class="fragment" data-fragment-index="4" -->

+++

<!-- .slide: style="text-align: left;"> -->
## Docker Compose

- Für Multi-Container Docker Anwendungen <!-- .element: class="fragment" data-fragment-index="1" -->
- docker-compose.yaml <!-- .element: class="fragment" data-fragment-index="2" -->
    - Definition der Container <!-- .element: class="fragment" data-fragment-index="2" -->
- docker-compose up/down  <!-- .element: class="fragment" data-fragment-index="3" -->
    - Start/Stop aller Anwendungen in einem Rutsch <!-- .element: class="fragment" data-fragment-index="3" -->
- Rudimentäre Funktionalitäten <!-- .element: class="fragment" data-fragment-index="4" -->
- Geeignet für sehr kleine (Dev-/Test-)Umgebungen <!-- .element: class="fragment" data-fragment-index="5" -->
    - schneller als schnellstes K8s-Setup ([k3s](https://k3s.io/) oder [kind](https://kind.sigs.k8s.io/docs/user/quick-start/))  <!-- .element: class="fragment" data-fragment-index="5" -->
- Mittel der Wahl ist aber Kubernetes <!-- .element: class="fragment" data-fragment-index="6" -->

---

# Monolith vs Microservices 

+++

<!-- .slide: style="text-align: left;" class="stretch"> -->
![image](https://i0.wp.com/mjaglan.github.io/images/docker-virtualbox/docker-vs-vm.png?w=840&ssl=1)

<aside class="notes">
  wie man auf diesem bild erkennen kann, ist eine microservice struktur deutlich kleiner als ein Monolith system

  liegt daran, dass sich alle systeme ein host os teilen können

  dazu zählen kernel aufgaben

  deployments schneller da nicht alle prozesse beendet und neugestartet werden müssen
</aside>

---

# Container Orchestrierung

+++

## Wieso?
- Orchestrierung von Containern <!-- .element: class="fragment" data-fragment-index="1" -->

<aside class="notes">
  Unter Orchestrierung versteht man das Deployment, Maintenance und Scaling

  Vorteile: Bessere Ressourcennutzung

  Bessere Bereitsstellung von Containern -> zero downtime rollouts
</aside>

+++

## Warum Kubernetes?
- Warum nicht Docker Swarm? <!-- .element: class="fragment" data-fragment-index="1" -->
- Mehr Flexibilität <!-- .element: class="fragment" data-fragment-index="2" -->
- Eingebautes Monitoring und Logging <!-- .element: class="fragment" data-fragment-index="3" -->
- Bereitstellung von Storage <!-- .element: class="fragment" data-fragment-index="4" -->
- Größere User-Base <!-- .element: class="fragment" data-fragment-index="5" -->

<aside class="notes">
  da wir einiges über docker gehört haben kann man sich fragen warum nicht docker swarm?

  es hat ja eine leichtere installation und kommt aus dem gleichen haus wie docker

  bei kubernetes hat man eine größere flexibilität, sodass man anwendungen bereitstellen kann, wie man möchtet

  kubernetes hat ein eingebautes monitoring und logging, docker swarm nicht

  speicher kann man bei kubernetes so hinzufügen, dass es als einzeiler genutzt werden kann
</aside>

---

# Prinzipien hinter Kubernetes

+++

### Der Pod
- kein Container
- beinhaltet mindestens einen Container
- kann init container beinhalten
- kann sidecar container beinhalten
- kleinste Einheit in Kubernetes

<aside class="notes">
  oft werden pods mit containern verglichen

  aber ein pod ist grundsätzlich eine gruppe (gruppe von walen vergleich)

  teilen sich speicher und Netzwerkressourcen

  befinden sich auf dem selber server

  erhält einen oder mehrere container > abhängig voneinader

  init container kann script vor start ausführen

  sidecar container kann daten im pod aktualisieren

  dazu morgen mehr
</aside>

+++

### Wo laufen Pods?
- Auf (Worker-)Nodes

<aside class="notes">
  Die pods laufen auf sogenannten nodes, dies sind die server auf welchen die diversen Kubernetes core programme laufen, dies wird morgen genauer behandelt
</aside>

+++

### Ordnungselemente
- ReplicaSet <!-- .element: class="fragment" data-fragment-index="1" -->
- Deployment <!-- .element: class="fragment" data-fragment-index="2" -->
- StatefulSet <!-- .element: class="fragment" data-fragment-index="3" -->
- DaemonSet <!-- .element: class="fragment" data-fragment-index="4" -->
- Job <!-- .element: class="fragment" data-fragment-index="5" -->
- CronJob <!-- .element: class="fragment" data-fragment-index="6" -->

<aside class="notes">
  Wie man sehen kann, ist das wichtigste Element der Pod

  Selten setzt man ihn einzeln ein

  Normalerweise nutzt man ein übergeortnetes Ordnungselement:

  ReplicaSet: Stellt eine genaue Anzahl an Pods sicher, wird äußerst selten explizit verwenden. Statt dessen:

  Deployment: Kombiniert ReplicaSets mit Versionierung für `stateless` Pods und bietet die Möglichkeit zum `staged rollout`

  StatefulSet: Wie Deployment, nur für für `stateful` Pods, d.h. Pods, die zur Laufzeit CRUD Operationen auf persistenten Daten ausführen

  DaemonSet: Garantiert, dass auf allen Nodes genau ein Pod läuft

  Job: Ein Pod, der nur einmal gestartet wird, um eine bestimmte Aufgabe zu erledigen

  CronJob: Für wiederkehrende Tasks zu bestimmten Zeitpunkten, z.B. um regelmäßige Backups zu erstellen
</aside>


---

# Fragen
- Hab ihr noch Fragen an uns?

---

# Ausblick Tag 2
- Architektur von Kubernetes
- Basis Objekte von Kubernetes

---

# Tag 2

+++

<!-- .slide: style="text-align: left;"> -->
## Agenda
1. Rewind
1. Architektur von Kubernetes
1. Einrichtung eurer Umgebung
1. kubectl
1. k9s
1. Basisobjekte Kubernetes

---

# Rewind

+++

<!-- .slide: style="text-align: left;"> -->
### Docker
- Container Runtime Engine <!-- .element: class="fragment" data-fragment-index="1" -->
- docker CLi <!-- .element: class="fragment" data-fragment-index="2" -->
    - Bau von Images <!-- .element: class="fragment" data-fragment-index="3" -->
    - Start von Images (Container) <!-- .element: class="fragment" data-fragment-index="4" -->
    - Image Registry (Verteilung von Images, vgl. AppStore) <!-- .element: class="fragment" data-fragment-index="5" -->
    - ... <!-- .element: class="fragment" data-fragment-index="5" -->

+++

<!-- .slide: style="text-align: left;"> -->
### Image bauen
- Schreibe ein Dockerfile (Rezept) <!-- .element: class="fragment" data-fragment-index="1" -->
    - Instruktionen (Zutaten) <!-- .element: class="fragment" data-fragment-index="2" -->
        - FROM, COPY, RUN, ENTRYPOINT, ... <!-- .element: class="fragment" data-fragment-index="3" -->
        - jede Instruktion = neuer Layer (vgl. mit Binärdatei) <!-- .element: class="fragment" data-fragment-index="4" -->
- docker build [-t TAG] CONTEXT <!-- .element: class="fragment" data-fragment-index="5" -->
    - Image = N übereinander gelegte RO Layer (overlay FS) <!-- .element: class="fragment" data-fragment-index="6" -->

+++

<!-- .slide: style="text-align: left;"> -->
### Image starten
- docker run [OPTS] IMAGE [COMMAND] [ARGS]
    - Container = N Image RO Layers plus ein leerer RW Layer on top <!-- .element: class="fragment" data-fragment-index="1" -->
    - RW Layer kann zur Laufzeit modifiziert werden <!-- .element: class="fragment" data-fragment-index="2" -->
    - Delete File/Dir <!-- .element: class="fragment" data-fragment-index="3" -->
        - Löscht, wenn im RW Layer vorhanden <!-- .element: class="fragment" data-fragment-index="4" -->
        - Versteckt, wenn in einem darunter liegenden RO Layer <!-- .element: class="fragment" data-fragment-index="5" -->
    - Stirbt PID 1, stirbt der Container <!-- .element: class="fragment" data-fragment-index="6" -->
        - docker rm CONTAINER => RW Layer wird gelöscht <!-- .element: class="fragment" data-fragment-index="7" -->

+++

<!-- .slide: style="text-align: left;"> -->
### Kubernetes
- Container Orchestrierungstool <!-- .element: class="fragment" data-fragment-index="1" -->
    - Verwaltet Pods <!-- .element: class="fragment" data-fragment-index="2" -->
        - besteht aus 1 bis N Containern <!-- .element: class="fragment" data-fragment-index="3" -->
        - eigene IP, eigenes Netzwerk <!-- .element: class="fragment" data-fragment-index="4" -->
    - Started, stoppt und überwacht Pods <!-- .element: class="fragment" data-fragment-index="5" -->
        - Verteilung auf Worker-Nodes <!-- .element: class="fragment" data-fragment-index="6" -->
        - Garantiert Pods Ressourcen (CPU/Memory) <!-- .element: class="fragment" data-fragment-index="7" -->
    - Self-Healing <!-- .element: class="fragment" data-fragment-index="8" -->
    - Dynamische Skalierung <!-- .element: class="fragment" data-fragment-index="9" -->
    - ... <!-- .element: class="fragment" data-fragment-index="9" -->

+++

<!-- .slide: style="text-align: left;"> -->
## Kubernetes Objekte
- Pod
    - kleinste deploybare Einheit
    - beinhaltet 1 bis N Container
    - eigener Netzbereich und IP
- ReplicaSet <!-- .element: class="fragment" data-fragment-index="1" -->
    - Stellt sicher, dass zu jeder Zeit genau N Pods laufen <!-- .element: class="fragment" data-fragment-index="2" -->
    - Matching über Labels <!-- .element: class="fragment" data-fragment-index="3" -->

+++

<!-- .slide: style="text-align: left;"> -->
- Deployment
    - Managed ein ReplicaSet
    - Bietet Versionierung und zero downtime Rollouts
- DeamonSet <!-- .element: class="fragment" data-fragment-index="1" -->
    - Spec wie Deployment nur ohne Replica Count <!-- .element: class="fragment" data-fragment-index="1" -->
    - Managed damit kein ReplicaSet <!-- .element: class="fragment" data-fragment-index="1" -->
    - Stattdessen je ein Replica pro Node <!-- .element: class="fragment" data-fragment-index="1" -->

+++

<!-- .slide: style="text-align: left;"> -->
- Service
    - Loadbalancer für Pods
    - Auch hier Matching via Labels
    - Typen <!-- .element: class="fragment" data-fragment-index="1" -->
        - None (headless) <!-- .element: class="fragment" data-fragment-index="1" -->
        - ClusterIP (Default) <!-- .element: class="fragment" data-fragment-index="2" -->
        - NodePort <!-- .element: class="fragment" data-fragment-index="3" -->
        - Loadbalancer <!-- .element: class="fragment" data-fragment-index="4" -->

Note:
Service - headless (erstellt für jeden Pod einen DNS entry innerhalb des Clusters (coredns), kein externer Zugriff möglich)

Service - ClusterIP (routet über die clusterinternen Pod IPs, kein externer Zugiff möglich)

Service - NodePort (öffnet auf jedem Node denselben Port, über den von außen der Service erreicht werden kann)

Service - Loadbalancer (exosed den Service ins Internet, bedarf eines Loadbalancers der den Traffic an der Service weiterleitet)

+++

<!-- .slide: style="text-align: left;"> -->
- StatefulSet
    - Spec ähnlich zu Deployment
    - Geordnetes Starten (einer nach dem anderen)
    - Geordnetes Stoppen in umgekehrter Reihenfolge
- ConfigMap <!-- .element: class="fragment" data-fragment-index="1" -->
    - Plain-Text Key-Value Store <!-- .element: class="fragment" data-fragment-index="1" -->
    - Kann in Pods, Deployments, STSs und DSs gemounted werden <!-- .element: class="fragment" data-fragment-index="2" -->
- Secret <!-- .element: class="fragment" data-fragment-index="3" -->
    - base64 encoded Data Store <!-- .element: class="fragment" data-fragment-index="3" -->
    - Kann in Pods, Deployments, STSs und DSs gemounted werden <!-- .element: class="fragment" data-fragment-index="4" -->

+++

<!-- .slide: style="text-align: left;"> -->
### Kubernetes Tools
- kubectl <!-- .element: class="fragment" data-fragment-index="1" -->
    - CLI zur Interaktion mit k8s Clustern <!-- .element: class="fragment" data-fragment-index="1" -->
- krew <!-- .element: class="fragment" data-fragment-index="2" -->
    - kubectl Plugin Manager <!-- .element: class="fragment" data-fragment-index="2" -->
- k9s <!-- .element: class="fragment" data-fragment-index="3" -->
    - Terminal UI zur Interaktion mit k8s Clustern <!-- .element: class="fragment" data-fragment-index="3" -->
- kind (Kubernetes in Docker) <!-- .element: class="fragment" data-fragment-index="4" -->
    - Single-Node k8s Cluster in Docker Container <!-- .element: class="fragment" data-fragment-index="4" -->

+++

<!-- .slide: style="text-align: left;"> -->
- helm
    - Paket-Manager für Kubernetes
    - vgl. mit apt für Ubuntu oder apk für Alpine
- Lens <!-- .element: class="fragment" data-fragment-index="1" -->
    - Graphical UI zur Interaktion mit k8s Clustern <!-- .element: class="fragment" data-fragment-index="1" -->
    - Nicht Teil des Workshops <!-- .element: class="fragment" data-fragment-index="1" -->

+++

<!-- .slide: style="text-align: left;"> -->
## kubectl Plugins

- [Liste](https://krew.sigs.k8s.io/plugins/) von verfügbaren Plugins
- Installation
    - kubectl krew install NAME [NAME...]

+++

<!-- .slide: style="text-align: left;"> -->
## Was ist w6p?

Go CLI executable ausschließlich für diesen Workshop  <!-- .element: class="fragment" data-fragment-index="1" -->
- w6p install TOOL  <!-- .element: class="fragment" data-fragment-index="2" -->
    - lokale Installation von gebräuchlichen k8s Tools  <!-- .element: class="fragment" data-fragment-index="2" -->
- w6p exercise CONTEXT -n NUMBER  <!-- .element: class="fragment" data-fragment-index="3" -->
    - Startet Aufgaben aus dem jeweiligen Kontext (docker oder k8s)  <!-- .element: class="fragment" data-fragment-index="3" -->
- w6p cluster  <!-- .element: class="fragment" data-fragment-index="4" -->
    - Startet/stoppt Single-Node Kubernetes Cluster in Container  <!-- .element: class="fragment" data-fragment-index="4" -->

---

<!-- .slide: style="text-align: left;"> -->
## Kubernetes

- Ursprünglich 2014 entwickelt von Google
- Abgegeben 2015 an die Cloud Native Compute Fondation (CNCF)

<aside class="notes">
  Weiterentwicklung von Google Borg

  aktuelle version 1.23.5
</aside>

+++

<!-- .slide: style="text-align: left;"> -->
### CNCF

- Cloud Native Computing Foundation
- 2015 Gegründet
- äber 500 Hersteller und Betreiber

<aside class="notes">
  CNCF ist die Cloud Native Computung Foundation

  untergeordnet der Linux Fondation

  Größte Unternehmen sind Amazon, Google, Apple, Microsoft

  x-cellent ist auch teil der CNCF
</aside>

+++

<!-- .slide: style="text-align: left;" class="stretch"> -->
## Architektur von Kubernetes
![image](https://upload.wikimedia.org/wikipedia/commons/b/be/Kubernetes.png)

<aside class="notes">
  2 arten der nodes, master/controlPlane und worker
</aside>

+++

<!-- .slide: style="text-align: left;"> -->
### Architektur des Clusters
1. Modular und austauschbar <!-- .element: class="fragment" data-fragment-index="1" -->
    1. Control-Plane <!-- .element: class="fragment" data-fragment-index="2" -->
        - etcd <!-- .element: class="fragment" data-fragment-index="3" -->
        - API-Server <!-- .element: class="fragment" data-fragment-index="4" -->
        - Scheduler <!-- .element: class="fragment" data-fragment-index="5" -->
        - Kube-Controller-Manager <!-- .element: class="fragment" data-fragment-index="6" -->
    1. Nodes <!-- .element: class="fragment" data-fragment-index="7" -->
        - Kubelet <!-- .element: class="fragment" data-fragment-index="8" -->
        - Kube-Proxy <!-- .element: class="fragment" data-fragment-index="9" -->
1. Open-Source  <!-- .element: class="fragment" data-fragment-index="10" -->

<aside class="notes">
  Das Kubernetes Konstrukt ist modular aufgebaut und komplett austauschbar.

  zusammengefasst werden einige diese unter dem Namen "ControlPlane" welche nur auf dem Master server laufen

  Kernkomponenten sind ...

  bei Nodes zählen alle, sowohl master als auch worker nodes

  Diese Komponenten sind komplett Open Source

  gleich kommen einzelheiten zu diesen Komponenten
</aside>

+++

<!-- .slide: style="text-align: left;"> -->
## Control-Plane

<aside class="notes">
die Controle Plane Server sind die nodes, welche für die Verwaltung des Clusters zuständig sind, normalerweise keine container auf diesen nodes
</aside>

+++

<!-- .slide: style="text-align: left;"> -->
### etcd
- entwickelt von CoreOS <!-- .element: class="fragment" data-fragment-index="1" -->
- key-value Database <!-- .element: class="fragment" data-fragment-index="2" -->
- kann nicht getauscht werden <!-- .element: class="fragment" data-fragment-index="3" -->
- speichert stand von cluster <!-- .element: class="fragment" data-fragment-index="4" -->
- Consistency notwending <!-- .element: class="fragment" data-fragment-index="5" -->

<aside class="notes">
  wird entwickelt und maintaint von coreos team

  open source tool

  quasi hardcoded in kubernetes core

  wichtigste K8s komponente

  speichert configuration, status, und alle metadaten

  wenn man etcd backup in neuen Cluster einspielt, baut es den cluster wie zuvor auf

  Consistency, also Wiedersruchsfreiheit ist beim etcd notwendig, anosnsten kann es zu ausfällen des Clusters kommen
</aside>

+++

<!-- .slide: style="text-align: left;"> -->
### API-Server
- Ansprechpunkt des Users <!-- .element: class="fragment" data-fragment-index="1" -->
- Validation der Daten <!-- .element: class="fragment" data-fragment-index="2" -->
- bekanntester ist kube-apiserver <!-- .element: class="fragment" data-fragment-index="3" -->
- horizontale skalierbarkeit <!-- .element: class="fragment" data-fragment-index="4" -->

<aside class="notes">
  immer wenn ihr im Cluster was arbeitet sprecht ihr mit dem API Server, egal ob mit kubectl, helm, k9s etc..

  Validiert, ob rechte vorhanden sind, und anfragen sinn ergeben

  updated values in etcd

  bekanntester api server tool kube-apiserver
</aside>

+++

<!-- .slide: style="text-align: left;"> -->
#### scheduler
- Verteilt workload <!-- .element: class="fragment" data-fragment-index="1" -->
- verantworlich für pods ohne node <!-- .element: class="fragment" data-fragment-index="2" -->

<aside class="notes">
  sobald ein neuer Pod am API Server erstellt wurde, von diesem in die etcd db geschrieben wurde

  nimmt der scheudler die werte und verteilt diese an nodes

  Faktoren bei entscheidung welche node

  Ressourcenanforderungen

  Hard/Software-einschränkungen

  bestimmte Flags z.B. niemals auf master/nur auf master
</aside>

+++

<!-- .slide: style="text-align: left;"> -->
### Kube-Controller-Manager
- bringt cluster von "ist" -> "soll" <!-- .element: class="fragment" data-fragment-index="1" -->
- Managed Nodes <!-- .element: class="fragment" data-fragment-index="2" -->
- mitteilung an scheuduler wenn node down <!-- .element: class="fragment" data-fragment-index="3" -->

<aside class="notes">
  bekommt von scheduler meldung, wenn pod zu node kommen soll und übermittelt dies

  überwacht nodes, wenn einer down ist, mitteilung an scheuduler, node down bitte pods neu verteilen
</aside>

+++

<!-- .slide: style="text-align: left;"> -->
## Nodes

<aside class="notes">
  die nachfolgende Software ist auf allen nodes, also master und worker installiert
<aside>

+++

<!-- .slide: style="text-align: left;"> -->
### Kubelet
- verwaltet pods <!-- .element: class="fragment" data-fragment-index="1" -->
- auf jeden node installiert <!-- .element: class="fragment" data-fragment-index="2" -->
- verantwortlich für status <!-- .element: class="fragment" data-fragment-index="3" -->

<aside class="notes">
  kubelet bekommt info von controller und fürt auf node aus

  startet stoppt updated pods auf node

  überwacht pods ob sie gewünschten status haben
</aside>

+++

<!-- .slide: style="text-align: left;"> -->
### Kube Proxy
- verwaltet Netzwerkanfragen <!-- .element: class="fragment" data-fragment-index="1" -->
- routet traffic zu gewünschten pod <!-- .element: class="fragment" data-fragment-index="2" -->
- loadbalancer <!-- .element: class="fragment" data-fragment-index="3" -->

<aside class="notes">
  der kube-proxy wird angefragt sobald eine Netzwerkanfrage zum node kommt und leitet diese weiter zum gewünschten container

  übernimmt auch loadbalancing funktionen, sollten meherere Pods das gleiche machen (mehere nginx zum beispiel)
</aside>

+++

<!-- .slide: style="text-align: left;"> -->
## Weitere Komponenten
- CNI <!-- .element: class="fragment" data-fragment-index="1" -->
- Container-Runtime <!-- .element: class="fragment" data-fragment-index="2" -->

<aside class="notes">
  es gibt noch weiter komponente

  Network Plugin, also CNI schreibt Network interfaces in die

  Container Runtime

  dies ist z.B. containerd, cri-o oder die deprecated docker engine

  meistens containerd

  diese Software ist zuständig um die Container laufen zu lassen
</aside> 

+++

<!-- .slide: style="text-align: left;"> -->
## Namespaces
- separierungseinheit in Kubernetes  <!-- .element: class="fragment" data-fragment-index="1" -->
- Objekte können welche in anderem Namespace nicht sehen  <!-- .element: class="fragment" data-fragment-index="2" -->
- 4 standart Namespaces  <!-- .element: class="fragment" data-fragment-index="3" -->
    - default  <!-- .element: class="fragment" data-fragment-index="4" -->
    - kube-node-lease <!-- .element: class="fragment" data-fragment-index="5" -->
    - kube-public <!-- .element: class="fragment" data-fragment-index="6" -->
    -kube-system <!-- .element: class="fragment" data-fragment-index="7" -->

<aside class="notes">
  Namespaces sind ein ganz wichtiger punkt in Kubernetes

  separiert im Cluster verschiedene Anwendungen

  Gleiche Anwendung kann im Cluster in verschiedenen Namespaces mit gleichen Namen laufen
  
  default: Objekte welche keinem Anderen Namespace zugeordnet werden

  kuube-node-lease: hält objecte welche mit jedem node zusammenhängen

  erlaubt dem kubelet hearbeats an die control plane zu schicken

  kube-public: wenn anwendungen im kompletten cluster sichtbar sein sollen

  kube-system objekte, welche vom Kubernetes system erstellt wurden

</aside>

+++

<!-- .slide: style="text-align: left;"> -->
### DNS
- CoreDNS und KubeDNS <!-- .element: class="fragment" data-fragment-index="1" -->
- FQDN in cluster <!-- .element: class="fragment" data-fragment-index="2" -->
    - POD.NAMESPACE.pod.cluster.local <!-- .element: class="fragment" data-fragment-index="2" -->
    - SERVICE.NAMESPACE.svc.cluster.local <!-- .element: class="fragment" data-fragment-index="3" -->

<aside class="notes">
  CoreDNS und KubeDNS sind die beiden größten DNS services in Kubernetes

  CoreDNS neuer und mittlererweile standart seit kubernetes 1.12

  KubeDNS ist mit interner Namensauflösung ca 10% schneller

  CoreDNS mit externer Namensauflösung ca 3x besser

  CoreDNS Ressourcen schonender

  im cluster kann man auch eine FQDN auflösung machen

  entweder zum pod mit dem podnamen.namespacenamen.pod.cluster.local

  oder die bessere art, zum service mit servicename.namespace.svc.cluster.local

  man kann wenn man im gleichen namespace ist alles nach servicename weg lassen
</aside>

+++

<!-- .slide: style="text-align: left;"> -->
## OpenSource

<aside class="notes">
  Alle diese Komponenten sind opensource und theoretisch austauschbar, auch wenn einige so sehr in den kubernetes core eingebaut sind, dass diese nur mit sehr hohen aufwand getauscht werden können
</aside>


+++

<!-- .slide: style="text-align: left;"> -->
## Ausfallsicherheit

- Container Health Check  <!-- .element: class="fragment" data-fragment-index="1" -->
    - readyness <!-- .element: class="fragment" data-fragment-index="2" -->
    - liveness <!-- .element: class="fragment" data-fragment-index="3" -->
- Hostsystemausfall <!-- .element: class="fragment" data-fragment-index="4" -->
- Update <!-- .element: class="fragment" data-fragment-index="5" -->

<aside class="notes">
  Kubernetes hat den großen Vorteil, dass die deployten Anwendungen Ausfallsicher sind

  dies wird erzielt, indem man bei container

  readyness checks, also checks die prüfen ob der container gestartet ist

  und liveness checks, also checks die fortlaufend prüfen, ob der container noch läuft

  definieren kann.

  Sollte dies einmal nicht der fall sein, dann versucht das Kubelet den status wieder herzustellen

  Bei einem Node ausfall sorgt der scheduler, dass die pods auf einem anderen Node gestartet wird

  updates werden bei kubernetes normalerweise so gemacht, dass die replicas nach und nach ausgetauscht werden

  dadurch ist die anwendung niemals komplett heruntergefahren und durchgehend erreichbar
</aside>

+++

<!-- .slide: style="text-align: left;"> -->
### YAML
- Alle k8s Objekte werden in [YAML](https://learn.getgrav.org/17/advanced/yaml) definiert
    - Wird Manifest genannt
    - Jedes Objekt hat eigene YAML Spec  <!-- .element: class="fragment" data-fragment-index="1" -->
- API Server versteht diese Specs  <!-- .element: class="fragment" data-fragment-index="2" -->
    - kubectl apply -f spec.yaml  <!-- .element: class="fragment" data-fragment-index="3" -->
        - Schickt den Inhalt von spec.yaml an k8s API Server  <!-- .element: class="fragment" data-fragment-index="4" -->
        - API Server prüft Inhalt und Berechtigungen (Admission Control)  <!-- .element: class="fragment" data-fragment-index="5" -->
        - Wenn ok -> API Server sorgt für Anlage/Update  <!-- .element: class="fragment" data-fragment-index="6" -->
        - Wenn nicht ok -> Reject  <!-- .element: class="fragment" data-fragment-index="7" -->

+++

<!-- .slide: style="text-align: left;"> -->
### Manifest: Beispiele

```yaml
apiVersion: apps/v1
kind: StatefulSet
metadata:
  name: web
spec:
  selector:
    matchLabels:
      app: nginx # has to match .spec.template.metadata.labels
  serviceName: "nginx"
  replicas: 3 # by default is 1
  template:
    metadata:
      labels:
        app: nginx # has to match .spec.selector.matchLabels
    spec:
      terminationGracePeriodSeconds: 10
      containers:
      - name: nginx
        image: nginx-slim:0.8
        ports:
        - containerPort: 80
          name: web
        volumeMounts:
        - name: www
          mountPath: /usr/share/nginx/html
  volumeClaimTemplates:
  - metadata:
      name: www
    spec:
      accessModes: [ "ReadWriteOnce" ]
      storageClassName: "my-storage-class"
      resources:
        requests:
          storage: 1Gi
```

<aside class="notes">
  erklären was in der yaml steht

  spaces als seperator, wenn es eingerückt drunter steht dann wird es weitergegeben

  step für step
</aside>

+++

<!-- .slide: style="text-align: left;"> -->
```yaml
    metadata:
      labels:
        app: nginx # has to match .spec.selector.matchLabels
    spec:
      terminationGracePeriodSeconds: 10
      containers:
      - name: nginx
        image: nginx-slim:0.8
        ports:
        - containerPort: 80
          name: web
        volumeMounts:
        - name: www
          mountPath: /usr/share/nginx/html
```

<aside class="notes">
  containers

  kennt man schon von docker

  name wie `--name` bei docker run 

  image: ist wie bei docker das image

  ports leitet den port aus dem container in den cluster

  volumeMounts: mountet ein volume in ein container

  metadata.labels hier werden labels definiert, welche mit selectoren genutzt werden können

  in dem fall app aber kann auch alles andere sein

  spec: der status in welchem das objekt sein soll, in diesem fall wird ein container definiert
</aside>

+++

<!-- .slide: style="text-align: left;"> -->
```yaml
apiVersion: apps/v1
kind: StatefulSet
metadata:
  name: web
  namespace: testing
```

<aside class="notes">
  definition des statefulSets

  apiVersion: nicht wichtig. diverse ressourcen haben andere apiVersionen. Werden im laufe der zeit mehr kennenlernen

  kind: in diesem fall statefulset

  metadata: 
    name: eindeutiger identifier
    namespace: Namespace in welches es deployt werden soll
</aside>

+++

<!-- .slide: style="text-align: left;"> -->
```yaml
spec:
  selector:
    matchLabels:
      app: nginx # has to match .spec.template.metadata.labels
  replicas: 3 # by default is 1
```

<aside class="notes">
  spec: unterschiedlich bei jedem object definiert, welchen status das object bekommen soll in dem fall ein statefulset
  
  selector: auf welche objekte das object matchen soll

  in diesem fall soll es labels matchen, welche den namen app hat

  wie unser container von vorhin

  replicas, definiert wie viele gleiche pods deployt werden sollen
</aside>

+++

<!-- .slide: style="text-align: left;"> -->
```yaml
  volumeClaimTemplates:
  - metadata:
      name: www
    spec:
      accessModes: [ "ReadWriteOnce" ]
      storageClassName: "my-storage-class"
      resources:
        requests:
          storage: 1Gi
```
<aside class="notes">
  Volume Claimes

  damit die pods auch daten schreiben können gibt es volume claims

  diese definieren, auf welchem filesystem diese daten abgelegt werden sollen

  braucht im hintergrund kubernetes object PersistantVolume

  metadata name, definiert wieder den namen des volumeclaims

  spec: wieder wie dieses object aussehen soll

  accessMode: ReadWriteOnce, gibt noch Many, wird aber seltenst genutzt

  storageClassName name von persistat volume

  ressources. request.storage wie viel speicherplatz reserviert werden soll
</aside>

---

<!-- .slide: style="text-align: left;"> -->
## Einrichtung eurer Umgebung

```sh
w6p exercise k8s
```

---

<!-- .slide: style="text-align: left;"> -->
## kubectl
CLI Tool für das Management von k8s Clustern

Man kann kubectl über folgende Wege mitteilen, mit welchem Cluster es sich verbinden soll: <!-- .element: class="fragment" data-fragment-index="1" -->
- via Command-Line Argument '--kubeconfig path/to/kubeconfig' <!-- .element: class="fragment" data-fragment-index="2" -->
- via Umgebungsvariable KUBECONFIG <!-- .element: class="fragment" data-fragment-index="3" -->
- via Default Kubeconfig-Datei ~/.kube/config <!-- .element: class="fragment" data-fragment-index="4" -->

+++

<!-- .slide: style="text-align: left;"> -->
## Weitere Parameter
- Welcher Namespace?
    - Alle: '--all-namespaces' oder '-A'
    - X: '--namespace X' oder '-n X'
- Was will ich tun? <!-- .element: class="fragment" data-fragment-index="1" -->
    - Ressource anschauen <!-- .element: class="fragment" data-fragment-index="2" -->
        - 'get pod', 'get deploy', 'get secret' <!-- .element: class="fragment" data-fragment-index="2" -->
    - Ressource erstellen <!-- .element: class="fragment" data-fragment-index="3" -->
        - 'run', 'create deploy', 'create ns' <!-- .element: class="fragment" data-fragment-index="3" -->
    - Ressource löschen <!-- .element: class="fragment" data-fragment-index="4" -->
        - 'delete pod', 'delete deploy', 'delete sts' <!-- .element: class="fragment" data-fragment-index="4" -->

+++

<!-- .slide: style="text-align: left;"> -->
## Konkrete Beispiele
- Anlegen einer ConfigMap im NS 'webapp':

```sh
kubectl create -n webapp cm db-config \
    --from-literal=db-user=dba \
    --from-literal=db-password=OH-NO
```

+++

<!-- .slide: style="text-align: left;"> -->
- Erstelle Pod Manifest:

```sh
export do='--dry-run=client -o yaml'
kubectl run my-pod --image ubuntu:20.04 $do > pod.yaml
```

Dieses Manifest kann jetzt bequem angepasst/vervollständigt werden <!-- .element: class="fragment" data-fragment-index="1" -->

+++

<!-- .slide: style="text-align: left;"> -->
- Erstelle Deployment Manifest:

```sh
export do='--dry-run=client -o yaml'
kubectl create deploy --image alpine:3.15 my-deploy $do > dp.yaml
```

+++

<!-- .slide: style="text-align: left;"> -->
## Aufgabe 1

```sh
w6p exercise k8s -n1
```
Lösung nach 20m

---

<!-- .slide: style="text-align: left;"> -->
## k9s

Terminal UI zum Management eines k8s Clusters
- Start
```sh
k9s [--kubeconfig path/to/kubeconfig] [-n NAMESPACE]
```
- Stop: Ctrl^c <!-- .element: class="fragment" data-fragment-index="1" -->
- Hilfe: '?' <!-- .element: class="fragment" data-fragment-index="2" -->
- Navigation mit Pfeiltasten <!-- .element: class="fragment" data-fragment-index="3" -->
- Tiefer reinspringen mit ENTER <!-- .element: class="fragment" data-fragment-index="4" -->
    - z.B. von selektiertem Deployment zu allen darüber verwalteten Pods <!-- .element: class="fragment" data-fragment-index="5" -->
    - Zurück mit ESC <!-- .element: class="fragment" data-fragment-index="6" -->

+++

- Ressourcen-Navigation
    - ':'
    - Gefolgt von der gewünschten Ressource
        - Context: 'context' oder 'ctx' <!-- .element: class="fragment" data-fragment-index="1" -->
        - Namespace: 'namespace' oder 'ns' <!-- .element: class="fragment" data-fragment-index="2" -->
        - Pod: 'pod' oder 'po' <!-- .element: class="fragment" data-fragment-index="3" -->
        - Deployment: 'deployment', 'deploy' oder 'dp' <!-- .element: class="fragment" data-fragment-index="4" -->
        - PodSecurityPolicy: 'psp' <!-- .element: class="fragment" data-fragment-index="5" -->
        - ... <!-- .element: class="fragment" data-fragment-index="5" -->
    - TAB sensitiv bei Auto-Vorschlägen <!-- .element: class="fragment" data-fragment-index="6" -->
    - Mit ENTER bestätigen <!-- .element: class="fragment" data-fragment-index="7" -->

+++

<!-- .slide: style="text-align: left;"> -->
- Context-Switch
    - ':ctx' ENTER Navigiere-zu-Context ENTER
    - Verbindet sich mit dem ausgewählten Cluster <!-- .element: class="fragment" data-fragment-index="1" -->
    - In der Folge werden nur noch Objekte dieses Clusters angezeigt <!-- .element: class="fragment" data-fragment-index="2" -->
- Namespace-Switch <!-- .element: class="fragment" data-fragment-index="3" -->
    - ':ns' ENTER Navigiere-zu-Namespace ENTER <!-- .element: class="fragment" data-fragment-index="3" -->
    - In der Folge wird die Sichtbarkeit <!-- .element: class="fragment" data-fragment-index="4" -->
        - erweitert auf alle Namespaces, wenn NS='all' <!-- .element: class="fragment" data-fragment-index="4" -->
        - eingeschränkt auf gewählten Namespace, sonst <!-- .element: class="fragment" data-fragment-index="4" -->

+++

<!-- .slide: style="text-align: left;"> -->
- Wenn eine Ressource X selektiert ist
    - Describe: 'd' (k describe X) <!-- .element: class="fragment" data-fragment-index="1" -->
    - Zeige YAML: 'y' (k get X -o yaml) <!-- .element: class="fragment" data-fragment-index="2" -->
    - Delete: 'Ctrl^d' (k delete X) <!-- .element: class="fragment" data-fragment-index="3" -->
    - Edit: 'e' (k edit X) <!-- .element: class="fragment" data-fragment-index="4" -->
        - vi oder vim <!-- .element: class="fragment" data-fragment-index="4" -->

+++

<!-- .slide: style="text-align: left;"> -->
- Je nachdem welche Resource ausgewählt ist, weitere Optionen möglich
    - Für Pods z.B: <!-- .element: class="fragment" data-fragment-index="1" -->
         - Logs: 'l' (k logs X) <!-- .element: class="fragment" data-fragment-index="1" -->
         - Shell: 's' (k exec -ti X sh) <!-- .element: class="fragment" data-fragment-index="1" -->
         - Port Forward: 'Shift^f' (k port-forward X 80) <!-- .element: class="fragment" data-fragment-index="1" -->
    - Für Deployments z.B: <!-- .element: class="fragment" data-fragment-index="2" -->
         - Logs: 'l' <!-- .element: class="fragment" data-fragment-index="2" -->
         - Scale: 's'  <!-- .element: class="fragment" data-fragment-index="2" -->
         - Restart: 'r'  <!-- .element: class="fragment" data-fragment-index="2" -->

+++

<!-- .slide: style="text-align: left;"> -->
## Aufgabe 2

```sh
w6p exercise k8s -n2
```
Lösung nach 20m

Note:
Ende Tag 2

---

# TAG 3

+++

# Agenda
- Recap
- Weitere Kubernetes Objekte
- Helm

---

<!-- .slide: style="text-align: left;"> -->
# Recap

+++

<!-- .slide: style="text-align: left;"> -->
## Docker
- Container Runtime Engine
- Bau von Images
- Starten von Images
- Verwalten von Images

+++

<!-- .slide: style="text-align: left;"> -->
## Kubernetes
- Container Orchestrierungstool
    - Verwalten von Pods
    - Starten stoppen und Überwachen
    - Self-Healing 
    - Dynamische Skalierung

+++

<!-- .slide: style="text-align: left;"> -->
### Objekte

+++

<!-- .slide: style="text-align: left;"> -->
- Pod
    - kleinste deploybare Einheit
    - beinhaltet 1 bis N Container
    - eigener Netzbereich und IP
- ReplicaSet <!-- .element: class="fragment" data-fragment-index="1" -->
    - Stellt sicher, dass zu jeder Zeit genau N Pods laufen <!-- .element: class="fragment" data-fragment-index="2" -->
    - Matching über Labels <!-- .element: class="fragment" data-fragment-index="3" -->

+++

<!-- .slide: style="text-align: left;"> -->
- Deployment
    - Managed ein ReplicaSet
    - Bietet Versionierung und zero downtime Rollouts
- DeamonSet <!-- .element: class="fragment" data-fragment-index="1" -->
    - Spec wie Deployment nur ohne Replica Count <!-- .element: class="fragment" data-fragment-index="1" -->
    - Managed damit kein ReplicaSet <!-- .element: class="fragment" data-fragment-index="1" -->
    - Stattdessen je ein Replica pro Node <!-- .element: class="fragment" data-fragment-index="1" -->

+++

<!-- .slide: style="text-align: left;"> -->
- Service
    - Loadbalancer für Pods
    - Auch hier Matching via Labels
    - Typen <!-- .element: class="fragment" data-fragment-index="1" -->
        - None (headless) <!-- .element: class="fragment" data-fragment-index="1" -->
        - ClusterIP (Default) <!-- .element: class="fragment" data-fragment-index="2" -->
        - NodePort <!-- .element: class="fragment" data-fragment-index="3" -->
        - Loadbalancer <!-- .element: class="fragment" data-fragment-index="4" -->

Note:
Service - headless (erstellt für jeden Pod einen DNS entry innerhalb des Clusters (coredns), kein externer Zugriff möglich)

Service - ClusterIP (routet über die clusterinternen Pod IPs, kein externer Zugiff möglich)

Service - NodePort (öffnet auf jedem Node denselben Port, über den von außen der Service erreicht werden kann)

Service - Loadbalancer (exosed den Service ins Internet, bedarf eines Loadbalancers der den Traffic an der Service weiterleitet)

+++

<!-- .slide: style="text-align: left;"> -->
- StatefulSet
    - Spec ähnlich zu Deployment
    - Geordnetes Starten (einer nach dem anderen)
    - Geordnetes Stoppen in umgekehrter Reihenfolge
- ConfigMap <!-- .element: class="fragment" data-fragment-index="1" -->
    - Plain-Text Key-Value Store <!-- .element: class="fragment" data-fragment-index="1" -->
    - Kann in Pods, Deployments, STSs und DSs gemounted werden <!-- .element: class="fragment" data-fragment-index="2" -->
- Secret <!-- .element: class="fragment" data-fragment-index="3" -->
    - base64 encoded Data Store <!-- .element: class="fragment" data-fragment-index="3" -->
    - Kann in Pods, Deployments, STSs und DSs gemounted werden <!-- .element: class="fragment" data-fragment-index="4" -->

+++

<!-- .slide: style="text-align: left;"> -->
### Kubernetes Tools
- kubectl <!-- .element: class="fragment" data-fragment-index="1" -->
    - CLI zur Interaktion mit k8s Clustern <!-- .element: class="fragment" data-fragment-index="1" -->
- krew <!-- .element: class="fragment" data-fragment-index="2" -->
    - kubectl Plugin Manager <!-- .element: class="fragment" data-fragment-index="2" -->
- k9s <!-- .element: class="fragment" data-fragment-index="3" -->
    - Terminal UI zur Interaktion mit k8s Clustern <!-- .element: class="fragment" data-fragment-index="3" -->
- kind (Kubernetes in Docker) <!-- .element: class="fragment" data-fragment-index="4" -->
    - Single-Node k8s Cluster in Docker Container <!-- .element: class="fragment" data-fragment-index="4" -->

+++

<!-- .slide: style="text-align: left;"> -->
- helm
    - Paket-Manager für Kubernetes
    - vgl. mit apt für Ubuntu oder apk für Alpine
- Lens <!-- .element: class="fragment" data-fragment-index="1" -->
    - Graphical UI zur Interaktion mit k8s Clustern <!-- .element: class="fragment" data-fragment-index="1" -->
    - Nicht Teil des Workshops <!-- .element: class="fragment" data-fragment-index="1" -->

---

<!-- .slide: style="text-align: left;"> -->
# Weitere Objekttypen in K8s

+++

<!-- .slide: style="text-align: left;"> -->
## Errinnerung Pod
- Kleinste Deploybare einheit
- Kann per yaml datei erstellt/modifiziert werden

+++

<!-- .slide: style="text-align: left;"> -->
### Aufgabe:
- Kaputte pod yaml https://github.com/x-cellent/k8s-workshop/blob/4d12a2b505babef8f0d06875f397aaf9d0147973/exercises/k8s/ex3%20-%20fix%20broken%20Pod/exercise.md
- Reparieren und in Namespace der Wahl deployen

+++

```yaml
apiversion: v1
Kind: pod
metadata:
  labels:
    app: frontend
  name: web
spec:
containers:
  name: web
    image: nginx
    tag: latest
    ports:
  - containerPort: 80
    resources:
      requests:
        cpu: "1.0"
        memory:"1G"
      limits:
       cpu: "1.0"
        memory: 1G
```

+++

<!-- .slide: style="text-align: left;"> -->
### Lösung
```yaml
apiVersion: v1 #Typo in apiVersion, V von Version muss groß sein
kind: Pod #Typo, kind muss klein sein
metadata:
  labels:
    app: frontend
  name: web
  namespace: ex1 # Optional, kann auch kubectl auch via "-n ex1" mitgegeben werden
spec:
  containers: #ab hier muss alles eingeruckt sein
  - name: web #listen in yaml werden beim ersten punkt mit `-`angegeben
    image: nginx:latest #Image und Tag definiert man in einer zeile mit `:` dazwischen
    ports:
    - containerPort: 80 #hier auch falsch eingerückt
    resources:
      requests:
        cpu: "1.0"
        memory: "1G" # zwischen memory und den 1G muss ein Leerzeichen sein
      limits:
        cpu: "1.0"
        memory: "1G" #1G muss in anführungszeichen sein
```

+++

<!-- .slide: style="text-align: left;"> -->
## Service
- Loadbalancer für Pods
- Matching via Labels

+++

![image](https://miro.medium.com/max/1400/0*X1VC6PMEMbxloLmh.png)

<aside class="notes">
Der Service leitet anfragen welche in die Nodes kommt an die Pods weiter
</aside>

+++

<!-- .slide: style="text-align: left;"> -->
#### Aufgabe
- erstelle ein Deployment (Objekt) des Pods aus vergangener aufgabe
    - yaml des pods hier zu finden: https://github.com/x-cellent/k8s-workshop/blob/main/exercises/k8s/ex3%20-%20fix%20broken%20Pod/pod.yaml
- erstelle anschließend ein Service (Objekt) um die Pods zu Loadbalancen
    - Deployment Lösung hier zu finden: https://github.com/x-cellent/k8s-workshop/blob/main/exercises/k8s/ex4%20-%20create%20Service/deplyoment.yaml

+++

#### Lösung
- erst deployment.yaml erstellen und in gewünschten Namespace deployen
    - https://github.com/x-cellent/k8s-workshop/blob/main/exercises/k8s/ex4%20-%20create%20Service/deplyoment.yaml
    - kubectl apply -d deployment.yaml -n web
- anschließend eine service.yaml erstellen 

+++

<!-- .slide: style="text-align: left;"> -->
```yaml
apiVersion: v1
kind: Service
metadata:
  name: web
spec:
  selector:
    app: frontend # selector hier muss dem label des deployments ensprechen
  ports:
    - protocol: TCP
      port: 80
      targetPort: 8081
```

+++

<!-- .slide: style="text-align: left;"> -->
- Diese yaml in den gewünsten Namespace deployen
    - k apply -f service.yaml -n web

+++

<!-- .slide: style="text-align: left;"> -->
## Port-Forwarding
- Zugriff in Container erlangen
- für debugging
- k9s kann diese Verwalten

+++

<!-- .slide: style="text-align: left;"> -->
### Aufgabe
- Deployment und service aus letzter Aufgabe muss im selben Namespace deployt sein
    - kubectl create ns web
    - kubectl apply -f https://github.com/x-cellent/k8s-workshop/blob/main/exercises/k8s/ex4%20-%20create%20Service/deplyoment.yaml -n web
    - kubectl apply -f https://github.com/x-cellent/k8s-workshop/blob/main/exercises/k8s/ex4%20-%20create%20Service/service.yaml -n web
- Anschließend bitte ein Port-Forwarding in einen Pod machen
    - Welche Wege gibt es?


+++

<!-- .slide: style="text-align: left;"> -->
### Lösung
- es gibt den weg mit kubectl
    - kubectl port-forward -n web service/web 8081:8081
    - kubectl port-forward -n web deployment/web 8081:80
    - kubectl port-forward -n web pod/web-6779b45f74-bvc7p 8081:80
    - kubectl port-forward -n web pod/web-6779b45f74-bvc7p :80 
        - hier bestimmt kubectl selber den local port
- mit k9s ist es auch möglich


+++

<!-- .slide: style="text-align: left;"> -->
## DaemonSet
- Jeder Node bekommt ein Replica
    - Log-Shipper
    - Monitoring Agent

<aside class="notes">
  daemonSets starten eine Replica auf jeder Node

  kann nicht passieren, dass bei einem Node ausfall die pods erst auf einer neuen node deployt wird

  oft verwendet für log collector, da diese die logs von allen pods auf allen nodes braucht

  ebenso bei monitoring collectorn, da diese das monitoring von allen nodes braucht
</aside>

+++

<!-- .slide: style="text-align: left;"> -->
### Aufgabe 
- Deploye ein DaemonSet mit einem nginx Pod in ein Namespace deiner Wahl
- Scale das DaemonSet auf 3 Pods
    - ist dies Möglich?
    - Warum? Warum nicht?

+++


### Lösung
```yaml
apiVersion: apps/v1
kind: DaemonSet
metadata:
  name: web-ds
  labels:
    app: web-ds
spec:
  selector:
    matchLabels:
      name: web-ds
  template:
    metadata:
      labels:
        name: web-ds
    spec:
      containers:
      - name: web-ds
        image: nginx:latest
        resources:
          limits:
            memory: 200Mi
          requests:
            cpu: 100m
            memory: 200Mi
```

+++

- scaling ist nicht möglich, da daemonSets mit den Nodes scalen

+++

<!-- .slide: style="text-align: left;"> -->
## StatefulSet
- Persistente Pods
- Geordnetes Update/Shutdown

<aside class="notes">
  StatefulSets sind sinnvoll, wenn man erzielen möchte, dass eine anwendung ihren status nicht verliert

  z.B Datenbanken sind klassische Anwendungen welche man in diesem zustand haben möchte.
</aside>

+++

<!-- .slide: style="text-align: left;"> -->
## Job
- Einmalige Ausführung eines Commands in einem Pod
    - Datenbank Backup

<aside class="notes">
  jobs sind praktisch um einzelne kommandos auszuführen

  z.B prüfen ob ein service im cluster erreichbar ist, datenbank backups zu erstellen
</aside>

+++

<!-- .slide: style="text-align: left;"> -->
### Aufgabe 
- Erstelle ein Job welcher einmalig die Zahl Pi auf 5000 Stellen genau berechnet.
- gebe Pi aus

+++

<!-- .slide: style="text-align: left;"> -->
#### Lösung
- von kubernetes doku das Manifest übernehmen und anpassen

+++

<!-- .slide: style="text-align: left;"> -->
```yaml
apiVersion: batch/v1
kind: Job
metadata:
  name: pi
spec:
  template:
    metadata:
      labels:
        job: pi
    spec:
      containers:
      - name: pi
        image: perl
        command: ["perl",  "-Mbignum=bpi", "-wle", "print bpi(5000)"]
      restartPolicy: Never
  backoffLimit: 4
```

+++

<!-- .slide: style="text-align: left;"> -->
- kubectl get logs -n ex7 pi-

+++

<!-- .slide: style="text-align: left;"> -->
## CronJobs
- Mischung aus klassischen CronJobs und Jobs
- Regelmäßige Ausführung eines Jobs
    - Datenbank Backups

<aside class="notes">
  wie klassische Linux Cronjobs 

  Regelmäßige außführung von Jobs

  man kann auch einmalig die Jobs eines Cronjobs außführen pratkisch für debugging
</aside>

+++

<!-- .slide: style="text-align: left;"> -->
### Aufgabe 
- erstelle einen Cronjob welcher minütlich das datum und deinen Namen ausgibt
- dieser Cronjob soll 5 erfolgreiche und 8 fehlgeschlagene versuche behalten
- teste diesen cronjob ohne eine minute zu warten

+++

<!-- .slide: style="text-align: left;"> -->
#### Lösung
- aus kubernetes Doku Manifest kopieren und anpassen

+++

<!-- .slide: style="text-align: left;"> -->
```yaml
apiVersion: batch/v1
kind: CronJob
metadata:
  name: hello
spec:
  successfulJobsHistoryLimit: 5
  failedJobsHistoryLimit: 8
  schedule: "*/1 * * * *"
  jobTemplate:
    spec:
      template:
        metadata:
          labels:
            cronjob: hello
        spec:
          containers:
          - name: hello
            image: busybox:1.28
            imagePullPolicy: IfNotPresent
            command:
            - /bin/sh
            - -c
            - date; echo Pascal
          restartPolicy: OnFailure
```

+++

<!-- .slide: style="text-align: left;"> -->
- erstellen eines einmaligen runs
```sh
 kubectl create job -n ex8 --from=cronjob/hello hello-test
```
- Output wieder sichtbar mit 
```sh
kubectl logs -n ex8 hello-
```

+++

<!-- .slide: style="text-align: left;"> -->
## ConfigMaps
- Speicherung von nicht vertraulichen daten
- Einbindung in Pods als
    - Umgebungsvariable
    - command-line argument
    - Datei (Volume)
- Kein Reload der Pods bei Änderung

<aside class="notes">
  in configmaps sollen nur nicht vertrauliche daten gespeichert werden

  es gibt mehrere wege diese in die container einzubinden

  pods reloaden nicht automatisch wenn configmaps geupdated wurden
  </aside>

+++

<!-- .slide: style="text-align: left;"> -->
### Aufgabe
- dieses deployment möchte eine ConfigMap einbinden
    - [Deployment.yaml](https://github.com/x-cellent/k8s-workshop/blob/main/exercises/k8s/ex9%20-%20ConfigMap%20-%20Deployment/deployment.yaml)
- diese ConfigMap
    - [configmap.yaml](https://github.com/x-cellent/k8s-workshop/blob/main/exercises/k8s/ex9%20-%20ConfigMap%20-%20Deployment/configmap.yaml)
- ändere die WorkerConnection und deploye die beiden Ressourcen

+++

<!-- .slide: style="text-align: left;"> -->
#### Lösung
- WorkerConnection in Zeile 13 Updaten, anschließend zurst die Configmap deployen
```sh
kubectl apply -f configmap.yaml -n ex9
```
- anschließend das deployment deplyoen
```sh
kubectl apply -f deployment.yaml -n ex9
```
- Wichtig! Beides in den gleichen Namespace

+++

<!-- .slide: style="text-align: left;"> -->
## Secret
- Speicherung vertraulicher Daten
- Unverschlüsselt in etcd DB
- Bessere Seperierung mittels Rollen
   - User darf Configmaps sehen aber keine Secrets

<aside class="notes">
  secrets gibt es um vertrauliche daten zu speichern

  standartmäßig liegen diese daten aber unverschlüsselt im etcd

  einbindung ähnlich wie bei configmaps
</aside>

+++

## PersistantVolume (PV)
- sehr viele Volume Typen
    - Lightbits, local und s3 bei der FI-TS
- Speichert Infos über Volumen und Storage
- überverzeichnis muss bereits erstellt sein


+++

- ReadWriteOnce oder ReadWriteMany
    - Once, nur ein Node darf auf das Volume schreiben
    - Many, mehrere dürfen
- ReadOnlyMany
    - mehere Nodes können das Volume ReadOnly Mounten

+++

### Aufgabe
- erstelle ein local PV mit 10 GB Capacity
- erstelle das Verzeichnis auf der Node
- dieser soll ReadWriteOnce sein
- dieser muss einen eindeutigen Namen haben

+++

#### Lösung
```yaml
apiVersion: v1
kind: PersistentVolume
metadata:
  name: example-pv
  labels:
    storage: local
spec:
  storageClassName: standard
  capacity:
    storage: 10Gi
  volumeMode: Filesystem
  accessModes:
  - ReadWriteOnce
  persistentVolumeReclaimPolicy: Delete
  local:
    path: /mnt/disks/ssd1
  nodeAffinity:
    required:
      nodeSelectorTerms:
      - matchExpressions:
        - key: kubernetes.io/hostname
          operator: In
          values:
          - k8s-workshop-cluster-control-plane
```

+++

## PersistantVolumeClaim (PVC)
- Reserviert Ressourcen eines PV`s
- wird anschließend ins deployment eingebaut
- Verknüpfung PV und PVC mit Selector labels oder direkt mit namen
    - bei local kein dynamisches (selector) mapping möglich
- Verknüpfung ist eine 1 zu 1 Verknüpfung
    - keine 2 PVC an einem PV

+++

### Aufgabe
- erstelle ein PVC
- erstelle ein postgresql statefulset
    - Tipp: Configmap und Secret müssen auch erstellt sein
      um env Variablen in den Container zu übergeben
- welches das PVC einbindet
- lasse die daten welche in der DB sind anzeigen

+++

#### Lösung
- Der PV der letzten aufgabe muss erstellt sein
```yaml
kind: PersistentVolumeClaim
apiVersion: v1
metadata:
  name: postgres-pv-claim
  labels:
    app: postgres
spec:
  storageClassName: standard
  capacity:
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 5Gi
  volumeName: example-pv
```

+++

- Configmap
```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: postgres-configuration
  labels:
    app: postgres
data:
  POSTGRES_DB: topdb
  POSTGRES_USER: user23
```
- Secret
```yaml
apiVersion: v1
kind: Secret
metadata:
  name: postgres-secret
type: Opaque
data:
  POSTGRES_PASSWORD: sicherespasswort
```

+++

- Statefulset
```yaml
apiVersion: apps/v1
kind: StatefulSet
metadata:
  name: postgres-statefulset
  labels:
    app: postgres
spec:
  serviceName: "postgres"
  replicas: 1
  selector:
    matchLabels:
      app: postgres
  template:
    metadata:
      labels:
        app: postgres
    spec:
      containers:
      - name: postgres
        image: postgres:12
        envFrom:
        - configMapRef:
            name: postgres-configuration
        - secretRef:
            name: postgres-secret
        ports:
        - containerPort: 5432
          name: postgresdb
        volumeMounts:
        - name: pv-data
          mountPath: /var/lib/postgresql/data
      volumes:
      - name: pv-data
        persistentVolumeClaim:
          claimName: postgres-pv-claim
```

+++

- Daten anzeigen lassen
```sh
kubectl exec -n postgresql -it postges-statefulset-0 -- /bin/bash
psql -U user23 topdb
```

---

<!-- .slide: style="text-align: left;"> -->
# Helm
- Package Manager für Kubernetes
- gegliedert in sogenannten Charts
- Große Softwarehersteller schreiben eigene Helm Charts
    - z.B. Gitlab
- praktisch um eine anwendung mit wenigen änderungen in verschiedenen umgebungen zu deployen
    - test/staging/production
- helm charts sind in sogenannten Repos gespeichert
    - chart ersteller meistens eigene Repo
    - nutzung ähnlich wie bei apt in ubuntu
        - adden, updaten installieren

+++

<!-- .slide: style="text-align: left;"> -->
![image](https://developer.ibm.com/developer/default/blogs/kubernetes-helm-3/images/helm3-arch.png)

+++

<!-- .slide: style="text-align: left;"> -->
## Aufbau eines Helm Charts
```sh
schulung
├── charts
├── Chart.yaml
├── templates
│   ├── deployment.yaml
│   ├── _helpers.tpl
│   ├── hpa.yaml
│   ├── ingress.yaml
│   ├── NOTES.txt
│   ├── serviceaccount.yaml
│   ├── service.yaml
│   └── tests
│       └── test-connection.yaml
└── values.yaml
```

+++

<!-- .slide: style="text-align: left;"> -->
## Aufbau eines Helm Charts
- das meiste spielt sich im templates ordner ab
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: {{ include "schulung.fullname" . }}
  labels:
    {{- include "schulung.labels" . | nindent 4 }}
spec:
  {{- if not .Values.autoscaling.enabled }}
  replicas: {{ .Values.replicaCount }}
  {{- end }}
  selector:
    matchLabels:
      {{- include "schulung.selectorLabels" . | nindent 6 }}
  template:
    metadata:
      {{- with .Values.podAnnotations }}
      annotations:
        {{- toYaml . | nindent 8 }}
      {{- end }}
      labels:
        {{- include "schulung.selectorLabels" . | nindent 8 }}
    spec:
      {{- with .Values.imagePullSecrets }}
      imagePullSecrets:
        {{- toYaml . | nindent 8 }}
      {{- end }}
      serviceAccountName: {{ include "schulung.serviceAccountName" . }}
      securityContext:
        {{- toYaml .Values.podSecurityContext | nindent 8 }}
      containers:
        - name: {{ .Chart.Name }}
          securityContext:
            {{- toYaml .Values.securityContext | nindent 12 }}
          image: "{{ .Values.image.repository }}:{{ .Values.image.tag | default .Chart.AppVersion }}"
          imagePullPolicy: {{ .Values.image.pullPolicy }}
          ports:
            - name: http
              containerPort: 80
              protocol: TCP
          livenessProbe:
            httpGet:
              path: /
              port: http
          readinessProbe:
            httpGet:
              path: /
              port: http
          resources:
            {{- toYaml .Values.resources | nindent 12 }}
      {{- with .Values.nodeSelector }}
      nodeSelector:
        {{- toYaml . | nindent 8 }}
      {{- end }}
      {{- with .Values.affinity }}
      affinity:
        {{- toYaml . | nindent 8 }}
      {{- end }}
      {{- with .Values.tolerations }}
      tolerations:
        {{- toYaml . | nindent 8 }}
      {{- end }}
```

+++

<!-- .slide: style="text-align: left;"> -->
## Aufbau eines Helm Charts
- wie so oft im yaml format
- das meiste bis alles templates
- anpassungen in der values.yaml

+++

<!-- .slide: style="text-align: left;"> -->
```yaml
# Default values for schulung.
# This is a YAML-formatted file.
# Declare variables to be passed into your templates.

replicaCount: 1

image:
  repository: nginx
  pullPolicy: IfNotPresent
  # Overrides the image tag whose default is the chart appVersion.
  tag: ""

imagePullSecrets: []
nameOverride: ""
fullnameOverride: ""

serviceAccount:
  create: true
  annotations: {}
  name: ""

podAnnotations: {}

podSecurityContext: {}

securityContext: {}

service:
  type: ClusterIP
  port: 80

ingress:
  enabled: false
  className: ""
  annotations: {}
  hosts:
    - host: chart-example.local
      paths:
        - path: /
          pathType: ImplementationSpecific
  tls: []

resources: {}
  # We usually recommend not to specify default resources and to leave this as a conscious
  # choice for the user. This also increases chances charts run on environments with little
  # resources, such as Minikube. If you do want to specify resources, uncomment the following
  # lines, adjust them as necessary, and remove the curly braces after 'resources:'.
  # limits:
  #   cpu: 100m
  #   memory: 128Mi
  # requests:
  #   cpu: 100m
  #   memory: 128Mi

autoscaling:
  enabled: false
  minReplicas: 1
  maxReplicas: 100
  targetCPUUtilizationPercentage: 80
  # targetMemoryUtilizationPercentage: 80

nodeSelector: {}

tolerations: []

affinity: {}
```

+++

<!-- .slide: style="text-align: left;"> -->
- Beispiel an Container part des Deployment template

+++

```yaml
      containers:
        - name: {{ .Chart.Name }}
          securityContext:
            {{- toYaml .Values.securityContext | nindent 12 }}
          image: "{{ .Values.image.repository }}:{{ .Values.image.tag | default .Chart.AppVersion }}"
          imagePullPolicy: {{ .Values.image.pullPolicy }}
          ports:
            - name: http
              containerPort: 80
              protocol: TCP
          livenessProbe:
            httpGet:
              path: /
              port: http
          readinessProbe:
            httpGet:
              path: /
              port: http
          resources:
            {{- toYaml .Values.resources | nindent 12 }}
```

+++

```yaml
image:
  repository: nginx
  pullPolicy: IfNotPresent
  # Overrides the image tag whose default is the chart appVersion.
  tag: ""

resources: {}
  # We usually recommend not to specify default resources and to leave this as a conscious
  # choice for the user. This also increases chances charts run on environments with little
  # resources, such as Minikube. If you do want to specify resources, uncomment the following
  # lines, adjust them as necessary, and remove the curly braces after 'resources:'.
  # limits:
  #   cpu: 100m
  #   memory: 128Mi
  # requests:
  #   cpu: 100m
  #   memory: 128Mi
```

+++

## Helm Commands
- helm install
    - installiert ein helm chart
    - mit -n namespace angebbar
    - mit --dry-run --debug kann man überprüfen ob das deployment klappen sollte
    - mit --version versionspinning
    - Syntax `helm install -n NAMESPACE RELEASE_NAME PFAD_ZUM_HELM_CHART

+++

- helm upgrade
    - upgraden eines helm charts auf neue revision
    - --install wichtige flag, macht, dass chart installiert wird wenns nicht da ist
    - mit --version versionspinning
- helm create
    - erstellen eines helm charts
    - erstellt die grundlegende ordner struktur

+++

- helm uninstall
    - deinstalliert ein chart, löscht alle ressourcen
- helm rollback
    - zurückspielen auf alte version des helm charts
- helm list
    - zeigt installierte helm charts
    - entweder mit -A für alle Namespaces oder -n mit Namespace angabe
- helm lint
    - überprüfung ob helm chart template keine fehler hat

+++

- helm repo
    - add 
        - hinzufügen eines repos
        - z.B. helm repo add bitnami
    - update
        - herunterladen welche charts in repos sind
        - z.B. in bitnami gibt es ein postgresql chart

+++

## Aufgabe:
1. erstelle ein Helm Chart für ein nginx deployment mit service
1. deploye dies in ein Namespace deiner wahl

+++

### Lösung
1. erst das chart erstellen
```sh
helm create NAME
helm create nginx-deployment
```
1. dann installieren
```sh
helm install -n NAMESPACE RELEASE_NAME PFAD_ZUM_HELM_CHART
helm install -n helm-namespace nginx-deployment ./nginx-deployment
```

+++

## Aufgabe:
1. passe die replicas mit helm an
1. Verifiziere, dass mehr pods laufen

+++

### Lösung

1. dann die values yaml anpassen und upgrade
```sh
helm upgrade -n NAMESPACE RELEASE_NAME PFAD_ZUM_HELM_CHART
helm upgrade -n helm-namespace nginx-deployment ./nginx-deployment
```
1. mit kubectl oder k9s anzeigen, dass die angegebenen Pods da sind
```sh
kubectl get pods -n NAMESPACE
kubectl get pods -n helm-namespace
```

+++

## Aufgabe:
1. mache ein Rollback auf eine alte Helm version

+++

### Lösung
1. mit helm rollback auf alte revision gehen
```sh
helm rollback -n NAMESPACE RELEASE_NAME REVISION
helm rollback -n helm-namespace nginx-deployment 1
```

+++

### Übersicht Helm
- Ist ein Packetmanager
- arbeitet mit Templates
- eine zentrale datei (values.yaml) um komplexe anwendungen zu deployen
- wird in Repos verwaltet

---

<!-- .slide: style="text-align: left;"> -->
## Container Runtime Interface
- API für Container Verwaltung (Starten/Stoppen)
- Wird von Kubernetes unterstützt
    - Konkrete Implemetierung damit austauschbar
- [Container Runtimes](https://kubernetes.io/docs/setup/production-environment/container-runtimes/)
    - containerd
    - CRI-O
    - Docker Engine

+++

<!-- .slide: style="text-align: left;" class="stretch"> -->
![image](images/docker-CRI-O-containerd-runc.png)

Notes:
containerd -  Linux-Daemon; Pulled Images aus Registry, verwaltet Speicher und Netzwerke, started/stoppt Containern via runc

runc – Low-Level-Container-Runtime; verwendet libcontainer - native Go-basierte Implementierung zum Starten und Stoppen von Containern

---

<!-- .slide: style="text-align: left;"> -->
## RBAC

Role Based Access Control
- Authentifizierung (Wer bin ich?) <!-- .element: class="fragment" data-fragment-index="1" -->
    - Analogie Ausreise: Perso <!-- .element: class="fragment" data-fragment-index="1" -->
- Autorisierung (Was darf ich?) <!-- .element: class="fragment" data-fragment-index="2" -->
    - Analogie Ausreise: Visa <!-- .element: class="fragment" data-fragment-index="2" -->
- Admission Control <!-- .element: class="fragment" data-fragment-index="3" -->
    - Stellt Authentifizierung und Autorisierung sicher <!-- .element: class="fragment" data-fragment-index="3" -->
    - Analogie Ausreise: <!-- .element: class="fragment" data-fragment-index="4" -->
        - Prüft Perso und Visa <!-- .element: class="fragment" data-fragment-index="4" -->
        - Zoll: Darf ich mein Gepäck einführen? <!-- .element: class="fragment" data-fragment-index="4" -->

+++

<!-- .slide: style="text-align: left;"> -->
## Admission Controller
-
Allow/Deny/Change APi-Requests
- Basiert auf Regeln und Policies


---

<!-- .slide: style="text-align: left;"> -->
# Literatur
- [Kubernetes Up & Running](https://www.amazon.de/Kubernetes-Up-Running-Brendan-Burns/dp/1492046531)
- [Kubernetes Best Practices](https://www.amazon.de/Kubernetes-Best-Practices-Blueprints-Applications/dp/1492056472)
- [Online](https://kubernetes.io/docs)

---

# Fragen
- Hab ihr noch Fragen an uns?
