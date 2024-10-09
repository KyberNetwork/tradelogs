## BUILDER
FROM golang:1.22-bullseye as builder

WORKDIR /src

COPY . .

RUN go build -o app ./cmd/tradelogs

## DEPLOY
FROM debian:bullseye

RUN apt-get update && \
    apt install -y ca-certificates && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /cmd

COPY --from=builder /src/app /cmd/app

COPY cmd/tradelogs/migrations migrations

CMD /cmd/app
