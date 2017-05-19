#!/usr/bin/env bash

# cut required due to OSX's output of wc with prepended tab
if [[ $(docker images --filter "reference=golang-compiler"  --quiet | wc -l | tr -d ' ') -eq 0 ]]; then
  ./infra/docker/golang-compiler/build.sh
fi

# Build image with binary
docker run --rm -t \
           -v "$PWD/randgen:/go/src/randgen:ro"\
           -v "$PWD/dist:/go/dist:rw" \
           golang-compiler \
           bash -c "cd /go/src/randgen && go build -o /go/dist/randgen main.go"

# Package the binary
./infra/docker/randgen-prod/build.sh
