5s
1.
Anlegen einer `index.html`:
```html
<html>
    <head>
        <title>Hello World!</title>
    </head>
    <body>
        <h1>Hello World!</h1>
    </body>
</html>
```

Einbinden dieser Datei als Docker Volume mit Port-Forwarding 8080 (Host) -> 80 (Container):
```sh
docker run -d -p 8080:80 --name web --volume $PWD/index.html:/usr/share/nginx/html/index.html nginx
```

2.
```sh
docker exec -it web bash
nginx -v
```

3.
```sh
docker rm -f web
```
