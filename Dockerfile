FROM golang:1.24.2

WORKDIR /app

ENV GOPROXY=https://goproxy.cn,direct

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o main ./cmd/main.go

CMD ["/app/main"]
