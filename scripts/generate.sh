#!/usr/bin/env bash

set -euo pipefail

function main() {
    # Get the last argument, i.e. PROJECT
    for project; do true; done
    if [[ -z "${project}" ]]; then
        usage
        exit 1;
    fi

    # =======================================
    # GENERATING PROCESS
    # =======================================
    pushd "$project"
      go generate -x ./proto
    popd
}

function usage() {
    echo "Usage:"
    echo "    ./scripts/generate <PROJECT>"
}

main "$@"
