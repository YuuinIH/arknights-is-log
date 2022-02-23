FROM golang:1.17 as builder

##
## Build
##
WORKDIR /tmp/build

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /tmp/build/islog

##
## Build
##
FROM alpine:3.14

WORKDIR /
COPY --from=builder /tmp/build/islog /islog
COPY --from=builder /tmp/build/config /config

EXPOSE 8080

ENTRYPOINT [ "/islog" ]
