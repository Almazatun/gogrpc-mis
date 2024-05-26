FROM golang:1.22.3-alpine3.18

WORKDIR /app

COPY go.mod .

RUN go mod download

# COPY go.sum .
# COPY .env /app

COPY . .

RUN go build -o main ./cmd

EXPOSE 3055

CMD ["./main"]
