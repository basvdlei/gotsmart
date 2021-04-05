FROM golang:1.16-alpine as builder
RUN apk update && apk add git
COPY . /go/src/github.com/robbertnoordzij/gotsmart
WORKDIR /go/src/github.com/robbertnoordzij/gotsmart
ENV CGO_ENABLED 0
RUN go get ./...
RUN go vet ./... && \
    go test ./... && \
    go build

FROM alpine:3.8
COPY --from=builder /go/src/github.com/robbertnoordzij/gotsmart/gotsmart \
	/usr/local/bin/gotsmart
EXPOSE 8080
ENTRYPOINT [ "/usr/local/bin/gotsmart" ]
CMD [ "-device", "/dev/ttyS0" ]
