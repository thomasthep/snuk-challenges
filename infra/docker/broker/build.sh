#!/usr/bin/env bash

docker build -t broker:latest -f ./infra/docker/broker/Dockerfile .
