FROM golang:1.15.6 AS builder

RUN mkdir /tmp/go
WORKDIR /tmp/go
COPY main.go go.mod ./

RUN CGO_ENABLED=0 GOOS=linux go build -o demo-webserver .

FROM centos:8

ARG WEBSERVER_VERSION=v1.0
ENV DEMO_WEBSERVER_VERSION=${WEBSERVER_VERSION}

COPY --from=builder /tmp/go/demo-webserver /usr/local/bin/
RUN chmod +x /usr/local/bin/demo-webserver
CMD ["/usr/local/bin/demo-webserver"]
