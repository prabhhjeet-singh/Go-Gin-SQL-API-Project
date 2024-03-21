FROM golang:alpine

ENV CGO_ENABLED=1

# Install build dependencies
RUN apk add --no-cache gcc musl-dev

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

WORKDIR /app/rest-api

RUN go build -o main .

EXPOSE 8081

CMD ["./main"]
