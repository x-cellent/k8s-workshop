20m
1. Anlegen einer geeigneten `index.html`:

```sh
cat <<EOF > index.html
<html>
    <head>
        <title>Hello World!</title>
    </head>
    <body>
        <h1>Hello World!</h1>
    </body>
</html>
EOF
```

Einbinden dieser Datei als Docker Volume mit Port-Forwarding Host (8080) -> Container (80):

```sh
docker run -d --name web -v $PWD/index.html:/usr/share/nginx/html/index.html -p 8080:80 nginx
```

2. Verbindung mit dem Container via Bash-Shell:

```sh
docker exec -it web bash
```

Ausgabe der NginX Version:

```sh
nginx -v
```

3. Entweder:

```sh
docker exec web /bin/bash -c "sed -i 's@World@Kubernetes@g' /usr/share/nginx/html/index.html"
```
O
der (weil wir die `index.html` in den Container gemounted haben):

```sh
sed -i 's@World@Kubernetes@g' index.html
```

4.

```sh
docker cp website.html web:/usr/share/nginx/html/index.html
```

Die lokale `index.html` ist an diese Stelle im Container gemounted, es ist ein und dieselbe Datei.

5.

```sh
docker rm -f web
```
