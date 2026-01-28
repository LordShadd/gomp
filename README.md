# GOMP

Open Multiplayer C-API implementation for Golang.

Under development. Do not use in production under any circumstances.

C-API: https://github.com/openmultiplayer/open.mp-capi

Building command:

```sh
CGO_ENABLED=1 GOOS=linux GOARCH=386 go build -buildmode=c-shared -o basic.so examples/basic/main.go
```
