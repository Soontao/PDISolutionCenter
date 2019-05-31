# build image
FROM golang:1.12-alpine3.9 AS build-env

# install build tools
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh wget unzip gcc libc-dev

# build
WORKDIR /app

# copy sources
COPY . .

# vendor build only can be executed outside the GOPATH
RUN go build -mod=vendor -o=app .

# fetch frontend release
RUN wget https://github.com/Soontao/PDISolutionCenterFront/releases/download/v1.5.0/dist.zip
RUN unzip dist.zip

# distribution image
FROM alpine:3.9

# add CAs
RUN apk --no-cache add ca-certificates

COPY --from=build-env /app/app .
COPY --from=build-env /app/dist ./static

# start
CMD ["./app"]