Erweitere die Gruppe `docker` um den eigenen Benutzer:

```sh
sudo usermod -aG docker ${USER}
```

Dies erleichtert den Umgang mit der `docker` CLI, da auf ein
vorangestelltes `sudo` verzichtet werden kann.

ACHTUNG:
Wenn man einen fremden Benutzer der `docker` Gruppe hinzufügt,
erteilt man ihm automatisch `sudo` Rechte!
Daher wird das nur auf Dev-Systemen empfohlen, keinesfalls auf
Produktivsystemen.
