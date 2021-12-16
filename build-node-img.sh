#!/bin/bash

docker build --no-cache -f Dockerfile -t susy:nodev2 .
docker build --no-cache -f Dockerfile.node -t susy:validatorv2 .
