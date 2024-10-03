## BUILDER
FROM golang:1.22-bullseye as builder

WORKDIR /src

COPY . .

RUN go build -o app ./cmd/tradelogs

RUN go build -o parse_log ./v2/cmd/parse_log
RUN go build -o backfill ./v2/cmd/backfill


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
WORKDIR /v2
COPY --from=builder /src/parse_log /v2/parse_log
COPY --from=builder /src/backfill /v2/backfill
COPY v2/cmd/migrations /v2/migrations

CMD /cmd/app
