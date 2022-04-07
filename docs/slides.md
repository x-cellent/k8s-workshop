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
- Verpacken der Software *mitsamt* aller Dependencies (Image)  <!-- .element: class="fragment" data-fragment-index="1" -->
    - Nichts darüber hinaus (Betriebssytem notwendig?)  <!-- .element: class="fragment" data-fragment-index="2" -->
- Container-Runtime für alle Plattformen  <!-- .element: class="fragment" data-fragment-index="3" -->

+++

<!-- .slide: style="text-align: left;"> -->
## Umsetzung
- Linux  <!-- .element: class="fragment" data-fragment-index="1" -->
- Idee: Container teilen sich Kernel  <!-- .element: class="fragment" data-fragment-index="2" -->
- LXC; basierend auf Kernel-Funktionalitäten  <!-- .element: class="fragment" data-fragment-index="3" -->
    - namespaces  <!-- .element: class="fragment" data-fragment-index="4" -->
    - cgroups  <!-- .element: class="fragment" data-fragment-index="5" -->
- Docker erweitert LXC um  <!-- .element: class="fragment" data-fragment-index="6" -->
    - CLI zum Starten und Verwalten von Containern  <!-- .element: class="fragment" data-fragment-index="7" -->
    - Image Registry  <!-- .element: class="fragment" data-fragment-index="8" -->
    - Networking  <!-- .element: class="fragment" data-fragment-index="9" -->
    - docker-compose  <!-- .element: class="fragment" data-fragment-index="10" -->

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
    - [Dockerfile](https://docs.docker.com/engine/reference/builder/) <!-- .element: class="fragment" data-fragment-index="3" -->
1. Container <!-- .element: class="fragment" data-fragment-index="4" -->
1. Image Registry <!-- .element: class="fragment" data-fragment-index="5" -->

+++

<!-- .slide: style="text-align: left;"> -->
## Dockerfile
- Image-*Rezept* mit u.a. folgenden Instruktionen:
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
bin/w6p exercise docker -n1
```
Zeit: ca. 5 min

+++

<!-- .slide: style="text-align: left;"> -->
## Container starten

```sh
docker run [--name NAME] [-i] [-t] [-d|--rm] [--net host|NETWORK] [-v HOST_PATH:CONTAINER_PATH] \
    [-p HOST_PORT:CONTAINER_PORT] [-u UID:GID] IMAGE [arg(s)]
```

- Noch viel mehr Flags möglich <!-- .element: class="fragment" data-fragment-index="1" -->
- [Referenz](https://docs.docker.com/engine/reference/run/) <!-- .element: class="fragment" data-fragment-index="2" -->

+++

## Command in Container triggern

```sh
docker exec [-i] [-t] CONTAINER COMMAND
```

Via Shell in den Container "springen":

```sh
docker exec [-i] [-t] CONTAINER COMMAND
```

+++

## Dateien kopieren

...vom Host in den Container:

```sh
docker cp HOST_FILE CONTAINER_NAME:CONTAINER_FILE
```

...vom Container in das Host-FS:

```sh
docker cp CONTAINER_NAME:CONTAINER_FILE HOST_FILE
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
- Warum nicht Docker Swarm?
- Mehr Flexibilität
- Eingebautes Monitoring und Logging
- Bereitstellung von Storage 
- Größere userbase

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
1. Architektur von Kubernetes
1. Einrichtung euerer Umgebung
1. Basisobjekte Kubernetes mit Übungen

---

# Kubernetes
*Kubernetes ist ein Open-Source-System  zur Automatisierung der Bereitstellung, Skalierung und Verwaltung von Container-Anwendungen*

+++

### Kubernetes

- Ursprünglich 2014 entwickelt von Google
- Abgegeben 2015 an die Cloud Native Compute Fondation (CNCF)

<aside class="notes">
  Weiterentwicklung von Google Borg

  aktuelle version 1.23.5
</aside>

+++

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

## Architektur von Kubernetes
![image](https://upload.wikimedia.org/wikipedia/commons/b/be/Kubernetes.png)

<aside class="notes">
  2 arten der nodes, master/controlPlane und worker
</aside>

+++

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

### Kube-Controller-Manager
- bringt cluster von "ist" -> "soll"
- Managed Nodes
- mitteilung an scheuduler wenn node down

<aside class="notes">
  bekommt von scheduler meldung, wenn pod zu node kommen soll und übermittelt dies

  überwacht nodes, wenn einer down ist, mitteilung an scheuduler, node down bitte pods neu verteilen
</aside>

+++

## Nodes

<aside class="notes">
  die nachfolgende Software ist auf allen nodes, also master und worker installiert
<aside>

+++

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

### Kube Proxy
- verwaltet Netzwerkanfragen
- routet traffic zu gewünschten pod
- loadbalancer

<aside class="notes">
  der kube-proxy wird angefragt sobald eine Netzwerkanfrage zum node kommt und leitet diese weiter zum gewünschten container

  übernimmt auch loadbalancing funktionen, sollten meherere Pods das gleiche machen (mehere nginx zum beispiel)
</aside>

+++

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

## OpenSource

<aside class="notes">
  Alle diese Komponenten sind opensource und theoretisch austauschbar, auch wenn einige so sehr in den kubernetes core eingebaut sind, dass diese nur mit sehr hohen aufwand getauscht werden können
</aside>

+++

## Namespaces
- separierungseinheit in Kubernetes

<aside class="notes">
  Namespaces sind ein ganz wichtiger punkt in Kubernetes

  separiert im Cluster verschiedene Anwendungen

  Gleiche Anwendung kann im Cluster in verschiedenen Namespaces mit gleichen Namen laufen
</aside>

+++

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

---

# Fragen
- Hab ihr noch Fragen an uns?

---

## Wichtige Ressourcen
1. kubectl cheat sheet: https://kubernetes.io/docs/reference/kubectl/cheatsheet
1. kubernetes docs: https://kubernetes.io/docs/concepts/overview

---

# Setup euerer Umgebung
- clonen dieses Repos und erstellen des kommandozeilen-tools

```sh
git clone https://github.com/x-cellent/k8s-workshop.git
cd k8s-workshop
make
```

<aside class="notes">
  wir haben ein repo mit allen schulungsunterlagen und einem in go geschriebenen kommandozeilen tool
</aside>

+++

## Install Tools
Kubernetes Dokumentation:
- kubectl <!-- .element: class="fragment" data-fragment-index="1" -->
- krew <!-- .element: class="fragment" data-fragment-index="2" -->
- helm <!-- .element: class="fragment" data-fragment-index="3" -->
- kind <!-- .element: class="fragment" data-fragment-index="4" -->
- k9s <!-- .element: class="fragment" data-fragment-index="5" -->

<aside class="notes">
  kurze hintergrundinfos über einzelne tools

  kubectl: basis tool zum interagieren mit kubernetes cluster

  krew: kubectl plugin manager

  helm: package manager für kubernetes

  kind: kubernetes in docker, um kleine test cluster aufzubauen und kleine manifeste in kubernetes zu deployen

  k9s: beschreibung von Sandro gebraucht
</aside>

+++

## Install kubectl plugins
- node-shell

---

# Objekttypen in K8s

<aside class="notes">
  gestern schon was gehört über Objekttypen
  
  heute genaue erklärung mit übungen
</aside>
+++

### Pod
- umfasst einen oder meherere Container
- niedrigstes verwaltbares Objekt
- jeder Pod bekommt IP addresse

<aside class="notes">
  wie ihr gestern schon erfahren habt, ist der pod das so ziemlich niedrigeste verwaltbare Objekt

  ein pod umfasst mindestens ein container kann aber auch mehrere umfassen

  pods haben ip addressen, da die aber dynamisch sind haben sie bei einem erneuten deploy eine neue IP
</aside>

+++

<!-- .slide: style="text-align: left;"> -->
## Aufgabe 1

```sh
bin/w6s exercise k8s -n1
```
Zeit: ca. 5m

+++

#### Lösungsbesprechung

<aside class="notes">
  ein teilnehmer erklärt seine lösung
</aside>

+++

### Service
- Objekt um Pod im Netzwerk erreichbar zu machen
- Loadbalancing
- Dynamische IP's von Pods

<aisde class="notes">
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

#### Lösungsbesprechung

<aside class="notes">
  ein teilnehmer erklärt seine lösung
</aside>

+++

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

#### Lösungsbesprechung

<aside class="notes">
  ein teilnehmer erklärt seine lösung
</aside>

+++

### Deployment
- Bessere art ReplicaSets zu verwalten
- Updates
- am weitesten verbreitete art

<aside class="notes">
  wie ihr herausgefunden habt, ist ein Deployment die bessere art ReplicaSets zu verwalten

  ein Deployment kann man Unterbrechungsfrei Updaten (wenn zumindestes 2 Replicas verfügbar)

  deployments werden am häufigsten genutzt um pods zu deployen
</aside>

+++

### DaemonSet
- jede Node bekommt ein Replica
- enorm ausfallsicher
- logs
- monitoring

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

#### Lösungsbesprechung
- deamonSet aus kubernetes Doku
- kubernetes Doku ist immer gut

<aside class="notes">
  ein teilnehmer erklärt seine lösung 
  
  kubernetes Doku ist immer ein guter ort sich infos zu holen, deployments anzuschauen/abzuschauen
</aside>

+++

### StatefulSet
- persistente Pods
- geordnetes Updaten

<aside class="notes">
  StatefulSets sind sinnvoll, wenn man erzielen möchte, dass eine anwendung ihren status nicht verliert

  z.B Datenbanken sind klassische Anwendungen welche man in diesem zustand haben möchte.
</aside>

+++

### Job
- ausführung eines commandes in einem pod
- datenbank backups

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

#### Lösungsbeschreibung

+++

### Cronjobs
- Mischung aus klassischen Cronjobs und Jobs
- regelmäßige ausführung eines jobs

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

+++

#### Lösungsbesprechung

<aside class="notes">
  ein teilnehmer erklärt seine lösung 

  kubectl create job nicht in cronjob k8s doku
</aside>

+++

### Configmaps
- speicherung von nicht vertraulichen daten
- einbindung in pods als
    - enviroment-variable
    - command-line argument
    - als datei in Volume
- kein reload von pods bei änderung von configmap

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

#### Lösungsbesprechung

<aside class="notes">
  ein teilnehmer erklärt seine lösung 
</aside>

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

#### Lösungsbesprechung

<aside class="notes">
  ein teilnehmer erklärt seine lösung 
</aside>

+++

### Secret
- speicherung vertraulicher daten
- unentschlüsselt in etcd db

<aside class="notes">
  secrets gibt es um vertrauliche daten zu speichern

  standartmäßig liegen diese daten aber unverschlüsselt im etcd

  einbindung ähnlich wie bei configmaps
</aside>

---

# Fragen
- Hab ihr noch Fragen an uns?