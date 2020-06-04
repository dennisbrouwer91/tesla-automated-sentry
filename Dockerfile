FROM golang:onbuild as builder
WORKDIR /build
ADD . .
RUN go get github.com/jsgoecke/tesla \
    && CGO_ENABLED=0 go build -ldflags='-w -s -extldflags "-static"' -a -o /build/main .

FROM scratch
COPY --from=builder /build/ /go/bin/
ENTRYPOINT ["/go/bin/main"]