# using maven lifecycle terminology here
.PHONY: package clean install deploy __clean __test __build_binary __build_image __push_image 

test: __test
install: __build_image
deploy: __push_image

__test:
	docker run --rm -v .:/root -w /root golang:1.21 /usr/local/go/bin/go test -v ./...	
__build_image:
	docker buildx build --platform=linux/amd64,linux/arm/v7 -t tastybug/gorki:latest-multi .
__push_image:
	docker buildx build --platform=linux/amd64,linux/arm/v7 -t tastybug/gorki:latest-multi --push .
