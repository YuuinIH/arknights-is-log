FROM golang:1.17 as builder

##
## Build
##
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN CGO_ENABLED=0 go build -v -o islog

##
## Build
##
FROM alpine:3.14

WORKDIR /root/
COPY --from=builder /app/islog ./islog
COPY --from=builder /app/config ./config

EXPOSE 8080

ENTRYPOINT ["./islog"]
