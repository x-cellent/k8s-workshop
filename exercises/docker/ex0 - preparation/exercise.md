Falls noch nicht geschehen, erweitere die Gruppe `docker` um den eigenen Benutzer:

```sh
sudo usermod -aG docker ${USER}
```

Beende anschließend deine aktuelle User-Session und melde
dich erneut an.
Folgender Befehl sollte jetzt keinen Fehler werfen:

```sh
docker images
```

Dies erleichtert den Umgang mit der `docker` CLI, da auf ein
vorangestelltes `sudo` verzichtet werden kann.

ACHTUNG:
Wenn man einen fremden Benutzer der `docker` Gruppe hinzufügt,
erteilt man ihm automatisch `sudo` Rechte!
Daher wird das nur auf Dev-Systemen empfohlen, keinesfalls auf
Produktivsystemen.
