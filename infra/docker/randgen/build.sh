#!/usr/bin/env bash

docker build -t randgen:latest -f ./infra/docker/randgen/Dockerfile .
