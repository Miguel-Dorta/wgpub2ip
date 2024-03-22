#!/bin/bash

PROGRAM_NAME="wgpub2ip"
OUT_DIR="build"
OS_ARCH_LIST=(
  "darwin/amd64"
  "darwin/arm64"
  "freebsd/386"
  "freebsd/amd64"
  "freebsd/arm"
  "freebsd/arm64"
  "linux/386"
  "linux/amd64"
  "linux/arm"
  "linux/arm64"
  "linux/mips"
  "linux/mips64"
  "linux/mips64le"
  "linux/mipsle"
  "linux/ppc64"
  "linux/ppc64le"
  "linux/riscv64"
  "linux/s390x"
  "netbsd/386"
  "netbsd/amd64"
  "netbsd/arm"
  "netbsd/arm64"
  "openbsd/386"
  "openbsd/amd64"
  "openbsd/arm"
  "openbsd/arm64"
  "openbsd/mips64"
  "plan9/386"
  "plan9/amd64"
  "plan9/arm"
  "solaris/amd64"
  "windows/386"
  "windows/amd64"
  "windows/arm"
  "windows/arm64"
)

mkdir -p "$OUT_DIR"

for OS_ARCH in "${OS_ARCH_LIST[@]}"; do
    IFS='/' read -r -a PARTS <<< "$OS_ARCH"
    OS="${PARTS[0]}"
    ARCH="${PARTS[1]}"

    # Export vars
    export GOOS="$OS"
    export GOARCH="$ARCH"

    # Beautify OS and ARCH
    if [ "$OS" = "darwin" ]; then
        OS="macos"
    fi
    if [ "$ARCH" = "386" ]; then
        ARCH="x86"
    elif [ "$ARCH" = "amd64" ]; then
        ARCH="x86_64"
    fi

    BINARY_NAME="$PROGRAM_NAME"
    if [ "$OS" = "windows" ]; then
        BINARY_NAME="$BINARY_NAME.exe"
    fi

    go build -o "$BINARY_NAME" -ldflags="-X main.Version=$(git describe --tags)" wgpub2ip.go
    zip "$OUT_DIR/${PROGRAM_NAME}_${OS}_${ARCH}" "$BINARY_NAME" LICENSE README.md

    rm "$BINARY_NAME"
done
