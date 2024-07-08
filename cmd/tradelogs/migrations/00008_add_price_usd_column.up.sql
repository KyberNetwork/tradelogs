CREATE TYPE tradelog_states AS ENUM ('new', 'processed');

ALTER TABLE tradelogs
    ADD COLUMN maker_token_price FLOAT NOT NULL DEFAULT 0,
    ADD COLUMN taker_token_price FLOAT NOT NULL DEFAULT 0,
    ADD COLUMN maker_usd_amount FLOAT NOT NULL DEFAULT 0,
    ADD COLUMN taker_usd_amount FLOAT NOT NULL DEFAULT 0,
    ADD COLUMN state tradelog_states NOT NULL DEFAULT 'new';

CREATE INDEX tradelogs_state_idx ON tradelogs (state);
