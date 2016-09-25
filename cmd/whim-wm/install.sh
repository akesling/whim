#!/usr/bin/env bash

PREFIX=/usr/local
BUILD_DIR=$(mktemp -d)

cd ${BUILD_DIR}

echo "Getting and building Whim WM in ${BUILD_DIR}."
go build github.com/akesling/whim/cmd/wm

echo "Installing ${pwd}/bin/wm to ${PREFIX}/bin/wm"
sudo cp -f wm ${PREFIX}/bin/wm
sudo chmod 755 ${PREFIX}/bin/wm
