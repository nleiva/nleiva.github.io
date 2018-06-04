FROM golang:1.10.2 as builder

# Optional WORKDIR (do I need to create the dir?)
RUN mkdir -p /go/src/github.com/nleiva/nleiva.github.io/
WORKDIR /go/src/github.com/nleiva/nleiva.github.io/
COPY ./. .

# Disable crosscompiling 
ENV CGO_ENABLED=0

# Ccompile Linux only
ENV GOOS=linux

# Build the binary with debug information removed
RUN go build -ldflags '-w -s' -a -installsuffix cgo -o web redirect/*.go

FROM alpine
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
COPY --from=builder /go/src/github.com/nleiva/nleiva.github.io/web /

# Copy root CA certificates
COPY --from=builder /go/src/github.com/nleiva/nleiva.github.io/redirect/ca-certificates.crt \
    /etc/ssl/certs/ 

# Copy static web content
COPY --from=builder /go/src/github.com/nleiva/nleiva.github.io/static/redirect.html \
    /static/ 

RUN update-ca-certificates

EXPOSE 80
EXPOSE 443
CMD ["/web"]