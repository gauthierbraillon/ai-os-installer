FROM golang:1.21-alpine AS builder

WORKDIR /build

COPY go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ai-os-installer ./cmd/ai-os-installer

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /build/ai-os-installer .

ENTRYPOINT ["/app/ai-os-installer"]
