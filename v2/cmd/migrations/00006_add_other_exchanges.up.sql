-- Zerox
alter table tradelogs_zerox
    add column tx_origin text default ''::text not null;

alter table tradelogs_zerox
    add column maker_token_origin_amount text default ''::text not null;

alter table tradelogs_zerox
    add column taker_token_origin_amount text default ''::text not null;

-- KyberSwap
create table tradelogs_kyberswap
(
    order_hash         text default ''::text not null,
    maker              text default ''::text not null,
    taker              text default ''::text not null,
    maker_token        text                  not null,
    taker_token        text                  not null,
    maker_token_amount text                  not null,
    taker_token_amount text                  not null,
    contract_address   text                  not null,
    block_number       bigint                not null,
    log_index          bigint                not null,
    tx_hash            text                  not null,
    timestamp          bigint                not null,
    event_hash         text default ''::text not null,
    tx_origin          text default ''::text not null,
    message_sender     text default ''::text not null,
    interact_contract  text default ''::text not null,
    maker_token_price  double precision,
    taker_token_price  double precision,
    maker_usd_amount   double precision,
    taker_usd_amount   double precision,
    primary key (block_number, log_index)
);

create index kyberswap_timestamp_idx
    on tradelogs_kyberswap (timestamp);

create index kyberswap_maker_idx
    on tradelogs_kyberswap (maker);

create index kyberswap_taker_idx
    on tradelogs_kyberswap (taker);

create index kyberswap_maker_token_idx
    on tradelogs_kyberswap (maker_token);

create index kyberswap_taker_token_idx
    on tradelogs_kyberswap (taker_token);

-- KyberSwap RFQ
create table tradelogs_kyberswap_rfq
(
    order_hash                text default ''::text not null,
    maker                     text default ''::text not null,
    taker                     text default ''::text not null,
    maker_token               text                  not null,
    taker_token               text                  not null,
    maker_token_amount        text                  not null,
    taker_token_amount        text                  not null,
    maker_token_origin_amount text                  not null,
    taker_token_origin_amount text                  not null,
    contract_address          text                  not null,
    block_number              bigint                not null,
    log_index                 bigint                not null,
    tx_hash                   text                  not null,
    timestamp                 bigint                not null,
    event_hash                text default ''::text not null,
    tx_origin                 text default ''::text not null,
    message_sender            text default ''::text not null,
    interact_contract         text default ''::text not null,
    maker_token_price         double precision,
    taker_token_price         double precision,
    maker_usd_amount          double precision,
    taker_usd_amount          double precision,
    primary key (block_number, log_index)
);

create index kyberswap_rfq_timestamp_idx
    on tradelogs_kyberswap_rfq (timestamp);

create index kyberswap_rfq_maker_idx
    on tradelogs_kyberswap_rfq (maker);

create index kyberswap_rfq_taker_idx
    on tradelogs_kyberswap_rfq (taker);

create index kyberswap_rfq_maker_token_idx
    on tradelogs_kyberswap_rfq (maker_token);

create index kyberswap_rfq_taker_token_idx
    on tradelogs_kyberswap_rfq (taker_token);

-- Paraswap
create table tradelogs_paraswap
(
    order_hash                text default ''::text not null,
    maker                     text default ''::text not null,
    taker                     text default ''::text not null,
    maker_token               text                  not null,
    taker_token               text                  not null,
    maker_token_amount        text                  not null,
    taker_token_amount        text                  not null,
    maker_token_origin_amount text                  not null,
    taker_token_origin_amount text                  not null,
    contract_address          text                  not null,
    block_number              bigint                not null,
    log_index                 bigint                not null,
    tx_hash                   text                  not null,
    timestamp                 bigint                not null,
    event_hash                text default ''::text not null,
    tx_origin                 text default ''::text not null,
    message_sender            text default ''::text not null,
    interact_contract         text default ''::text not null,
    maker_token_price         double precision,
    taker_token_price         double precision,
    maker_usd_amount          double precision,
    taker_usd_amount          double precision,
    primary key (block_number, log_index)
);

create index paraswap_timestamp_idx
    on tradelogs_paraswap (timestamp);

create index paraswap_maker_idx
    on tradelogs_paraswap (maker);

create index paraswap_taker_idx
    on tradelogs_paraswap (taker);

create index paraswap_maker_token_idx
    on tradelogs_paraswap (maker_token);

create index paraswap_taker_token_idx
    on tradelogs_paraswap (taker_token);

-- Hashflow V3
create table tradelogs_hashflow_v3
(
    order_hash                text default ''::text not null,
    maker                     text default ''::text not null,
    taker                     text default ''::text not null,
    maker_token               text                  not null,
    taker_token               text                  not null,
    maker_token_amount        text                  not null,
    taker_token_amount        text                  not null,
    maker_token_origin_amount text                  not null,
    taker_token_origin_amount text                  not null,
    contract_address          text                  not null,
    block_number              bigint                not null,
    log_index                 bigint                not null,
    tx_hash                   text                  not null,
    timestamp                 bigint                not null,
    event_hash                text default ''::text not null,
    tx_origin                 text default ''::text not null,
    message_sender            text default ''::text not null,
    interact_contract         text default ''::text not null,
    maker_token_price         double precision,
    taker_token_price         double precision,
    maker_usd_amount          double precision,
    taker_usd_amount          double precision,
    primary key (block_number, log_index)
);

create index hashflow_v3_timestamp_idx
    on tradelogs_hashflow_v3 (timestamp);

create index hashflow_v3_maker_idx
    on tradelogs_hashflow_v3 (maker);

create index hashflow_v3_taker_idx
    on tradelogs_hashflow_v3 (taker);

create index hashflow_v3_maker_token_idx
    on tradelogs_hashflow_v3 (maker_token);

create index hashflow_v3_taker_token_idx
    on tradelogs_hashflow_v3 (taker_token);

-- 1Inch V6
create table tradelogs_oneinch_v6
(
    order_hash                text default ''::text not null,
    maker                     text default ''::text not null,
    taker                     text default ''::text not null,
    maker_token               text                  not null,
    taker_token               text                  not null,
    maker_token_amount        text                  not null,
    taker_token_amount        text                  not null,
    maker_token_origin_amount text                  not null,
    taker_token_origin_amount text                  not null,
    contract_address          text                  not null,
    block_number              bigint                not null,
    log_index                 bigint                not null,
    tx_hash                   text                  not null,
    timestamp                 bigint                not null,
    event_hash                text default ''::text not null,
    tx_origin                 text default ''::text not null,
    message_sender            text default ''::text not null,
    interact_contract         text default ''::text not null,
    maker_token_price         double precision,
    taker_token_price         double precision,
    maker_usd_amount          double precision,
    taker_usd_amount          double precision,
    maker_traits              text default ''::text not null,
    primary key (block_number, log_index)
);

create index oneinch_v6_timestamp_idx
    on tradelogs_oneinch_v6 (timestamp);

create index oneinch_v6_maker_idx
    on tradelogs_oneinch_v6 (maker);

create index oneinch_v6_taker_idx
    on tradelogs_oneinch_v6 (taker);

create index oneinch_v6_maker_token_idx
    on tradelogs_oneinch_v6 (maker_token);

create index oneinch_v6_taker_token_idx
    on tradelogs_oneinch_v6 (taker_token);

-- Uniswapx
create table tradelogs_uniswapx
(
    order_hash         text default ''::text not null,
    maker              text default ''::text not null,
    taker              text default ''::text not null,
    maker_token        text                  not null,
    taker_token        text                  not null,
    maker_token_amount text                  not null,
    taker_token_amount text                  not null,
    contract_address   text                  not null,
    block_number       bigint                not null,
    log_index          bigint                not null,
    tx_hash            text                  not null,
    timestamp          bigint                not null,
    event_hash         text default ''::text not null,
    tx_origin          text default ''::text not null,
    message_sender     text default ''::text not null,
    interact_contract  text default ''::text not null,
    maker_token_price  double precision,
    taker_token_price  double precision,
    maker_usd_amount   double precision,
    taker_usd_amount   double precision,
    primary key (block_number, log_index)
);

create index uniswapx_timestamp_idx
    on tradelogs_uniswapx (timestamp);

create index uniswapx_maker_idx
    on tradelogs_uniswapx (maker);

create index uniswapx_taker_idx
    on tradelogs_uniswapx (taker);

create index uniswapx_maker_token_idx
    on tradelogs_uniswapx (maker_token);

create index uniswapx_taker_token_idx
    on tradelogs_uniswapx (taker_token);

-- Bebop
create table tradelogs_bebop
(
    order_hash                text default ''::text not null,
    maker                     text default ''::text not null,
    taker                     text default ''::text not null,
    maker_token               text                  not null,
    taker_token               text                  not null,
    maker_token_amount        text                  not null,
    taker_token_amount        text                  not null,
    maker_token_origin_amount text                  not null,
    taker_token_origin_amount text                  not null,
    contract_address          text                  not null,
    block_number              bigint                not null,
    log_index                 bigint                not null,
    trade_index               bigint                not null,
    tx_hash                   text                  not null,
    timestamp                 bigint                not null,
    event_hash                text default ''::text not null,
    tx_origin                 text default ''::text not null,
    message_sender            text default ''::text not null,
    interact_contract         text default ''::text not null,
    maker_token_price         double precision,
    taker_token_price         double precision,
    maker_usd_amount          double precision,
    taker_usd_amount          double precision,
    primary key (block_number, log_index, trade_index)
);

create index bebop_timestamp_idx
    on tradelogs_bebop (timestamp);

create index bebop_maker_idx
    on tradelogs_bebop (maker);

create index bebop_taker_idx
    on tradelogs_bebop (taker);

create index bebop_maker_token_idx
    on tradelogs_bebop (maker_token);

create index bebop_taker_token_idx
    on tradelogs_bebop (taker_token);

-- ZeroX RFQ
create table tradelogs_zerox_v3
(
    order_hash                text default ''::text not null,
    maker                     text default ''::text not null,
    taker                     text default ''::text not null,
    maker_token               text                  not null,
    taker_token               text                  not null,
    maker_token_amount        text                  not null,
    taker_token_amount        text                  not null,
    maker_token_origin_amount text                  not null,
    taker_token_origin_amount text                  not null,
    contract_address          text                  not null,
    block_number              bigint                not null,
    log_index                 bigint                not null,
    tx_hash                   text                  not null,
    timestamp                 bigint                not null,
    event_hash                text default ''::text not null,
    tx_origin                 text default ''::text not null,
    message_sender            text default ''::text not null,
    interact_contract         text default ''::text not null,
    maker_token_price         double precision,
    taker_token_price         double precision,
    maker_usd_amount          double precision,
    taker_usd_amount          double precision,
    primary key (block_number, log_index)
);

create index zerox_v3_timestamp_idx
    on tradelogs_zerox_v3 (timestamp);

create index zerox_v3_maker_idx
    on tradelogs_zerox_v3 (maker);

create index zerox_v3_taker_idx
    on tradelogs_zerox_v3 (taker);

create index zerox_v3_maker_token_idx
    on tradelogs_zerox_v3 (maker_token);

create index zerox_v3_taker_token_idx
    on tradelogs_zerox_v3 (taker_token);

-- PancakeSwap
create table tradelogs_pancakeswap
(
    order_hash         text default ''::text not null,
    maker              text default ''::text not null,
    taker              text default ''::text not null,
    maker_token        text                  not null,
    taker_token        text                  not null,
    maker_token_amount text                  not null,
    taker_token_amount text                  not null,
    contract_address   text                  not null,
    block_number       bigint                not null,
    log_index          bigint                not null,
    tx_hash            text                  not null,
    timestamp          bigint                not null,
    event_hash         text default ''::text not null,
    tx_origin          text default ''::text not null,
    message_sender     text default ''::text not null,
    interact_contract  text default ''::text not null,
    maker_token_price  double precision,
    taker_token_price  double precision,
    maker_usd_amount   double precision,
    taker_usd_amount   double precision,
    primary key (block_number, log_index)
);

create index pancakeswap_timestamp_idx
    on tradelogs_pancakeswap (timestamp);

create index pancakeswap_maker_idx
    on tradelogs_pancakeswap (maker);

create index pancakeswap_taker_idx
    on tradelogs_pancakeswap (taker);

create index pancakeswap_maker_token_idx
    on tradelogs_pancakeswap (maker_token);

create index pancakeswap_taker_token_idx
    on tradelogs_pancakeswap (taker_token);