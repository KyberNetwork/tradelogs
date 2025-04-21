alter table cow_transfer_event 
    add COLUMN transfer_id serial,
    add COLUMN transfer_type text default 'ERC20';
alter table cow_transfer_event drop constraint cow_transfer_event_pkey;
alter table cow_transfer_event add primary key (transfer_id);

create table cow_trade_callframe
(
    tx_hash    text not null,
    call_frame text not null,
    primary key (tx_hash)
);
