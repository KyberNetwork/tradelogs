## BUILDER
FROM golang:1.22-bullseye as builder

WORKDIR /src

COPY . .

RUN go build -o app ./cmd/tradelogs

RUN go build -o parse_log ./v2/cmd/parse_log


## DEPLOY
FROM debian:bullseye

RUN apt-get update && \
    apt install -y ca-certificates && \
    rm -rf /var/lib/apt/lists/*

### tradelogs v1
WORKDIR /cmd

COPY --from=builder /src/app /cmd/app

COPY cmd/tradelogs/migrations migrations

### tradelogs v2
WORKDIR /parse_log
COPY --from=builder /src/parse_log /parse_log/parse_log
COPY v2/cmd/parse_log/migrations migrations

ENTRYPOINT /cmd/app
