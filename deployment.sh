#!/usr/bin/env bash

# sets gin to release mode
export GIN_MODE=release

while getops u:f:a: flag
do
  case "${flag}" in
    --skip-test) skip-test="${OPTARG}";;
  esac
done

if [ "${skip-test}" == true ]; then
  # run tests
  if go test ./...; then
    echo "test passed"
  else
    echo "Test failed. Exiting..."
    exit 0
  fi
fi

# build go project to executable
GOOS="linux" GOARCH="amd64" go build .

# build container image
docker build -t ropc-service .

# deploy image
docker-compose up --remove-orphans -d
