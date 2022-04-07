15m
1.

```sh
docker run -d --name registry -p 5050:5000 registry
```

2.

```sh
docker run -t --net host nc -vz http://localhost 5050
```

3.

```sh
docker tag nc localhost:5050
docker push localhost:5050/nc
```

4.

```sh
docker rmi localhost:5050/nc
docker rmi nc
```

5.

```sh
docker run -t --net host nc -vz http://localhost 5050
```

6.

```sh
docker pull localhost:5050/nc
```

7.

```sh
docker tag localhost:5050/nc nc
docker rmi localhost:5050/nc
docker run -t --net host nc -vz http://localhost 5050
```

8.

```sh
docker rm -f registry
```

9.

```sh
docker run -t --net host nc -vz http://localhost 5050
```
