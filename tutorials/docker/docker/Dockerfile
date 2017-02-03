# Grab image from Docker hub
FROM golang:1.7

MAINTAINER wchan@freewheel.tv

RUN mkdir example
COPY ./main.go example
WORKDIR example
RUN go build .

ENTRYPOINT ["./example"]
