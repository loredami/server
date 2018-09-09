FROM instrumentisto/dep as dep

FROM golang:latest as builder
WORKDIR /go/src/github.com/loredami/server
COPY . /go/src/github.com/loredami/server
COPY --from=dep /usr/local/bin/dep /usr/local/bin/
RUN dep ensure
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/server .

FROM alpine:latest
WORKDIR /root/
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/server server
ENTRYPOINT ./server