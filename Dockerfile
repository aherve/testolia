FROM golang:alpine

RUN mkdir /app
WORKDIR /app

ADD ./*.go ./
RUN go build -o testolia

EXPOSE 8080

ENTRYPOINT ["./testolia"]
