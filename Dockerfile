FROM golang:1.18-alpine AS builder
WORKDIR /app/clean-architecture
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN go build -o clean-architecture clean-architecture/cmd

FROM alpine
RUN apk update && apk add --no-cache tzdata
WORKDIR /app
COPY --from=builder /app/clean-architecture /app/
ENTRYPOINT ["./clean-architecture"]