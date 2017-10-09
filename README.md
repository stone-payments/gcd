# gcd

[![Build Status](https://travis-ci.org/stone-payments/gcd.svg?branch=master)](https://travis-ci.org/stone-payments/gcd)
[![Go Report Card](https://goreportcard.com/badge/github.com/stone-payments/gcd)](https://goreportcard.com/report/github.com/stone-payments/gcd)

## Description
This project is a garbage collector for docker images and container, it exists because today our machines contains a lot of garbage from Docker components.

## Binary

### Building binary

```bash
> git clone https://github.com/stone-payments/gcd.git
> make build
```

### Running builded binary
```bash
> ./bin/gcd
```


### Parameters

- __-docker-host__ Set docker host target
- __-sweep-interval__ Set interval between sweep, this parameter only measures in second
- __-remove-images__ Set enable to remove images that isn't any container dependencies
- __-remove-healthy-containers-exited__ Set enable to remove containers exited with code 0

## Docker

### Use docker image from Docker Hub
```bash
> docker run --name gcd -v /var/run/docker.sock:/var/run/docker.sock guiferpa/gcd
```

### Building docker image
```bash
> make build-image
```
> :warning: This project use multi-stage build to build docker image, required use >=17.05 docker version

### Running builded image
```bash
> docker images

REPOSITORY               TAG                 IMAGE ID            CREATED             SIZE
guiferpa/gcd             latest              04ba50851638        16 seconds ago      9.7 MB
docker.io/golang         1.8-onbuild         5d82e356477f        2 weeks ago         699 MB
docker.io/alpine         latest              7328f6f8b418        6 weeks ago         3.97 MB

> docker run --name gcd -v /var/run/docker.sock:/var/run/docker.sock guiferpa/gcd
```

### Environment variables

- __GCD_DOCKER_HOST:__ A env variable to set __-docker-host__, by default use `/var/run/docker.sock:/var/run/docker.sock`
- __GCD_SWEEP_INTERVAL:__ A env variable to set __-sweep-interval__, by default use 60 seconds
- __GCD_REMOVE_IMAGES__: A env variable to set __-remove-images__, by default use `true`
- __GCD_REMOVE_HEALTHY_CONTAINERS_EXITED__: A env variable to set __-remove-healthy-containers-exited__, by default use `true`

### License
[MIT](https://github.com/stone-payments/gcd/blob/master/LICENSE)
