# colmap Example

This is a demonstration of how to integrate the `colmap` module (Go rewrite of ColAndreas) into a gomp-based plugin. This plugin initializes the Map Database and uses the `OnPlayerClickMap` event to RayCast collisions and teleport the player directly to the map's ground.

## Prerequisites

- C++ Compiler (GCC/Clang)
- *Note: Bullet Physics 3 is internally vendored and statically built via `colmap/install_bullet32.sh`, so no system-wide installation (like `libbullet-dev`) is required.*

## Requirements

Before starting the server to test this plugin, you must provide the map collision file (`ColAndreas.cadb`):
1. Download or generate it using the official [ColAndreas Wizard](https://github.com/Pottus/ColAndreas).
2. Place the generated file inside your server's `scriptfiles/` folder (so it exists at `scriptfiles/ColAndreas.cadb`).

## Building the Example

```sh
CGO_ENABLED=1 GOOS=linux GOARCH=386 go build -buildmode=c-shared -o colmap_example.so
```

Then load `colmap_example.so` in your `server.cfg` under `plugins`.

## Usage

When ingame:
1. Open the pause menu.
2. Go to the Map.
3. Right click anywhere on the Map (creates a targeted teleport).
4. The plugin will perform a vertical raycast from Z: 1000 down to Z: -1000.
5. If a collision is found, you will be teleported securely on top of the ground and notified of the model ID hit.
