FROM alpine:3
LABEL maintainer="tastybug@tastybug.com"
RUN mkdir /app
ADD ./cmd/gorki/gorki_amd64 /app/gorki
WORKDIR /app
CMD ["/app/gorki"]