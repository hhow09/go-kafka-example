FROM golang:1.19-alpine3.17
WORKDIR /app
COPY . .
RUN go build -o log-parser

CMD ["/app/log-parser"]