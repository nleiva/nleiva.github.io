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
RUN go build -ldflags '-w -s' -a -installsuffix cgo -o web redirect/main.go

FROM alpine
COPY --from=builder /go/src/github.com/nleiva/nleiva.github.io/web /
EXPOSE 80
CMD ["/web"]