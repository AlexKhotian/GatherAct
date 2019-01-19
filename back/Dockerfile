#FROM golang:1.12-rc-alpine3.8
FROM golang:latest

COPY . /src

WORKDIR /src

RUN go mod vendor

RUN go build -o back

CMD ["/src/back"]