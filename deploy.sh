#!/bin/sh

set -e

if [ -z "$1" ]; then
  echo "Usage: $0 <tag>"
  exit 0
fi

echo "Docker build with tag: $1"
docker build -t bondhansarker/test-go-app:"$1" .
docker push bondhansarker/test-go-app:"$1"