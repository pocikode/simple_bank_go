# Build stage
FROM golang:1.22.4-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

# Run stage
FROM alpine:3.20
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env.example ./app.env
COPY start.sh .
COPY wait-for.sh .
COPY db/migration ./db/migration

EXPOSE 8012
CMD ["/app/main"]
ENTRYPOINT ["/app/start.sh"]