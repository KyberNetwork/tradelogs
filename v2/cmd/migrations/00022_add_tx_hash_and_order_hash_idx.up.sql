create index tradelogs_oneinch_v6_tx_hash_idx
    on tradelogs_oneinch_v6 USING hash (tx_hash);
create index tradelogs_zerox_tx_hash_idx
    on tradelogs_zerox USING hash (tx_hash);
create index tradelogs_kyberswap_tx_hash_idx
    on tradelogs_kyberswap USING hash (tx_hash);
create index tradelogs_kyberswap_rfq_tx_hash_idx
    on tradelogs_kyberswap_rfq USING hash (tx_hash);
create index tradelogs_paraswap_tx_hash_idx
    on tradelogs_paraswap USING hash (tx_hash);
create index tradelogs_hashflow_v3_tx_hash_idx
    on tradelogs_hashflow_v3 USING hash (tx_hash);
create index tradelogs_uniswapx_tx_hash_idx
    on tradelogs_uniswapx USING hash (tx_hash);
create index tradelogs_bebop_tx_hash_idx
    on tradelogs_bebop USING hash (tx_hash);
create index tradelogs_zerox_v3_tx_hash_idx
    on tradelogs_zerox_v3 USING hash (tx_hash);
create index tradelogs_pancakeswap_tx_hash_idx
    on tradelogs_pancakeswap USING hash (tx_hash);

create index tradelogs_zerox_order_hash_idx
    on tradelogs_zerox USING hash (order_hash);
create index tradelogs_kyberswap_order_hash_idx
    on tradelogs_kyberswap USING hash (order_hash);
create index tradelogs_kyberswap_rfq_order_hash_idx
    on tradelogs_kyberswap_rfq USING hash (order_hash);
create index tradelogs_paraswap_order_hash_idx
    on tradelogs_paraswap USING hash (order_hash);
create index tradelogs_hashflow_v3_order_hash_idx
    on tradelogs_hashflow_v3 USING hash (order_hash);
create index tradelogs_oneinch_v6_order_hash_idx
    on tradelogs_oneinch_v6 USING hash (order_hash);
create index tradelogs_uniswapx_order_hash_idx
    on tradelogs_uniswapx USING hash (order_hash);
create index tradelogs_bebop_order_hash_idx
    on tradelogs_bebop USING hash (order_hash);
create index tradelogs_zerox_v3_order_hash_idx
    on tradelogs_zerox_v3 USING hash (order_hash);
create index tradelogs_pancakeswap_order_hash_idx
    on tradelogs_pancakeswap USING hash (order_hash);

create index tradelogs_cow_protocol_tx_hash_idx
    on tradelogs_cow_protocol USING hash (tx_hash);
create index cow_trade_callframe_tx_hash_idx
    on cow_trade_callframe USING hash (tx_hash);