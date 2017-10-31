# gcd

[![Build Status](https://travis-ci.org/stone-payments/gcd.svg?branch=master)](https://travis-ci.org/stone-payments/gcd)
[![Coverage Status](https://coveralls.io/repos/github/stone-payments/gcd/badge.svg?branch=master)](https://coveralls.io/github/stone-payments/gcd?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/stone-payments/gcd)](https://goreportcard.com/report/github.com/stone-payments/gcd)

## Description
This project is a garbage collector for docker images and containers. It was created to alleviate the problem of old
and unused containers laying around on Docker utilizing storage resources.

## Usage
### Running locally built ELF
```bash
./bin/gcd
```

### Running pre-built Docker image from Docker Hub
```bash
docker run --name gcd -v /var/run/docker.sock:/var/run/docker.sock guiferpa/gcd
```
### Parameters
gcd can be parametrized via command-line arguments and, if you're using the Docker image, by environment variables.

#### CLI arguments
- **-docker-host**: set docker host target (default: `/var/run/docker.sock`).
- **-sweep-interval**: set interval in seconds between sweeps (default: `60`).
- **-remove-images**: remove images that aren't container dependencies (default: `true`).
- **-remove-healthy-containers-exited**: remove containers exited with code 0 (default: `true`).

### Environment variables
- **GCD_DOCKER_HOST**: set docker host target (default: `/var/run/docker.sock`).
- **GCD_SWEEP_INTERVAL**: set interval in seconds between sweeps (default: `60`).
- **GCD_REMOVE_IMAGES**: remove images that aren't container dependencies (default: `true`).
- **GCD_REMOVE_HEALTHY_CONTAINERS_EXITED**: remove containers exited with code 0 (default: `true`).

## Building
### Building ELF
```bash
git clone https://github.com/stone-payments/gcd.git
make build
```

### Building docker image
> :warning: This project uses multi-stage build to build the docker image with the included Dockerfile. Docker >=17.05
> is required.
```bash
make build-image
```

### License
This code is licensed under the MIT license.
