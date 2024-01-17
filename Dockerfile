FROM --platform=$BUILDPLATFORM golang:1.21 as build
WORKDIR /root

COPY . /root

RUN go test -v ./...
RUN cd cmd/gorki; go build -o gorki -v

FROM alpine:3
LABEL maintainer="tastybug@tastybug.com"

RUN mkdir /app
COPY --from=build /root/cmd/gorki/gorki /app/gorki
WORKDIR /app
CMD ["/app/gorki"]
