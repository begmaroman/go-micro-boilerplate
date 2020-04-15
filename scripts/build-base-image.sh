#!/usr/bin/env bash

set -xeuo pipefail

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
ROOT="$DIR/.."

# Here we define the image. It is also possible to use the remote image here.
# If you use AWS container registry, use it here.
IMAGE="go-micro-boilerplate-base:1.12-stretch-0"

pushd "$ROOT" > /dev/null
    # Build the base image
    docker build -f backend/Dockerfile -t "$IMAGE" .

    # When it's successfully built, you can push in in your remote registry.
    # Use the command like this: docker push "$IMAGE"
popd > /dev/null
