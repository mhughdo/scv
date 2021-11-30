# syntax=docker/dockerfile:1

FROM golang:1.17.3-alpine AS builder

ENV CGO_ENABLED=0
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /scv-api ./cmd/api


FROM gcr.io/distroless/base-debian11 as deployer

WORKDIR /

COPY --from=builder /scv-api /scv-api

EXPOSE 4000

USER nonroot:nonroot

ENTRYPOINT ["/scv-api"]