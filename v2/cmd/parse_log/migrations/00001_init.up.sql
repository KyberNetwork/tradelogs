CREATE TYPE tradelog_states AS ENUM ('new', 'processed');

create table tradelogs_zerox
(
    order_hash         text             default ''::text               not null,
    maker              text             default ''::text               not null,
    taker              text             default ''::text               not null,
    maker_token        text                                            not null,
    taker_token        text                                            not null,
    maker_token_amount text                                            not null,
    taker_token_amount text                                            not null,
    contract_address   text                                            not null,
    block_number       bigint                                          not null,
    log_index          bigint                                          not null,
    tx_hash            text                                            not null,
    timestamp          bigint                                          not null,
    event_hash         text             default ''::text               not null,
    maker_token_price  double precision default 0                      not null,
    taker_token_price  double precision default 0                      not null,
    maker_usd_amount   double precision default 0                      not null,
    taker_usd_amount   double precision default 0                      not null,
    state              tradelog_states  default 'new'::tradelog_states not null,
    primary key (block_number, log_index)
);

create index zerox_timestamp_idx
    on tradelogs_zerox (timestamp);

create index zerox_maker_idx
    on tradelogs_zerox (maker);

create index zerox_taker_idx
    on tradelogs_zerox (taker);

create index zerox_maker_token_idx
    on tradelogs_zerox (maker_token);

create index zerox_taker_token_idx
    on tradelogs_zerox (taker_token);

create index zerox_state_idx
    on tradelogs_zerox (state);

create table tradelogs_state
(
    key   text not null constraint tradelogs_state_pk primary key,
    value text not null
);