ALTER TABLE tradelogs_uniswapx
ADD COLUMN expiration bigint,
ADD COLUMN decay_start_time BIGINT DEFAULT NULL,
ADD COLUMN decay_end_time BIGINT DEFAULT NULL,
ADD COLUMN exclusive_filler TEXT NOT NULL DEFAULT '';

ALTER TABLE tradelogs_pancakeswap
ADD COLUMN expiration bigint;

