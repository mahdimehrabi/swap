FROM golang:1.22-alpine AS base
RUN apk add build-base
RUN echo $GOPATH

copy . /apps


RUN go build /app/cmd/main.go

CMD /app/main
