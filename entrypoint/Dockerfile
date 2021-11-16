# Builder image for iCosts telegram bot on base alpine linux
FROM golang:1.15-alpine3.14 AS builder

ARG GIT_AUTH
ENV GOPRIVATE "github.com"
ENV GOPROXY "http://proxy.golang.org"

## Install git and download dependencies
RUN \
    sed -i 's/https/http/g' /etc/apk/repositories && \
    apk add --no-cache git 

## Copying code to image
ADD ./ /go/src/entrypoint

## Building application inside image
RUN cd /go/src/entrypoint && go build -o entrypoint

# Runtime image for iCosts telegram bot on base alpine linux
FROM alpine:3.14

## Set Moscow time zone and create user for running application
RUN addgroup -S entrypoint && adduser -S -G entrypoint entrypoint

USER entrypoint
WORKDIR /home/entrypoint

## Copying application from temporary image with permissions for user 'icosts'
COPY --from=builder --chown=entrypoint:entrypoint /go/src/entrypoint/entrypoint /usr/bin/entrypoint

ENTRYPOINT [ "entrypoint" ]
CMD [""]