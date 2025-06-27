FROM golang:1.24.4-bookworm AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o server ./server/

FROM debian:bookworm

COPY --from=builder /app/server /

ENTRYPOINT ["/server"]
