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
RUN go build -mod=vendor .

# fetch frontend release
RUN wget https://github.com/Soontao/PDISolutionCenterFront/releases/download/v1.1.1/dist.zip
RUN unzip dist.zip
RUN cp -R dist/ static/ 

# distribution image
FROM alpine:3.9

# add CAs
RUN apk --no-cache add ca-certificates

COPY --from=build-env /app/app .
COPY --from=build-env /app/app/static ./static

# start
CMD ["./app"]