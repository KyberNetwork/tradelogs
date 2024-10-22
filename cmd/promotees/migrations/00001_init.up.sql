CREATE TABLE IF NOT EXISTS promoteesName (
    promoter TEXT NOT NULL,
    name TEXT NOT NULL,
    PRIMARY KEY(promoter)
);

CREATE TABLE IF NOT EXISTS promotees (
    promoter    TEXT NOT NULL,
	promotee    TEXT NOT NULL,
	timestamp   BIGINT, 
	event_hash   TEXT,
	chain_id     TEXT NOT NULL,
	block_number BIGINT,
    PRIMARY KEY(promoter, promotee, chain_id)
);

CREATE INDEX block_number_idx ON promotees (block_number);