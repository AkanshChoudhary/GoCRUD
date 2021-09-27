FROM golang:alpine

LABEL maintainer="Quique <akanshchoudhary79@gmail.com>"

WORKDIR /app

COPY go.mod .

COPY go.sum .

COPY . .

ENV PORT 8000

RUN go build

CMD ["./go_crud"]

