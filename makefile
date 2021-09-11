GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOINSTALL=$(GOCMD) install
GOTEST=$(GOCMD) test
BASE_BINARY_NAME=gorki
BINARY_UNIX=$(BASE_BINARY_NAME)_amd64

# using maven lifecycle terminology here
.PHONY: package clean install deploy __clean __test __build_binary __build_image __push_image 

package: clean __test __build_binary
test: clean __test
clean: __clean
install: package __build_image
deploy: install __push_image

__clean:
	$(GOCLEAN)
	rm -f cmd/gorki/$(BINARY_UNIX)
__test:
	$(GOTEST) -v ./...
__build_binary:
	cd cmd/gorki && env GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v && $(GOINSTALL)
__build_image:
	docker build -t "tastybug/gorki" .
__push_image:
	docker push "tastybug/gorki"
