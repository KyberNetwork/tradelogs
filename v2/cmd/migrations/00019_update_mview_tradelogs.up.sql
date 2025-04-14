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