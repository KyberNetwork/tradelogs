alter table tradelogs_zerox
    alter column maker_token_price drop not null;

alter table tradelogs_zerox
    alter column maker_token_price drop default;

alter table tradelogs_zerox
    alter column taker_token_price drop not null;

alter table tradelogs_zerox
    alter column taker_token_price drop default;

alter table tradelogs_zerox
    alter column maker_usd_amount drop not null;

alter table tradelogs_zerox
    alter column maker_usd_amount drop default;

alter table tradelogs_zerox
    alter column taker_usd_amount drop not null;

alter table tradelogs_zerox
    alter column taker_usd_amount drop default;

alter table tradelogs_zerox drop column state;
