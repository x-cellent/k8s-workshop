15m
1.

```sh
docker run -d --name redis redis
```

2.

```sh
docker inspect --format='{{.NetworkSettings.IPAddress}}' redis
# 172.17.0.2
```

3.

```sh
docker inspect --format='{{.Config.ExposedPorts}}' redis
# map[6379/tcp:{}]
```

4.

```sh
docker inspect --format='{{.Config.Volumes}}' redis
# map[/data:{}]
```

5.

```sh
docker inspect --format='{{.Config.Entrypoint}}' redis
# [docker-entrypoint.sh]
docker exec -t redis find / -type f -name docker-entrypoint.sh
# /usr/local/bin/docker-entrypoint.sh
# Find entrypoint arg(s)
docker inspect --format='{{.Config.Cmd}}' redis
# [redis-server]
```

Es wird ein redis-server gestartet.

6. Da die Container Kernel Ressourcen verwenden, sind Container Prozesse
normale Kernel-Prozesse:

```sh
docker exec -it redis bash
$ apt update
$ apt install -y procps
$ ps aux
# PID 1 process is 'redis-server *:6379*'
$ exit
ps aux | grep redis-server
# z.B. 3847193
```




