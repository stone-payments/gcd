FROM golang:1.8-onbuild AS builder

MAINTAINER Guilherme Paixão <gpaixao@stone.com.br>

ENV WD /go/src/github.com/stone-payments/gcd

COPY . $WD

WORKDIR $WD

RUN CGO_ENABLED=0 go build -o ./bin/gcd ./

FROM alpine

ENV WD /go/src/github.com/stone-payments/gcd

COPY --from=builder $WD/bin/gcd /usr/bin

ENV GCD_DOCKER_HOST "unix:///var/run/docker.sock"
ENV GCD_SWEEP_INTERVAL "60"
ENV GCD_REMOVE_IMAGES "true"
ENV GCD_REMOVE_HEALTHY_CONTAINERS_EXITED "true"

CMD gcd -docker-host=$GCD_DOCKER_HOST \ 
        -sweep-interval=$GCD_SWEEP_INTERVAL \
        -remove-images=$GCD_REMOVE_IMAGES \
        -remove-healthy-containers-exited=$GCD_REMOVE_HEALTHY_CONTAINERS_EXITED
