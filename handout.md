# Docker

Container Runtime Engine

Tool zum bereitstellen von Containern

## CLI

Bau von Images
  - Instruktionen (Zutaten)
    - FROM, Definition eines Base Images
    - COPY, dateien in Container Kopieren
    - RUN, Definieren von Flags
    - ENTRYPOINT, Definieren eines Commands welches beim start ausgeführt werden soll

Start von Images
  - Container, N Image RO Layers plus ein leerer RW Layer on top 
  - RW Layer kann zur Laufzeit modifiziert werden
  - Delete File/Dir
        - Löscht, wenn im RW Layer vorhanden 
        - Versteckt, wenn in einem darunter liegenden RO Layer
    - Stirbt PID 1, stirbt der Container 
        - docker rm CONTAINER => RW Layer wird gelöscht 

- Image Registry vgl. AppStore


# Kubernetes

### Kubernetes
- Container Orchestrierungstool 
    - Verwaltet Pods 
        - besteht aus 1 bis N Containern 
        - eigene IP, eigenes Netzwerk 
    - Started, stoppt und überwacht Pods 
        - Verteilung auf Worker-Nodes 
        - Garantiert Pods Ressourcen (CPU/Memory) 
    - Self-Healing 
    - Dynamische Skalierung

### Objekte
- Pod
    - kleinste deploybare Einheit
    - beinhaltet 1 bis N Container
    - eigener Netzbereich und IP
- ReplicaSet
    - Stellt sicher, dass zu jeder Zeit genau N Pods laufen 
    - Matching über Labels 
- Deployment
    - Managed ein ReplicaSet
    - Bietet Versionierung und zero downtime Rollouts
- DeamonSet 
    - Spec wie Deployment nur ohne Replica Count 
    - Managed damit kein ReplicaSet 
    - Stattdessen je ein Replica pro Node
- Service
    - Loadbalancer für Pods
    - Auch hier Matching via Labels
    - Typen 
        - None (headless) 
        - ClusterIP (Default) 
        - NodePort 
        - Loadbalancer
- StatefulSet
    - Spec ähnlich zu Deployment
    - Geordnetes Starten (einer nach dem anderen)
    - Geordnetes Stoppen in umgekehrter Reihenfolge
- ConfigMap 
    - Plain-Text Key-Value Store 
    - Kann in Pods, Deployments, STSs und DSs gemounted werden 
- Secret 
    - base64 encoded Data Store
    - Kann in Pods, Deployments, STSs und DSs gemounted werden

### Kubernetes Tools
- kubectl 
    - CLI zur Interaktion mit k8s Clustern
    - cheatsheet https://kubernetes.io/de/docs/reference/kubectl/cheatsheet/
- krew 
    - kubectl Plugin Manager 
- k9s 
    - Terminal UI zur Interaktion mit k8s Clustern
- kind (Kubernetes in Docker) 
    - Single-Node k8s Cluster in Docker Container 
- helm
    - Paket-Manager für Kubernetes
    - vgl. mit apt für Ubuntu oder apk für Alpine
- Lens 
    - Graphical UI zur Interaktion mit k8s Clustern 
    - Nicht Teil des Workshops 
