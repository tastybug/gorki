FROM golang:1.14-alpine3.12
LABEL maintainer="tastybug@tastybug.com"
RUN mkdir /app
ADD ./bloggo /app/
WORKDIR /app
CMD ["/app/bloggo"]