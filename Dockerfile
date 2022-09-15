FROM golang:1.17-alpine
EXPOSE 8010
RUN mkdir /app
ADD . /app
WORKDIR /app
CMD ["/app/main"]
RUN go build -o main . 