UPDATE tradelogs SET 
maker_token = taker_token,
taker_token = maker_token,
maker_token_amount = taker_token_amount,
taker_token_amount = maker_token_amount
WHERE taker = '0x807cf9a772d5a3f9cefbc1192e939d62f0d9bd38';
