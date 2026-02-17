FROM golang:1.22-alpine@sha256:1699c10032ca2582ec89a24a1312d986a3f094aed3d5c1147b19880afe40e052 AS builder

WORKDIR /build

COPY go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ai-os-installer ./cmd/ai-os-installer

FROM alpine:3.21@sha256:c3f8e73fdb79deaebaa2037150150191b9dcbfba68b4a46d70103204c53f4709

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /build/ai-os-installer .

ENTRYPOINT ["/app/ai-os-installer"]
