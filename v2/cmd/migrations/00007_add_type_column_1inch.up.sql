alter table tradelogs_oneinch_v6 add column type text default ''::text not null;

alter table tradelogs_oneinch_v6
    alter column maker_traits drop default,
    alter column maker_traits type json using maker_traits::json,
    alter column maker_traits set default '{}'::json;