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
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux GIN_MODE=release go build -ldflags="-w -s" -o /go/bin/trimbleReader cmd/trimbleReader/main.go

############################
# STEP 2 build a small image
############################

FROM alpine AS senseiveReader
# Copy our static executable.
COPY --from=builder /go/bin/senseiveReader /go/bin/app
ENTRYPOINT ["/go/bin/app"]

FROM alpine AS trimbleReader
# Copy our static executable.
COPY --from=builder /go/bin/trimbleReader /go/bin/app
ENTRYPOINT ["/go/bin/app"]
