CREATE TABLE errorlogs (
    address text NOT NULL,
    topics text NOT NULL,
    data bytea NOT NULL,
    block_number int8 NOT NULL,
    tx_hash text NOT NULL,
    tx_index int8 NOT NULL,
    block_hash text NOT NULL,
    log_index int8 NOT NULL,
    time int8 NOT NULL,
    CONSTRAINT errorlogs_pkey PRIMARY KEY (block_number, log_index)
);
