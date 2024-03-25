FROM golang:1.17.6-buster

ADD ./ /rtforum

WORKDIR /rtforum

RUN go build -o app

EXPOSE 8009

CMD ["./app"]