UPDATE tradelogs
SET state = 'new'
WHERE
    (
        maker_token = '0xac57de9c1a09fec648e93eb98875b212db0d460b'
        OR taker_token = '0xac57de9c1a09fec648e93eb98875b212db0d460b'
    )
    AND state = 'processed';
