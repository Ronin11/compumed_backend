FROM golang:1.22

WORKDIR /api_server

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o api_server .

EXPOSE 8080

CMD ["./api_server"]
