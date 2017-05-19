#!/usr/bin/env bash

# cut required due to OSX's output of wc with prepended tab
if [[ $(docker images --filter "reference=golang-compiler"  --quiet | wc -l | tr -d ' ') -eq 0 ]]; then
  ./infra/docker/golang-compiler/build.sh
fi

# Build broker image with binary
docker run --rm -t \
           -v "$PWD/broker:/go/src/broker:rw"\
           -v "$PWD/dist:/go/dist:rw" \
           golang-compiler \
           bash -c "cd /go/src/broker && glide install && go build -o /go/dist/broker main.go"

# Build randgen image with binary
docker run --rm -t \
           -v "$PWD/randgen:/go/src/randgen:rw"\
           -v "$PWD/dist:/go/dist:rw" \
           golang-compiler \
           bash -c "cd /go/src/randgen && glide install && go build -o /go/dist/randgen main.go"

# Package the binary / also handled by docker-compose
# ./infra/docker/broker/build.sh
# ./infra/docker/randgen/build.sh

docker-compose up --build
docker-compose down
