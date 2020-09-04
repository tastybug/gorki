GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOINSTALL=$(GOCMD) install
GOTEST=$(GOCMD) test
BASE_BINARY_NAME=gorki
BINARY_UNIX=$(BASE_BINARY_NAME)_amd64

# using maven lifecycle terminology here
clean: __clean
package: clean __test __build_binary
install: package __build_docker_image
deploy: install __push_docker_image_to_hub

__clean:
	$(GOCLEAN)
	rm -f $(BINARY_UNIX)
__test:
	$(GOTEST) -v ./...
__build_binary:
	env GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v
	$(GOINSTALL)
__build_docker_image:
	docker build -t "tastybug/gorki" .
__push_docker_image_to_hub:
	docker push "tastybug/gorki"