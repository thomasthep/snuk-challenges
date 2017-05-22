#!/usr/bin/env bash

docker build -t collector:latest -f ./infra/docker/collector/Dockerfile .
