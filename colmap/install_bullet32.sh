#!/bin/bash
set -e

# Change to script directory
cd "$(dirname "$0")"

echo "Preparing 32-bit Bullet3 installation..."
BUILD_DIR="$(pwd)/bullet32"
mkdir -p "$BUILD_DIR"

echo "1. Cloning official repository (shallow clone)..."
rm -rf "$BUILD_DIR/src_repo"
git clone --depth 1 https://github.com/bulletphysics/bullet3.git "$BUILD_DIR/src_repo"

echo "2. Configuring CMake for 32-bit cross-compilation (-m32)..."
cd "$BUILD_DIR/src_repo"

cmake -B build -S . \
    -DCMAKE_CXX_FLAGS="-m32 -fPIC" \
    -DCMAKE_C_FLAGS="-m32 -fPIC" \
    -DUSE_DOUBLE_PRECISION=OFF \
    -DBUILD_SHARED_LIBS=OFF \
    -DBUILD_CPU_DEMOS=OFF \
    -DBUILD_OPENGL3_DEMOS=OFF \
    -DBUILD_BULLET2_DEMOS=OFF \
    -DBUILD_UNIT_TESTS=OFF \
    -DBUILD_EXTRAS=OFF \
    -DINSTALL_LIBS=ON \
    -DCMAKE_INSTALL_PREFIX="$BUILD_DIR"

echo "3. Compiling the physical libraries..."
cmake --build build -j$(nproc)

echo "4. Installing to local prefix and cleaning up..."
cmake --install build

cd ../..
rm -rf "$BUILD_DIR/src_repo"

echo ""
echo "Success! 32-bit static libraries (libBulletDynamics.a, etc) generated in colmap/bullet32/lib"
