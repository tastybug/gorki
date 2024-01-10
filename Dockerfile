FROM alpine:3
LABEL maintainer="tastybug@tastybug.com"

ARG gorki_arch
RUN mkdir /app
ADD "./cmd/gorki/gorki_${gorki_arch}" /app/gorki
WORKDIR /app
CMD ["/app/gorki"]
