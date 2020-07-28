# Bloggo

Golang built very simple, opinionated static site generator.
Open tasks are tracked in the [Todo File](./todo.diff).

### Site Notes

* Fav Icon from https://favicon.io/favicon-generator/, Background Color `#209CEE`, Font `Basic, 110`, Background `Rounded`.
* [Google PageSpeed](https://developers.google.com/speed/pagespeed/insights/?url=https%3A%2F%2Fsleepy-easley-16d3e7.netlify.app/index.html)
* [Writ CSS Reference Sheet](https://writ.cmcenroe.me/reference.html)

### Development

* `go test bloggo` to run all tests
* `go build` builds the executable
* `go run main.go` runs the `main.go` entry point function

To aid development of the site, run a local webserver: 
`docker run -d -p 8080:80 --name local-serve --rm -v $PWD/target:/var/www/html/website tastybug/dockerized-nginx-local-serve nginx`
to quickly check how the site looks like.

### Prepare Docker Image Build
```shell script
env GOOS=linux GOARCH=amd64 go build bloggo
docker build -t "tastybug/bloggo" .
```
To test from within the container, run
`docker run --name bloggo --rm -ti -v $PWD/site:/app/site "tastybug/bloggo" /bin/ash
`

### Run Dockerized Bloggo
```shell script
docker run --name bloggo --rm -v $PWD/site:/app/site "tastybug/bloggo"
```
