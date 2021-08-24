## How to run this example

1. run grpc server

```sh
$ go run grpc/server.go
```

2. run sysadmserver

```sh
$ go run sysadmServer/main.go
```

3. use curl command to test it

```sh
$ curl -v 'http://localhost:8052/rest/n/thinkerou'
```
