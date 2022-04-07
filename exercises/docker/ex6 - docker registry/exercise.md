1. Starte lokal eine Docker Registry, die auf dem Host-Port `5050` lauscht.

2. Verifiziere, dass die Registry auf Port 5050 erreichbar ist und mit `... succeeded` antwortet.
Nutze hierfür das Image `nc` von Aufgabe 1.

3. Pushe das Image `nc` in diese private Registry.

4. Lösche das Image `nc` vollständig vom Host (lokale Registry).

5. Verifiziere, dass 2. nun nicht mehr funktioniert.

6. Pulle das `nc` Image aus der privaten Registry.

7. Verifiziere, dass 2. wieder geht.

8. Lösche die private Registry.

9. Verifiziere, dass 2. jetzt erwartungsgemäß mit `... failed: Connection refused` antwortet.
