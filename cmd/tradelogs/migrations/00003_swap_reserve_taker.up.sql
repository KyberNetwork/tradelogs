UPDATE on_chain_trade SET 
input_asset_symbol = output_asset_symbol,
output_asset_symbol = input_asset_symbol,
input_amount = output_amount,
output_amount = input_amount
WHERE taker = '0x807cf9a772d5a3f9cefbc1192e939d62f0d9bd38';
