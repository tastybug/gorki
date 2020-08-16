GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
BASE_BINARY_NAME=gorki
BINARY_UNIX=$(BASE_BINARY_NAME)_amd64

default: clean test compile install
push: clean  test compile install docker-push
clean:
	$(GOCLEAN)
	rm -f $(BINARY_UNIX)
test:
	$(GOTEST) -v ./...
compile:
	env GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v
install:
	docker build -t "tastybug/gorki" .
docker-push:
	docker push "tastybug/gorki"