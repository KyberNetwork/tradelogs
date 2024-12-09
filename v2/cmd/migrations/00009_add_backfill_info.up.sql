insert into tradelogs_state (key, value) values ('enable_system_backfill', false);

update backfill set deploy_block = 10247094, backfilled_block = 0 where exchange = 'zerox';

insert into backfill (exchange, deploy_block, backfilled_block) values
('zeroxV3', 19921330, 0),
('paraswap', 14853783, 0),
('kyberswap', 16366767, 0),
('kyberswapRFQ', 18182728, 0),
('1inchV6', 19212918, 0),
('hashflowV3', 18424981, 0),
('uniswapx', 19814533, 0),
('bebop', 19783283, 0),
('pancakeswap', 20782532, 0);