FROM golang:1.18-alpine3.16 AS builder

ARG target
ENV target=${target}

WORKDIR /build

ENV GOPROXY https://goproxy.cn
COPY go.mod .
RUN go mod download
RUN go mod tidy

COPY . .
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a -o serve ./applications/${target}/

FROM alpine:3.16 AS doutok-serve

WORKDIR /app
RUN mkdir tmp
COPY --from=builder /build/serve /app
COPY --from=builder /build/config /app/config

ENTRYPOINT ["/app/serve"]
