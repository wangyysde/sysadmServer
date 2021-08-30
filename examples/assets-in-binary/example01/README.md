# Building a single binary containing templates

This is a complete example to create a single binary with the
[wangyysde/sysadmSErver][sysadmServer] Web Server with HTML templates.

[gin]: https://github.com/wangyysde/sysadmServer

## How to use

### Prepare Packages

```sh
go get github.com/wangyysde/sysadmServer
go get github.com/jessevdk/go-assets-builder
```

### Generate assets.go

```sh
go-assets-builder html -o assets.go
```

### Build the server

```sh
go build -o assets-in-binary
```

### Run

```sh
./assets-in-binary
```
