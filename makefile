GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
BINARY_NAME=gorki
BINARY_UNIX=$(BINARY_NAME)_amd64

default: clean test compile install
publish: clean  test compile install deploy
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
test:
	$(GOTEST) -v ./...
compile:
	$(GOBUILD) -o $(BINARY_NAME) -v
	env GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v
install:
	docker build -t "tastybug/gorki" .
deploy:
	docker push "tastybug/gorki"
run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)
