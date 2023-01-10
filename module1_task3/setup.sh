#!/usr/bin/env bash

# Install required packages
apt-get update
apt-get install -y make hugo make

# Build the website
make build

# If the build fails, print the error message and exit with the appropriate error code
if [ $? -ne 0 ]; then
  make build 2>&1 | tee /dev/stderr
  exit $?
fi

# If the build succeeds, exit with a success code
exit 0
