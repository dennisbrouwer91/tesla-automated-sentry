FROM golang:onbuild
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go get github.com/jsgoecke/tesla
RUN go build -o main .
CMD ["/app/main --daemon true"]