#### development stage
FROM golang:1.17-buster AS development

# set envvar
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GO111MODULE='on'

# set workdir
WORKDIR /source

# get project dependencies
COPY go.mod go.sum /source/
RUN go mod download

# copy files
COPY . /source

#### builder stage
FROM development AS builder
RUN go build -o ./app ./cmd/main.go

# Run
FROM alpine as run
COPY --from=builder /source/app /
ENTRYPOINT ["/app"]
