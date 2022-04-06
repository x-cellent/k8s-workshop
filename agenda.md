# Kubernetes Training für FI-TS/DevOps 2022

1. Vorstellung (30 min)

1. Agenda/ausblick (10 min)

1. Containerisierung (3 stunden)
    - Warum?
    - Unterschiede zu VMs
    - Das Dockerfile
    - Wichtige Befehle
    - Aufgabe: Erzeugen, Starten und Zugriff auf einen Container

12:30-13:30 Mittag

1. Monolithen vs. Microservices (10 min)

1. Container-Orchestrierung (20 min)
    - Warum?
    - Warum Kubernetes?
    - Cloud-Native Computing Foundation

1. Prinzipien hinter Kubernetes  (1h-1h 30m)
    - Pod != Container
    - Pods werden auf Nodes verteilt
    - Flaches Layer-3 für alle Pods
    - Oberhalb Pods gibt es Ordnungselemente
    - Komponenten folgen dem Controller-Prinzip

1. Fragen (30 min)

1. optional ausblick auf morgen (10m)

1. Komponenten eines Kubernetes-Clusters
    - Basis: apiserver, scheduler, kubelet, kube-controller-manager, kube-proxy, etcd
    - Verpflichtende Neben-Projekte: CNI, Container-Runtime
    - Optional: DNS und cloud-controller-manager

1. Basis Objekt-Typen in Kubernetes (jeweils mit Aufgaben)
    - Metadaten: kind, apiversion
    - Basis: Node, Namespace, Pod, Ressource-Limits
    - Konfigurationsparameter: ConfigMap, Secret
    - Workload: ReplicaSet, Deployment, DaemonSet, StatefulSet, Job
    - Zugriff: Service
    - Aufgabe zu: Port-Forwarding, Init-Container, SideCar-Container-Pattern

1. Fragen (30 min)

1. Helm
    - Was ist Helm? Vorteile und Alternativen
    - Aufbau eines Helm-Charts
    - helm Kommandos
    - Aufgabe: helm Chart erstellen

1. Weitere Objekt-Typen und Advanced Usage
    - Logging und Metriken
    - Koordinierter Zugriff auf Services: Ingress, IngressClass
    - Aufgabe: Controller und Regel deployen, per port-forward prüfen
    - Storage: HostPath, PersistentVolume, PersistentVolumeClaim, StorageClass
    - Aufgabe: verschiedene Storage-Klassen für Pods verwenden
    - Reglementierung von Traffic zwischen Pods: NetworkPolicy
    - Aufgabe: Explizite Freigabe einer Verbindung
    - Horizontal Pod Auto Scaler (HPA)
    - Restriktionen für das Scheduling von Pods auf Nodes: Affinity, AntiAffinity of Pods and Nodes
    - Aufgaben: Pods sollen auf Subset der Nodes verteilt werden, Pods dürfen nicht zusammen auf einem Node ausgeführt werden
    - RBAC+Sicherheit: Role, RoleBinding, ServiceAccount, PodSecurityPolicy (Deprecated)
    - Aufgabe: Roles&Role-Binding + Erstellung eines Anwendungs-Deployments, das einen ServiceAccount verwendet

1. Tipps und Tricks für den Alltag
    - Einrichtung der eigenen Shell
    - kubectx, kubens, stern, lens
    - NetworkPolicy Editor von Cilium
    - Verwendung von GitLab aus

1. Finance Cloud-Native
    - Generelle Architektur und Konzepte
        - "Kubernetes Cluster as Cattle"
        - Verwaltung durch Gardener: Soil, Seed, Shoot, Reconcilation
        - metal-stack als IaaS-Schicht: Partition, Machine, Firewall, Image, Networks 
        - Netzwerk (Internet, MPLS)
        - Storage: lokal, Lightbits, S3
        - DBaaS: Postgres
    - cloudctl und metalctl
    - Cluster Lifecycle: Erstellung, Update der Kubernetes Version, Update Worker-Image (mit Aufgaben)
    - Unterschiede zu lokalem Kubernetes / Hyperscalern
        - Privates Netz für Worker
        - Jeder Cluster hat eine Firewall
        - Loadbalancer für externen Zugriff wird direkt auf den Kubernetes Nodes produziert (metal-lb)
        - Zusätzliche Objekt-Typen: Firewall, ClusterwideNetworkPolicy
        - Zusätzliche Objekte: Firewall, Headless-Services zum Monitoring der Firewall
    - Status-Seite: https://status.fits.cloud/
 
1. Prometheus, Grafana, Thanos
    - Metrics-API
    - Exporter
    - AlertManager und AlertRules
    - Verknüpfung zu Chat-Tools
    - Dashboard
    - PromQL
    - Aggregation mit Thanos