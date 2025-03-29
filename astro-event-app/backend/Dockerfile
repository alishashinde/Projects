# Backend Dockerfile
FROM golang:1.20

WORKDIR /app

COPY . .

RUN go mod tidy

CMD ["go", "run", "main.go"]
