FROM golang:1.18.1-alpine3.15 as builder
WORKDIR $GOPATH/src
COPY ./ .
RUN go build -v
RUN cp main /

FROM alpine:3.15
COPY --from=builder /main /
CMD ["/main"]
