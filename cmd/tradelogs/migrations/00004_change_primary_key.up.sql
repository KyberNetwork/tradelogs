ALTER TABLE tradelogs drop constraint tradelogs_pkey;

ALTER table tradelogs add primary key (block_number, log_index, order_hash);