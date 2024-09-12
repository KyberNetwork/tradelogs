UPDATE tradelogs
SET state = 'new'
WHERE
    (
        (maker_token = '0x0000000000000000000000000000000000000000' AND maker_token_price = 0) OR 
        (taker_token = '0x0000000000000000000000000000000000000000' AND taker_token_price = 0)
    )
    AND state = 'processed';
