CREATE TABLE IF NOT EXISTS otc_order_filled
(
    order_hash                  TEXT PRIMARY KEY NOT NULL UNIQUE,
    maker                       TEXT NOT NULL,
    taker                       TEXT NOT NULL,
    maker_token                 TEXT NOT NULL,
    taker_token                 TEXT NOT NULL,
    maker_token_filled_amount   TEXT NOT NULL,
    taker_token_filled_amount   TEXT NOT NULL,
    address                     TEXT NOT NULL,
    block_number                BIGINT,
    tx_hash                     TEXT NOT NULL,
    timestamp                   BIGINT
);
CREATE INDEX order_hash_idx ON otc_order_filled (order_hash);
CREATE INDEX timestamp_idx ON otc_order_filled (timestamp);
CREATE INDEX maker_idx ON otc_order_filled (maker);
CREATE INDEX taker_idx ON otc_order_filled (taker); 