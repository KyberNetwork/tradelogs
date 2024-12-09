-- backfill 1inchV6 from current block
update backfill set backfilled_block = 0 where exchange = '1inchV6';