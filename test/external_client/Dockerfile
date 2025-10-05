FROM golang:1.25.0-alpine AS builder

WORKDIR /app

COPY go.* ./

# Must have
COPY services/api_1/go.* ./services/api_1/

# Needed for tests
COPY pkg/types/go.* ./pkg/types/
COPY test/external_client/go.* ./test/external_client/

RUN go work sync
RUN go mod download

COPY pkg/types/ ./pkg/types/
COPY test/external_client ./test/external_client

CMD ["go", "test", "./test/external_client/..."]