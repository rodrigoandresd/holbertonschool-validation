#!/usr/bin/env bash

# Install required packages
apt-get update && apt-get install -y hugo make

# Build the website
make build
