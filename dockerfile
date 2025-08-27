FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY ./go.mod ./
COPY ./go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o server

FROM alpine:3.18
WORKDIR /app

COPY --from=builder /app/server .

EXPOSE 9999

CMD ["./server"]
