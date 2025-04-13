create table tradelogs_cow_protocol
(
    owner             text default ''::text not null,
    sell_token        text                  not null,
    buy_token         text                  not null,
    sell_amount       text                  not null,
    buy_amount        text                  not null,
    fee_amount        text                  not null,
    order_uid         text                  not null,
    raw_trade_data    text                  not null,
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
    tx_hash      text   not null,
    timestamp    bigint not null,
    block_number bigint not null,
    from_address text   not null,
    to_address   text   not null,
    token        text   not null,
    amount       text   not null,
    token_price  double precision,
    amount_usd   double precision,
    primary key (tx_hash,timestamp,from_address,to_address,token,amount)
);

create index cow_transfer_event_timestamp_idx
    on cow_transfer_event (timestamp DESC);

create index cow_transfer_event_tx_hash_idx
    on cow_transfer_event (tx_hash);

DROP MATERIALIZED VIEW IF EXISTS mview_tradelogs;
CREATE MATERIALIZED VIEW mview_tradelogs AS
WITH all_trades AS (
  SELECT
      maker,
      taker,
      maker_token,
      taker_token,
      CASE 
        WHEN maker_token_amount ~ '^[0-9]+(\.[0-9]+)?$' THEN REPLACE(maker_token_amount, ',', '.')::DECIMAL
        ELSE NULL
      END AS maker_token_amount,
      CASE 
          WHEN taker_token_amount ~ '^[0-9]+(\.[0-9]+)?$' THEN REPLACE(taker_token_amount, ',', '.')::DECIMAL
          ELSE NULL
      END AS taker_token_amount,
      contract_address,
      block_number,
      tx_hash,
      timestamp,
      event_hash,
      tx_origin,
      maker_token_price,
      taker_token_price,
      maker_usd_amount,
      taker_usd_amount,
      CASE 
        WHEN maker_usd_amount IS NULL THEN NULL 
        WHEN maker_usd_amount > 0 THEN maker_usd_amount
        ELSE taker_usd_amount
      END AS usd_volume,
      'zerox' AS contract,
      '' AS type
  FROM public.tradelogs_zerox
  WHERE timestamp >= EXTRACT(EPOCH FROM CURRENT_TIMESTAMP - INTERVAL '1 day') * 1000
  AND maker_usd_amount > 0 AND taker_usd_amount > 0

  
  UNION ALL
  
  SELECT
      maker,
      taker,
      maker_token,
      taker_token,
      CASE 
        WHEN maker_token_amount ~ '^[0-9]+(\.[0-9]+)?$' THEN REPLACE(maker_token_amount, ',', '.')::DECIMAL
        ELSE NULL
      END AS maker_token_amount,
      CASE 
          WHEN taker_token_amount ~ '^[0-9]+(\.[0-9]+)?$' THEN REPLACE(taker_token_amount, ',', '.')::DECIMAL
          ELSE NULL
      END AS taker_token_amount,
      contract_address,
      block_number,
      tx_hash,
      timestamp,
      event_hash,
      tx_origin,
      maker_token_price,
      taker_token_price,
      maker_usd_amount,
      taker_usd_amount,
      CASE 
        WHEN maker_usd_amount IS NULL THEN NULL 
        WHEN maker_usd_amount > 0 THEN maker_usd_amount
        ELSE taker_usd_amount
      END AS usd_volume,
      'kyberswap' AS contract,
      '' AS type
  FROM public.tradelogs_kyberswap 
  WHERE timestamp >= EXTRACT(EPOCH FROM CURRENT_TIMESTAMP - INTERVAL '1 day') * 1000
  AND maker_usd_amount > 0 AND taker_usd_amount > 0
  
  UNION ALL
  
  SELECT
      maker,
      taker,
      maker_token,
      taker_token,
      CASE 
        WHEN maker_token_amount ~ '^[0-9]+(\.[0-9]+)?$' THEN REPLACE(maker_token_amount, ',', '.')::DECIMAL
        ELSE NULL
      END AS maker_token_amount,
      CASE 
          WHEN taker_token_amount ~ '^[0-9]+(\.[0-9]+)?$' THEN REPLACE(taker_token_amount, ',', '.')::DECIMAL
          ELSE NULL
      END AS taker_token_amount,
      contract_address,
      block_number,
      tx_hash,
      timestamp,
      event_hash,
      tx_origin,
      maker_token_price,
      taker_token_price,
      maker_usd_amount,
      taker_usd_amount,
      CASE 
        WHEN maker_usd_amount IS NULL THEN NULL 
        WHEN maker_usd_amount > 0 THEN maker_usd_amount
        ELSE taker_usd_amount
      END AS usd_volume,
      'kyberswap_rfq' AS contract,
      '' AS type
  FROM public.tradelogs_kyberswap_rfq 
  WHERE timestamp >= EXTRACT(EPOCH FROM CURRENT_TIMESTAMP - INTERVAL '1 day') * 1000
  AND maker_usd_amount > 0 AND taker_usd_amount > 0
  
  UNION ALL
  
  SELECT
      maker,
      taker,
      maker_token,
      taker_token,
      CASE 
        WHEN maker_token_amount ~ '^[0-9]+(\.[0-9]+)?$' THEN REPLACE(maker_token_amount, ',', '.')::DECIMAL
        ELSE NULL
      END AS maker_token_amount,
      CASE 
          WHEN taker_token_amount ~ '^[0-9]+(\.[0-9]+)?$' THEN REPLACE(taker_token_amount, ',', '.')::DECIMAL
          ELSE NULL
      END AS taker_token_amount,
      contract_address,
      block_number,
      tx_hash,
      timestamp,
      event_hash,
      tx_origin,
      maker_token_price,
      taker_token_price,
      maker_usd_amount,
      taker_usd_amount,
      CASE 
        WHEN maker_usd_amount IS NULL THEN NULL 
        WHEN maker_usd_amount > 0 THEN maker_usd_amount
        ELSE taker_usd_amount
      END AS usd_volume,
      'paraswap' AS contract,
      '' AS type
  FROM public.tradelogs_paraswap 
  WHERE timestamp >= EXTRACT(EPOCH FROM CURRENT_TIMESTAMP - INTERVAL '1 day') * 1000
  AND maker_usd_amount > 0 AND taker_usd_amount > 0
  
  UNION ALL
  
  SELECT
      maker,
      taker,
      maker_token,
      taker_token,
      CASE 
        WHEN maker_token_amount ~ '^[0-9]+(\.[0-9]+)?$' THEN REPLACE(maker_token_amount, ',', '.')::DECIMAL
        ELSE NULL
      END AS maker_token_amount,
      CASE 
          WHEN taker_token_amount ~ '^[0-9]+(\.[0-9]+)?$' THEN REPLACE(taker_token_amount, ',', '.')::DECIMAL
          ELSE NULL
      END AS taker_token_amount,
      contract_address,
      block_number,
      tx_hash,
      timestamp,
      event_hash,
      tx_origin,
      maker_token_price,
      taker_token_price,
      maker_usd_amount,
      taker_usd_amount,
      CASE 
        WHEN maker_usd_amount IS NULL THEN NULL 
        WHEN maker_usd_amount > 0 THEN maker_usd_amount
        ELSE taker_usd_amount
      END AS usd_volume,
      'hashflow_v3' AS contract,
      '' AS type
  FROM public.tradelogs_hashflow_v3 
  WHERE timestamp >= EXTRACT(EPOCH FROM CURRENT_TIMESTAMP - INTERVAL '1 day') * 1000
  AND maker_usd_amount > 0 AND taker_usd_amount > 0
  
  UNION ALL
  
  SELECT
      maker,
      taker,
      maker_token,
      taker_token,
      CASE 
        WHEN maker_token_amount ~ '^[0-9]+(\.[0-9]+)?$' THEN REPLACE(maker_token_amount, ',', '.')::DECIMAL
        ELSE NULL
      END AS maker_token_amount,
      CASE 
          WHEN taker_token_amount ~ '^[0-9]+(\.[0-9]+)?$' THEN REPLACE(taker_token_amount, ',', '.')::DECIMAL
          ELSE NULL
      END AS taker_token_amount,
      contract_address,
      block_number,
      tx_hash,
      timestamp,
      event_hash,
      tx_origin,
      maker_token_price,
      taker_token_price,
      maker_usd_amount,
      taker_usd_amount,
      CASE 
        WHEN maker_usd_amount IS NULL THEN NULL 
        WHEN maker_usd_amount > 0 THEN maker_usd_amount
        ELSE taker_usd_amount
      END AS usd_volume,
      'oneinch_v6' AS contract,
      type
  FROM public.tradelogs_oneinch_v6 
  WHERE timestamp >= EXTRACT(EPOCH FROM CURRENT_TIMESTAMP - INTERVAL '1 day') * 1000
  AND maker_usd_amount > 0 AND taker_usd_amount > 0
  
  UNION ALL
  
  SELECT
      maker,
      taker,
      maker_token,
      taker_token,
      CASE 
        WHEN maker_token_amount ~ '^[0-9]+(\.[0-9]+)?$' THEN REPLACE(maker_token_amount, ',', '.')::DECIMAL
        ELSE NULL
      END AS maker_token_amount,
      CASE 
          WHEN taker_token_amount ~ '^[0-9]+(\.[0-9]+)?$' THEN REPLACE(taker_token_amount, ',', '.')::DECIMAL
          ELSE NULL
      END AS taker_token_amount,
      contract_address,
      block_number,
      tx_hash,
      timestamp,
      event_hash,
      tx_origin,
      maker_token_price,
      taker_token_price,
      maker_usd_amount,
      taker_usd_amount,
      CASE 
        WHEN maker_usd_amount IS NULL THEN NULL 
        WHEN maker_usd_amount > 0 THEN maker_usd_amount
        ELSE taker_usd_amount
      END AS usd_volume,
      'uniswapx' AS contract,
      '' AS type
  FROM public.tradelogs_uniswapx 
  WHERE timestamp >= EXTRACT(EPOCH FROM CURRENT_TIMESTAMP - INTERVAL '1 day') * 1000
  AND maker_usd_amount > 0 AND taker_usd_amount > 0
  
  UNION ALL
  
  SELECT
      maker,
      taker,
      maker_token,
      taker_token,
      CASE 
        WHEN maker_token_amount ~ '^[0-9]+(\.[0-9]+)?$' THEN REPLACE(maker_token_amount, ',', '.')::DECIMAL
        ELSE NULL
      END AS maker_token_amount,
      CASE 
          WHEN taker_token_amount ~ '^[0-9]+(\.[0-9]+)?$' THEN REPLACE(taker_token_amount, ',', '.')::DECIMAL
          ELSE NULL
      END AS taker_token_amount,
      contract_address,
      block_number,
      tx_hash,
      timestamp,
      event_hash,
      tx_origin,
      maker_token_price,
      taker_token_price,
      maker_usd_amount,
      taker_usd_amount,
      CASE 
        WHEN maker_usd_amount IS NULL THEN NULL 
        WHEN maker_usd_amount > 0 THEN maker_usd_amount
        ELSE taker_usd_amount
      END AS usd_volume,
      'bebop' AS contract,
      '' AS type
  FROM public.tradelogs_bebop 
  WHERE timestamp >= EXTRACT(EPOCH FROM CURRENT_TIMESTAMP - INTERVAL '1 day') * 1000
  AND maker_usd_amount > 0 AND taker_usd_amount > 0
  
  UNION ALL
  
  SELECT
      maker,
      taker,
      maker_token,
      taker_token,
      CASE 
        WHEN maker_token_amount ~ '^[0-9]+(\.[0-9]+)?$' THEN REPLACE(maker_token_amount, ',', '.')::DECIMAL
        ELSE NULL
      END AS maker_token_amount,
      CASE 
          WHEN taker_token_amount ~ '^[0-9]+(\.[0-9]+)?$' THEN REPLACE(taker_token_amount, ',', '.')::DECIMAL
          ELSE NULL
      END AS taker_token_amount,
      contract_address,
      block_number,
      tx_hash,
      timestamp,
      event_hash,
      tx_origin,
      maker_token_price,
      taker_token_price,
      maker_usd_amount,
      taker_usd_amount,
      CASE 
        WHEN maker_usd_amount IS NULL THEN NULL 
        WHEN maker_usd_amount > 0 THEN maker_usd_amount
        ELSE taker_usd_amount
      END AS usd_volume,
      'zerox_v3' AS contract,
      '' AS type
  FROM public.tradelogs_zerox_v3
  WHERE timestamp >= EXTRACT(EPOCH FROM CURRENT_TIMESTAMP - INTERVAL '1 day') * 1000
  AND maker_usd_amount > 0 AND taker_usd_amount > 0
  
  UNION ALL
  
  SELECT
      maker,
      taker,
      maker_token,
      taker_token,
      CASE 
        WHEN maker_token_amount ~ '^[0-9]+(\.[0-9]+)?$' THEN REPLACE(maker_token_amount, ',', '.')::DECIMAL
        ELSE NULL
      END AS maker_token_amount,
      CASE 
          WHEN taker_token_amount ~ '^[0-9]+(\.[0-9]+)?$' THEN REPLACE(taker_token_amount, ',', '.')::DECIMAL
          ELSE NULL
      END AS taker_token_amount,
      contract_address,
      block_number,
      tx_hash,
      timestamp,
      event_hash,
      tx_origin,
      maker_token_price,
      taker_token_price,
      maker_usd_amount,
      taker_usd_amount,
      CASE 
        WHEN maker_usd_amount IS NULL THEN NULL 
        WHEN maker_usd_amount > 0 THEN maker_usd_amount
        ELSE taker_usd_amount
      END AS usd_volume,
      'pancakeswap' AS contract,
      '' AS type
  FROM public.tradelogs_pancakeswap 
  WHERE timestamp >= EXTRACT(EPOCH FROM CURRENT_TIMESTAMP - INTERVAL '1 day') * 1000
  AND maker_usd_amount > 0 AND taker_usd_amount > 0
)
SELECT * FROM all_trades ORDER BY timestamp DESC;