[Task Tracker](./todo.diff)

# Gorki

Gorki is a simple, opinionated static site generator written in Go. It allows you to generate a blog from Markdown written articles.

### Development

* `make`: Runs tests, builds the default binary.
* `make all`: Runs tests, builds for local env and amd64, creates and pushes docker image.

To aid development of the site, run a local webserver: 
```shell script
docker run -d -p 8080:80 --name local-serve --rm -v $PWD/target:/var/www/html/website tastybug/dockerized-nginx-local-serve nginx
```
to quickly check how the site looks like.

### Testing the Dockerized Gorki
To test from within the container, run
```shell script
docker run --name gorki --rm -ti -v $PWD/site:/app/site "tastybug/gorki" /bin/ash
```

### Run Dockerized Gorki
```shell script
docker run --name gorki --rm -v $PWD/site:/app/site "tastybug/gorki"
```
