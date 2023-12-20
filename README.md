# traefik plugin fortune

```shell
docker compose up
```

```console
$ curl localhost/who
Hostname: a3c72290d619
IP: 127.0.0.1
IP: 192.168.144.3
RemoteAddr: 192.168.144.2:39244
GET /who HTTP/1.1
Host: localhost
User-Agent: curl/7.81.0
Accept: */*
Accept-Encoding: gzip
X-Fortune: Good
X-Forwarded-For: 192.168.144.1
X-Forwarded-Host: localhost
X-Forwarded-Port: 80
X-Forwarded-Proto: http
X-Forwarded-Server: 6733240255a7
X-Real-Ip: 192.168.144.1
```
