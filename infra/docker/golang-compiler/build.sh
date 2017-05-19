#!/usr/bin/env bash

docker build -t golang-compiler:latest -f ./infra/docker/golang-compiler/Dockerfile .
