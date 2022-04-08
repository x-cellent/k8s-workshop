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
```

=> bin/w6p <!-- .element: class="fragment" data-fragment-index="1" -->

+++

<!-- .slide: style="text-align: left;"> -->
## Binary

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
bin/w6p exercise docker -n 1
```
Zeit: ca. 5 min

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
bin/w6p exercise docker -n2
```
Zeit: ca. 20 min

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
bin/w6p exercise docker -n3
```
Zeit: ca. 15 min

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
bin/w6p exercise docker -n4
```
Zeit: ca. 5 min

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
bin/w6p exercise docker -n5
```
Zeit: ca. 15 min

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
bin/w6p exercise docker -n6
```
Zeit: ca. 15 min

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
1. Einrichtung euerer Umgebung
1. Basisobjekte Kubernetes mit Übungen

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

- helm
    - Paket-Manager für Kubernetes
    - vgl. mit apt für Ubuntu oder apk für Alpine
- Lens <!-- .element: class="fragment" data-fragment-index="1" -->
    - Graphical UI zur Interaktion mit k8s Clustern <!-- .element: class="fragment" data-fragment-index="1" -->
    - Nicht Teil des Workshops <!-- .element: class="fragment" data-fragment-index="1" -->

---

<!-- .slide: style="text-align: left;"> -->
### Kubernetes

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

### etcd
- entwickelt von CoreOS
- key-value Database
- kann nicht getauscht werden
- speichert stand von cluster
- Consistency notwending

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
- Ansprechpunkt des Users
- Validation der Daten
- bekanntester ist kube-apiserver
- horizontale skalierbarkeit

<aside class="notes">
  immer wenn ihr im Cluster was arbeitet sprecht ihr mit dem API Server, egal ob mit kubectl, helm, k9s etc..

  Validiert, ob rechte vorhanden sind, und anfragen sinn ergeben

  updated values in etcd

  bekanntester api server tool kube-apiserver
</aside>

+++

<!-- .slide: style="text-align: left;"> -->
#### scheduler
- Verteilt workload
- verantworlich für pods ohne node

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
- bringt cluster von "ist" -> "soll"
- Managed Nodes
- mitteilung an scheuduler wenn node down

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
- verwaltet pods
- auf jeden node installiert
- verantwortlich für status

<aside class="notes">
  kubelet bekommt info von controller und fürt auf node aus

  startet stoppt updated pods auf node

  überwacht pods ob sie gewünschten status haben
</aside>

+++

<!-- .slide: style="text-align: left;"> -->
### Kube Proxy
- verwaltet Netzwerkanfragen
- routet traffic zu gewünschten pod
- loadbalancer

<aside class="notes">
  der kube-proxy wird angefragt sobald eine Netzwerkanfrage zum node kommt und leitet diese weiter zum gewünschten container

  übernimmt auch loadbalancing funktionen, sollten meherere Pods das gleiche machen (mehere nginx zum beispiel)
</aside>

+++

<!-- .slide: style="text-align: left;"> -->
## Weitere Komponenten
- CNI
- Container-Runtime

<aside class="notes">
  es gibt noch weiter komponente

  Network Plugin, also CNI schreibt Network interfaces in die

  Container Runtime

  dies ist z.B. containerd, cri-o oder die docker engine

  meistens containerd

  diese Software ist zuständig um die Container laufen zu lassen
</aside> 

+++

<!-- .slide: style="text-align: left;"> -->
## OpenSource

<aside class="notes">
  Alle diese Komponenten sind opensource und theoretisch austauschbar, auch wenn einige so sehr in den kubernetes core eingebaut sind, dass diese nur mit sehr hohen aufwand getauscht werden können
</aside>

+++

<!-- .slide: style="text-align: left;"> -->
## Namespaces
- separierungseinheit in Kubernetes

<aside class="notes">
  Namespaces sind ein ganz wichtiger punkt in Kubernetes

  separiert im Cluster verschiedene Anwendungen

  Gleiche Anwendung kann im Cluster in verschiedenen Namespaces mit gleichen Namen laufen
</aside>

+++

<!-- .slide: style="text-align: left;"> -->
## Ausfallsicherheit

1. Container Health Check
    1. readyness
    1. liveness
1. Hostsystemausfall
1. Update

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
### Manifeste
- in yaml definiert

+++

### Beispiel

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
## Install Tools
Kubernetes Dokumentation:
- kubectl <!-- .element: class="fragment" data-fragment-index="1" -->
- krew <!-- .element: class="fragment" data-fragment-index="2" -->
- kind <!-- .element: class="fragment" data-fragment-index="3" -->
- k9s <!-- .element: class="fragment" data-fragment-index="4" -->
- helm <!-- .element: class="fragment" data-fragment-index="5" -->

<aside class="notes">
  kurze hintergrundinfos über einzelne tools

  kubectl: basis tool zum interagieren mit kubernetes cluster

  krew: kubectl plugin manager

  helm: package manager für kubernetes

  kind: kubernetes in docker, um kleine test cluster aufzubauen und kleine manifeste in kubernetes zu deployen

  k9s: beschreibung von Sandro gebraucht
</aside>

+++

<!-- .slide: style="text-align: left;"> -->
## kubectl plugins
- [krew](https://krew.sigs.k8s.io/docs/user-guide/setup/install/)
- [Liste](https://krew.sigs.k8s.io/plugins/) von verfügbaren Plugins

### Install <!-- .element: class="fragment" data-fragment-index="1" -->
- node-shell <!-- .element: class="fragment" data-fragment-index="1" -->
- df-pv <!-- .element: class="fragment" data-fragment-index="2" -->

---

# Objekttypen in K8s

+++

<!-- .slide: style="text-align: left;"> -->
### Pod
- Umfasst einen oder meherere Container
- Kleinste verwaltbares Objekt
- Jeder Pod bekommt eine IP Addresse (ClusterIP)

<aside class="notes">
  Ein Pod beinhaltet 1 bis n Container

  Pods haben ip addressen, da die aber dynamisch sind haben sie bei Neustart eine neue IP
</aside>

+++

<!-- .slide: style="text-align: left;"> -->
## Aufgabe 1

```sh
bin/w6s exercise k8s -n1
```
Zeit: ca. 5m

+++

<!-- .slide: style="text-align: left;"> -->
### Service
- Objekt um Pod im Netzwerk erreichbar zu machen
- Loadbalancing
- Dynamische IP's von Pods

<aside class="notes">
  Services werden genutzt um pods im Netzwerk erreichbar zu machen

  hat eine loadbalancing funktion, wenn mehere pods mit gleichem Label im Namespace sind wird die last aufgeteilt

  geht an die pods mit einem Label, daher sind dynamische IPs bei Pods keine Probleme
</aside>

+++

<!-- .slide: style="text-align: left;"> -->
## Aufgabe 2

```sh
bin/w6s exercise k8s -n2
```
Zeit: ca. 5m

+++

<!-- .slide: style="text-align: left;"> -->
### ReplicaSets
- Pods Replizieren
- Nachträglich nicht änderbar

<aside class="notes">
  ein replicaset wird genutzt um mehere identische Pods zu deployen

  problem bei replicaset ist, dass es nachträglich nicht änderbar ist

  um neue version von pod zu deployen muss erst das alte replicaset gelöscht und das neue deployt werden

  frage an teilnehmer: zu was führt das? > Downtime
</aside>

+++

<!-- .slide: style="text-align: left;"> -->
## Aufgabe 3

```sh
bin/w6s exercise k8s -n3
```
Zeit: ca. 5m

+++

<!-- .slide: style="text-align: left;"> -->
### Deployment
- Art ReplicaSets zu verwalten
- Updates
- am weitesten verbreitete art

<aside class="notes">
  wie ihr herausgefunden habt, ist ein Deployment die bessere art ReplicaSets zu verwalten

  ein Deployment kann man Unterbrechungsfrei Updaten (wenn zumindestes 2 Replicas verfügbar)

  deployments werden am häufigsten genutzt um pods zu deployen
</aside>

+++

<!-- .slide: style="text-align: left;"> -->
### DaemonSet
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
## Aufgabe 4

```sh
bin/w6s exercise k8s -n4
```
Zeit: ca. 5m

+++

<!-- .slide: style="text-align: left;"> -->
### StatefulSet
- Persistente Pods
- Geordnetes Update/Shutdown

<aside class="notes">
  StatefulSets sind sinnvoll, wenn man erzielen möchte, dass eine anwendung ihren status nicht verliert

  z.B Datenbanken sind klassische Anwendungen welche man in diesem zustand haben möchte.
</aside>

+++

<!-- .slide: style="text-align: left;"> -->
### Job
- Einmalige Ausführung eines Commands in einem Pod
    - Datenbank Backup

<aside class="notes">
  jobs sind praktisch um einzelne kommandos auszuführen

  z.B prüfen ob ein service im cluster erreichbar ist, datenbank backups zu erstellen
</aside>

+++

<!-- .slide: style="text-align: left;"> -->
## Aufgabe 5

```sh
bin/w6s exercise k8s -n5
```
Zeit: ca. 5m

+++

<!-- .slide: style="text-align: left;"> -->
### CronJobs
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
## Aufgabe 6

```sh
bin/w6s exercise k8s -n6
```
Zeit: ca. 5m

<aside class="notes">
  kubectl create job nicht in cronjob k8s doku
</aside>

+++

<!-- .slide: style="text-align: left;"> -->
### ConfigMaps
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
## Aufgabe 7

```sh
bin/w6s exercise k8s -n7
```
Zeit: ca. 5m

+++

<!-- .slide: style="text-align: left;"> -->
## Aufgabe 8

```sh
bin/w6s exercise k8s -n8
```
Zeit: ca. 5m

<aside class="notes">
  Diesmal die Aufgabe vor dem API Objekt

  Teilnehmer sollen secrects finden 
</aside>

+++

<!-- .slide: style="text-align: left;"> -->
### Secret
- Speicherung vertraulicher Daten
- Unverschlüsselt in etcd DB

<aside class="notes">
  secrets gibt es um vertrauliche daten zu speichern

  standartmäßig liegen diese daten aber unverschlüsselt im etcd

  einbindung ähnlich wie bei configmaps
</aside>

---

<!-- .slide: style="text-align: left;"> -->
# Buchempfehlungen
- [Kubernetes Up & Running](https://www.amazon.de/Kubernetes-Up-Running-Brendan-Burns/dp/1492046531)
- [Kubernetes Best Practices](https://www.amazon.de/Kubernetes-Best-Practices-Blueprints-Applications/dp/1492056472)

---

# Fragen
- Hab ihr noch Fragen an uns?
