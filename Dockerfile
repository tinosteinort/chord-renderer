FROM golang:1.14.2-alpine AS builder

WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o chord-renderer .


FROM alpine:latest

COPY --from=builder /build/chord-renderer /app/
COPY --from=builder /build/font /app/font

WORKDIR /app
ENTRYPOINT ["./chord-renderer"]
