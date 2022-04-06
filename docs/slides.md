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
- Erwartungen?

---

# Agenda

---

<!-- .slide: style="text-align: left;"> -->
## Tag 1
1. Container
1. Monolithen vs. Microservices
1. Container-Orchestrierung
1. Prinzipien hinter Kubernetes 

---

# Container

<aside class="notes">
  
</aside>

+++

## Vorteile Containarisierung
1. kleinere Images <!-- .element: class="fragment" data-fragment-index="1" -->
1. Geringerer Ressourcenverbrauch <!-- .element: class="fragment" data-fragment-index="2" -->
1. Erhöhte Sicherheit <!-- .element: class="fragment" data-fragment-index="3" -->
1. Abhängigkeiten mit im Image <!-- .element: class="fragment" data-fragment-index="4" -->

<aside class="notes">
  base Ubuntu Server Image ca 2 GB
  base Ubuntu Container Image ca 27 MB
  Alpine noch kleiner ca 2.7 MB
  
  Da kein Komplettes OS installiert wird
  Kernel wird sich geteilt mit host system
</aside>

+++

## Das Dockerfile
1. Datei zum Image bauen

```Dockerfile
FROM alpine:3.9 #base Image
RUN apk add --no-cache mysql-client #Commands welche man ausführen möchte, in diesem fall mysql-client installieren
ENTRYPOINT ["mysql"] #Startcommand welcher der container ausführen soll
```

+++

## Das Dockerfile !TO_DO!
1. Multi-Stage Dockerfiles auch Möglich <!-- .element: class="fragment" data-fragment-index="1" -->
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
  Unter Orchestrierung versteht man das Deployment Maintainance und Scaling

  Vorteile: Besser ressourcen nutzung

  bessere Bereitsstellung von Containern -> zero Downtime Updates
</aside>

+++

## Warum Kubernetes?
- warum nicht docker swarm?
- mehr flexibilität
- eingebautes monitoring und logging
- Bereitstellung von Storage 
- grösere userbase

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
- Auf Nodes

<aside class="notes">
  Die pods laufen auf sogenannten nodes, dies sind die server auf welchen die diversen Kubernetes core programme laufen, dies wird morgen genauer behandelt
</aside>

+++

### Ordnungselemente
- deployment <!-- .element: class="fragment" data-fragment-index="1" -->
- daemonSet <!-- .element: class="fragment" data-fragment-index="2" -->
- ReplicaSet <!-- .element: class="fragment" data-fragment-index="3" -->
- StatefulSet <!-- .element: class="fragment" data-fragment-index="4" -->
- Job <!-- .element: class="fragment" data-fragment-index="5" -->

<aside class="notes">
  wie man sehen kann, ist das wichtigste element der pod

  seltenst setzt man ihn einzeln ein

  normalerweise nutzt man ein übergeortnetes ordnungselement

  welche sind

  ReplicaSet: zum sicherstellen, eine genaue anzahl an pods zu haben, wird nur seltenst genutzt ansonsten:

  deployment: wird genutzt um replicasets auszurollen, da dies automatisierte rolling updates anbietet

  daemonset: vergewissert, dass auf allen nodes ein pod läuft

  StatefulSet: dazu da, eine statische anwendung zu deployen, nicht austauschbar datenbanken zum beispiel

  Job: ein pod welcher kurzzeitig für eine aufgabe ausgeführt wird

  weiterführung zum cronjob zum beispiel um backups auszuführen
</aside>


---

# Fragen
- Hab ihr noch Fragen an uns?

---

# Ausblick auf Morgen
- Architektur von Kubernetes
- Basis Objekte von Kubernetes

---

# Agenda

---

<!-- .slide: style="text-align: left;"> -->
## Tag 2
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
        - Scheudler <!-- .element: class="fragment" data-fragment-index="5" -->
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
  
  ueberwacht nodes, wenn einer down ist, mitteilung an scheuduler, node down bitte pods neu verteilen
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

## OpenSource

<aside class="notes">
  Alle diese Komponenten sind opensource und theoretisch austauschbar, auch wenn einige so sehr in den kubernetes core eingebaut sind, dass diese nur mit sehr hohen aufwand getauscht werden können
</aside>

+++

## Namepsaces
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

  dies wird erziehlt, indem man bei container 

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
Kubernetes Dokumentation:​
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

#### Aufgabe
- Bitte starte aufgabe k8s 1
```sh 
bin/w6s exercise k8s -n 1
```
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

#### Aufgabe
- Bitte starte aufgabe k8s 2
```sh
bin/w6s exercise k8s -n 2
```

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

#### Aufgabe
- Bitte starte aufgabe k8s 3
```sh
bin/w6s exercise k8s -n 3
```

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

#### Aufgabe
- Bitte starte aufgabe k8s 4
```sh
bin/w6s exercise k8s -n 4
```

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

#### Aufgabe

- Bitte starte aufgabe k8s 5
```sh
bin/w6s exercise k8s -n 5
```

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

#### Aufgabe

- Bitte starte aufgabe k8s 6
```sh
bin/w6s exercise k8s -n 6
```

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

#### Aufgabe
- Bitte starte aufgabe k8s 7
```sh
bin/w6s exercise k8s -n 7
```

+++

#### Lösungsbesprechung

<aside class="notes">
  ein teilnehmer erklärt seine lösung 
</aside>

+++

### Aufgabe
- Bitte starte aufgabe k8s 8
```sh
bin/w6s exercise k8s -n 8
```

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