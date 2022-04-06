
Ordner Anlegen namens nginx-volume in welcher eine Hello World datei liegt, diesen mit einem Docker Run command als volume einbinden

```sh
docker run -d -p 8080:80 --name web --volume $PWD/nginx-volume/:/usr/share/nginx/html/ nginx
```

```sh
docker exec -it web bash
```