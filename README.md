# GOMP

Open Multiplayer C-API implementation for Golang.

Under development. Do not use in production under any circumstances.

C-API: https://github.com/openmultiplayer/open.mp-capi

Building command:

```sh
CGO_ENABLED=1 GOOS=linux GOARCH=386 go build -buildmode=c-shared -o basic.so examples/basic/main.go
```

## Extra

### colmap

A Go port of [ColAndreas](https://github.com/Pottus/ColAndreas). Provides GTA SA map collision via raycasting using Bullet 3 directly via CGO.

**Prerequisites:**
- C++ Compiler (GCC/Clang)
- *Note: Bullet Physics 3 is internally vendored and statically built via `colmap/install_bullet32.sh`, so no system-wide installation (like `libbullet-dev`) is required.*

**Map Database:**
You must generate the collision file (`ColAndreas.cadb`) using the official [ColAndreas Wizard](https://github.com/Pottus/ColAndreas) and place it in your server's `scriptfiles/` directory before testing.

For a complete working example of how to initialize the database and interact with it (e.g., raycasting on map clicks), check the [`_examples/colmap/`](_examples/colmap/) directory.
