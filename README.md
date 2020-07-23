# Bloggo

Golang built very simple, opinionated static site generator.
Open tasks are tracked in the [Todo File](./todo.diff).

### Site Notes

* Fav Icon from https://favicon.io/favicon-generator/, Background Color `#209CEE`, Font `Basic, 110`, Background `Rounded`.

### Building, Running

* `go test bloggo` to run all tests
* `go build` builds the executable
* `go run main.go` runs the `main.go` entry point function

To aid development of the site, run a local webserver: 
`docker run -d -p 8080:80 --name local-serve --rm -v $PWD/target:/var/www/html/website tastybug/dockerized-nginx-local-serve nginx`
to quickly check how the site looks like.

