create table backfill (
    exchange text not null primary key,
    deploy_block bigint not null,
    backfilled_block bigint not null default 0
);

insert into backfill (exchange, deploy_block)
values ('zerox', 20926953)