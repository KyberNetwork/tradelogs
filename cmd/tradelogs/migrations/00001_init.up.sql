CREATE TABLE IF NOT EXISTS tradelogs
(
    order_hash          TEXT NOT NULL UNIQUE,
    maker               TEXT,
    taker               TEXT,
    maker_token         TEXT NOT NULL,
    taker_token         TEXT NOT NULL,
    maker_token_amount  TEXT NOT NULL,
    taker_token_amount  TEXT NOT NULL,
    contract_address    TEXT NOT NULL,
    block_number        BIGINT NOT NULL,
    log_index           BIGINT NOT NULL,
    tx_hash             TEXT NOT NULL,
    timestamp           BIGINT NOT NULL,
    PRIMARY KEY(block_number, log_index)
);

CREATE INDEX timestamp_idx ON tradelogs (timestamp);
CREATE INDEX maker_idx ON tradelogs (maker);
CREATE INDEX taker_idx ON tradelogs (taker); 
CREATE INDEX maker_token_idx ON tradelogs (maker_token);
CREATE INDEX taker_token_idx ON tradelogs (taker_token);