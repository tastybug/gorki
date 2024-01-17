[![Go Report Card](https://goreportcard.com/badge/github.com/tastybug/gorki)](https://goreportcard.com/report/github.com/tastybug/gorki) | [![Master Build Status](https://travis-ci.com/tastybug/gorki.svg?branch=master)](https://travis-ci.com/tastybug/gorki) | [Task Tracker](./todo.diff) | [Docker Hub](https://hub.docker.com/repository/docker/tastybug/gorki)

# Gorki

Gorki is a simple, opinionated static site generator written in Go. It allows you to generate a blog from Markdown written articles.

### Development

* `make package` runs tests, builds the default binary.
* `make deploy` runs tests, builds for local env and amd64, creates and pushes docker image.

To aid development of the site, run a local webserver: 
```shell script
docker run -d -p 8080:80 --name local-serve --rm -v $PWD/target:/var/www/html/website tastybug/dockerized-nginx-local-serve nginx
```
to quickly check how the site looks like.

### Testing the Dockerized Gorki
To test from within the container, run
```shell script
docker run --name gorki --rm -ti -v $PWD/site:/app/site "tastybug/gorki:latest-multi" /bin/ash
```

### Run Dockerized Gorki
```shell script
docker run --name gorki --rm -v $PWD/site:/app/site "tastybug/gorki:latest-multi"
```

### Build and Publish

```
# this step only on the very first run
docker run --privileged --rm tonistiigi/binfmt --install all
docker buildx build --platform=linux/amd64,linux/arm/v7 -t tastybug/gorki:latest-multi --push .
```
