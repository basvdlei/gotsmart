FROM docker.io/library/golang:1.18-alpine as builder
WORKDIR /usr/src/app
COPY . .
ENV CGO_ENABLED 0
RUN go vet ./... && \
    go test ./... && \
    go build -o gotsmart

FROM alpine:3.13
COPY --from=builder /usr/src/app/gotsmart \
	/usr/local/bin/gotsmart
EXPOSE 8080
ENTRYPOINT [ "/usr/local/bin/gotsmart" ]
CMD [ "-device", "/dev/ttyS0" ]
