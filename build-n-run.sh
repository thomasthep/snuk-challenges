#!/usr/bin/env bash
set -e

# cut required due to OSX's output of wc with prepended tab
if [[ $(docker images --filter "reference=golang-compiler"  --quiet | wc -l | tr -d ' ') -eq 0 ]]; then
  ./infra/docker/golang-compiler/build.sh
fi

# Build collector binary
docker run --rm -t \
           -v "$PWD/collector:/go/src/github.com/thomasthep/snuk-challenges/collector:rw"\
           -v "$PWD/dist:/go/dist:rw" \
           golang-compiler \
           bash -c "cd /go/src/github.com/thomasthep/snuk-challenges/collector && glide install && go build -o /go/dist/collector main.go"

# Build randgen binary
docker run --rm -t \
           -v "$PWD/randgen:/go/src/github.com/thomasthep/snuk-challenges/randgen:rw"\
           -v "$PWD/dist:/go/dist:rw" \
           golang-compiler \
           bash -c "cd /go/src/github.com/thomasthep/snuk-challenges/randgen && glide install && go build -o /go/dist/randgen main.go"

# Package the binary / also handled by docker-compose
# ./infra/docker/collector/build.sh
# ./infra/docker/randgen/build.sh

docker-compose up --build
docker-compose down
