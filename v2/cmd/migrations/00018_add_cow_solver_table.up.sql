create table tradelogs_cow_protocol
(
    owner             text default ''::text not null,
    sell_token        text                  not null,
    buy_token         text                  not null,
    sell_amount       text                  not null,
    buy_amount        text                  not null,
    fee_amount        text                  not null,
    order_uid         text                  not null,
    contract_address  text                  not null,
    block_number      bigint                not null,
    tx_hash           text                  not null,
    timestamp         bigint                not null,
    event_hash        text default ''::text not null,
    tx_origin         text default ''::text not null,
    message_sender    text default ''::text not null,
    interact_contract text default ''::text not null,
    log_index         bigint                not null,
    buy_token_price   double precision,
    sell_token_price  double precision,
    buy_usd_amount    double precision,
    sell_usd_amount   double precision,
    primary key (block_number, log_index)
);

create index cow_protocol_timestamp_idx
    on tradelogs_cow_protocol (timestamp DESC);

create index cow_protocol_owner_idx
    on tradelogs_cow_protocol (owner);

create index cow_protocol_order_uid_idx
    on tradelogs_cow_protocol (order_uid);

create index cow_protocol_sell_token_idx
    on tradelogs_cow_protocol (sell_token);

create index cow_protocol_buy_token_idx
    on tradelogs_cow_protocol (buy_token);

create table cow_transfer_event
(
    transfer_id  text   not null,
    tx_hash      text   not null,
    timestamp    bigint not null,
    block_number bigint not null,
    from_address text   not null,
    to_address   text   not null,
    token        text   not null,
    amount       text   not null,
    token_price  double precision,
    amount_usd   double precision,
    primary key (transfer_id)
);

create index cow_transfer_event_timestamp_idx
    on cow_transfer_event (timestamp DESC);

create index cow_transfer_event_tx_hash_idx
    on cow_transfer_event (tx_hash);
