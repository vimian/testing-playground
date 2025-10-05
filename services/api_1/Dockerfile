FROM golang:1.25.0-alpine AS builder

WORKDIR /app

COPY go.* ./

# Must have
COPY test/external_client/go.* ./test/external_client/

# Needed for this service
COPY pkg/types/go.* ./pkg/types/
COPY services/api_1/go.* ./services/api_1/

RUN go work sync
RUN go mod download

COPY pkg/types/ ./pkg/types/
COPY services/api_1/ ./services/api_1/

RUN go test ./services/api_1/...

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main ./services/api_1

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main /app/main

EXPOSE 8080

CMD ["/app/main"]