update tradelogs_bebop set expiration = 0 where expiration is null;
alter table tradelogs_bebop alter column expiration set not null;
alter table tradelogs_bebop alter column expiration set default 0;

update tradelogs_hashflow_v3 set expiration = 0 where expiration is null;
alter table tradelogs_hashflow_v3 alter column expiration set not null;
alter table tradelogs_hashflow_v3 alter column expiration set default 0;

update tradelogs_kyberswap_rfq set expiration = 0 where expiration is null;
alter table tradelogs_kyberswap_rfq alter column expiration set not null;
alter table tradelogs_kyberswap_rfq alter column expiration set default 0;

update tradelogs_oneinch_v6 set expiration = 0 where expiration is null;
alter table tradelogs_oneinch_v6 alter column expiration set not null;
alter table tradelogs_oneinch_v6 alter column expiration set default 0;

update tradelogs_paraswap set expiration = 0 where expiration is null;
alter table tradelogs_paraswap alter column expiration set not null;
alter table tradelogs_paraswap alter column expiration set default 0;

update tradelogs_zerox set expiration = 0 where expiration is null;
alter table tradelogs_zerox alter column expiration set not null;
alter table tradelogs_zerox alter column expiration set default 0;

update tradelogs_zerox_v3 set expiration = 0 where expiration is null;
alter table tradelogs_zerox_v3 alter column expiration set not null;
alter table tradelogs_zerox_v3 alter column expiration set default 0;
