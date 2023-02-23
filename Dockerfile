FROM golang:1.18-alpine3.16 AS builder

WORKDIR /build
RUN adduser -u 10001 -D app-runner

ENV GOPROXY https://goproxy.cn
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a -o apiServe ./applications/api/
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a -o commentServe ./applications/comment/
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a -o favoriteServe ./applications/favorite/
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a -o feedServe ./applications/feed/
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a -o messageServe ./applications/message/
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a -o publishServe ./applications/publish/
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a -o relationServe ./applications/relation/
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a -o userServe ./applications/user/

FROM alpine:3.16 AS doutok-api-serve

WORKDIR /app
COPY --from=builder /build/apiServe /app
COPY --from=builder /build/config /app/config

USER app-runner
ENTRYPOINT ["/app/apiServe"]

FROM alpine:3.16 AS doutok-comment-serve

WORKDIR /app
COPY --from=builder /build/commentServe /app
COPY --from=builder /build/config /app/config

USER app-runner
ENTRYPOINT ["/app/commentServe"]

FROM alpine:3.16 AS doutok-favorite-serve

WORKDIR /app
COPY --from=builder /build/favoriteServe /app
COPY --from=builder /build/config /app/config

USER app-runner
ENTRYPOINT ["/app/favoriteServe"]

FROM alpine:3.16 AS doutok-feed-serve

WORKDIR /app
COPY --from=builder /build/feedServe /app
COPY --from=builder /build/config /app/config

USER app-runner
ENTRYPOINT ["/app/feedServe"]

FROM alpine:3.16 AS doutok-message-serve

WORKDIR /app
COPY --from=builder /build/messageServe /app
COPY --from=builder /build/config /app/config

USER app-runner
ENTRYPOINT ["/app/messageServe"]

FROM alpine:3.16 AS doutok-publish-serve

WORKDIR /app
COPY --from=builder /build/publishServe /app
COPY --from=builder /build/config /app/config

USER app-runner
ENTRYPOINT ["/app/publishServe"]

FROM alpine:3.16 AS doutok-relation-serve

WORKDIR /app
COPY --from=builder /build/relationServe /app
COPY --from=builder /build/config /app/config

USER app-runner
ENTRYPOINT ["/app/relationServe"]

FROM alpine:3.16 AS doutok-user-serve

WORKDIR /app
COPY --from=builder /build/userServe /app
COPY --from=builder /build/config /app/config

USER app-runner
ENTRYPOINT ["/app/userServe"]
