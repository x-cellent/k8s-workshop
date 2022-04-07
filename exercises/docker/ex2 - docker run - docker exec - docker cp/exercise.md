1. Erzeuge einen `nginx` Docker Container names `web`, der auf der Welcome-Page `Hello World!` ausgibt.
Erstelle dazu eine lokale `index.html` Datei, auf die der Container zugreift.
Starte den Container im Hintergrund (detached) so, dass du die Welcome-Page unter `http://localhost:8080` erreichen kannst.

2. Ermittle die im Container verwendete `nginx` Version.

3. Ändere die Ausgabe des laufenden Containers zu `Hello Kubernetes!`.
Verifiziere dies im Browser durch Drücken von F5.

4. Kopiere die Datei `website.html` über die Welcome-Page im Container.
Verifiziere die Änderung im Browser.
Warum hat sich die lokale `index.html` ebenfalls geändert?

5. Lösche den Container wieder.
