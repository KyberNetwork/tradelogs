ALTER TABLE tradelogs_uniswapx
ADD COLUMN expiration BIGINT NOT NULL DEFAULT 0,
ADD COLUMN decay_start_time BIGINT DEFAULT NULL,
ADD COLUMN decay_end_time BIGINT DEFAULT NULL,
ADD COLUMN exclusive_filler TEXT NOT NULL DEFAULT '';

ALTER TABLE tradelogs_pancakeswap
ADD COLUMN expiration BIGINT NOT NULL DEFAULT 0;

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
      maker_token_price,
      taker_token_price,
      maker_usd_amount,
      taker_usd_amount,
      'zerox' AS contract,
      '' AS type
  FROM public.tradelogs_zerox
  WHERE timestamp >= EXTRACT(EPOCH FROM CURRENT_TIMESTAMP - INTERVAL '3 months') * 1000
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
      maker_token_price,
      taker_token_price,
      maker_usd_amount,
      taker_usd_amount,
      'kyberswap' AS contract,
      '' AS type
  FROM public.tradelogs_kyberswap 
  WHERE timestamp >= EXTRACT(EPOCH FROM CURRENT_TIMESTAMP - INTERVAL '3 months') * 1000
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
      maker_token_price,
      taker_token_price,
      maker_usd_amount,
      taker_usd_amount,
      'kyberswap_rfq' AS contract,
      '' AS type
  FROM public.tradelogs_kyberswap_rfq 
  WHERE timestamp >= EXTRACT(EPOCH FROM CURRENT_TIMESTAMP - INTERVAL '3 months') * 1000
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
      maker_token_price,
      taker_token_price,
      maker_usd_amount,
      taker_usd_amount,
      'paraswap' AS contract,
      '' AS type
  FROM public.tradelogs_paraswap 
  WHERE timestamp >= EXTRACT(EPOCH FROM CURRENT_TIMESTAMP - INTERVAL '3 months') * 1000
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
      maker_token_price,
      taker_token_price,
      maker_usd_amount,
      taker_usd_amount,
      'hashflow_v3' AS contract,
      '' AS type
  FROM public.tradelogs_hashflow_v3 
  WHERE timestamp >= EXTRACT(EPOCH FROM CURRENT_TIMESTAMP - INTERVAL '3 months') * 1000
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
      maker_token_price,
      taker_token_price,
      maker_usd_amount,
      taker_usd_amount,
      'oneinch_v6' AS contract,
      type
  FROM public.tradelogs_oneinch_v6 
  WHERE timestamp >= EXTRACT(EPOCH FROM CURRENT_TIMESTAMP - INTERVAL '3 months') * 1000
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
      maker_token_price,
      taker_token_price,
      maker_usd_amount,
      taker_usd_amount,
      'uniswapx' AS contract,
      '' AS type
  FROM public.tradelogs_uniswapx 
  WHERE timestamp >= EXTRACT(EPOCH FROM CURRENT_TIMESTAMP - INTERVAL '3 months') * 1000
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
      maker_token_price,
      taker_token_price,
      maker_usd_amount,
      taker_usd_amount,
      'bebop' AS contract,
      '' AS type
  FROM public.tradelogs_bebop 
  WHERE timestamp >= EXTRACT(EPOCH FROM CURRENT_TIMESTAMP - INTERVAL '3 months') * 1000
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
      maker_token_price,
      taker_token_price,
      maker_usd_amount,
      taker_usd_amount,
      'zerox_v3' AS contract,
      '' AS type
  FROM public.tradelogs_zerox_v3
  WHERE timestamp >= EXTRACT(EPOCH FROM CURRENT_TIMESTAMP - INTERVAL '3 months') * 1000
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
      maker_token_price,
      taker_token_price,
      maker_usd_amount,
      taker_usd_amount,
      'pancakeswap' AS contract,
      '' AS type
  FROM public.tradelogs_pancakeswap 
  WHERE timestamp >= EXTRACT(EPOCH FROM CURRENT_TIMESTAMP - INTERVAL '3 months') * 1000
  AND maker_usd_amount > 0 AND taker_usd_amount > 0
)
SELECT * FROM all_trades ORDER BY timestamp DESC;
