############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder
# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git build-base
WORKDIR $GOPATH/src/mypackage/myapp/
COPY . .
# Build the binary.
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux GIN_MODE=release go build -ldflags="-w -s" -o /go/bin/senseiveReader cmd/senseiveReader/main.go

############################
# STEP 2 build a small image
############################

FROM alpine AS seseiveReader
# Copy our static executable.
COPY --from=builder /go/bin/senseiveReader /go/bin/app
RUN mkdir "backup"
ENTRYPOINT ["/go/bin/app"]
