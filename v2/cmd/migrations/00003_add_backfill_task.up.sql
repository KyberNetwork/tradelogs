CREATE TYPE backfill_status AS ENUM ('processing', 'failed', 'done', 'canceled');

create table backfill_task
(
    id              serial primary key,
    exchange        text                                                  not null,
    from_block      bigint                                                not null,
    to_block        bigint                                                not null,
    processed_block bigint          default 0                             not null,
    created_at      timestamptz     default now()                         not null,
    updated_at      timestamptz     default now()                         not null,
    status          backfill_status default 'processing'::backfill_status not null
);

