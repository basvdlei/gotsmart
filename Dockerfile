FROM golang:1.16-alpine as builder
WORKDIR /go/src/github.com/basvdlei/gotsmart
COPY . .
ENV CGO_ENABLED 0
RUN go vet ./... && \
    go test ./... && \
    go build

FROM alpine:3.13
COPY --from=builder /go/src/github.com/basvdlei/gotsmart/gotsmart \
	/usr/local/bin/gotsmart
EXPOSE 8080
ENTRYPOINT [ "/usr/local/bin/gotsmart" ]
CMD [ "-device", "/dev/ttyS0" ]
