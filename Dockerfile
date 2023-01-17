FROM golang:1.19

WORKDIR /subnet

COPY . .

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]