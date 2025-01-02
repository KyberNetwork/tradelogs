ALTER TABLE txorigin ADD COLUMN type TEXT;
UPDATE txorigin SET type = 'cow';

INSERT INTO txorigin (address, name, type) VALUES 
('0x421925abb2e3311571f18e545e37fb983318f274', 'ktt-operator-00', 'vt'),
('0xd2e018c16dc6a4ce18e7880bed5b1b3b64fc35c7', 'ktt-operator-01', 'vt'),
('0x86ee24a1effa6bdf75a311c9a3b73ca0b0ec3207', 'ktt-operator-10', 'vt'),
('0x866bdd48566d984f94c55f5ff8988208b3aa0008', 'ktt-operator-11', 'vt'),
('0x869f1a409aa364ed2d7f1a0327fe8844d547f909', 'ktt-operator-12', 'vt'),
('0x8686ca74753f976ff7ead5a692096f2cd1388c00', 'ktt-operator-02', 'vt'),
('0x8686301a430112d76fef8331821aef42c8c48686', 'ktt-operator-03', 'vt,uniswapx,pcsx'),
('0x8686e88c147fadc829bfef8e2d853e3d91ff8101', 'ktt-operator-04', 'uniswapx,pcsx'),
('0x868601d3693003cfd2f6b8ddd6945569ce5d7602', 'ktt-operator-05', '1inch-resolver'),
('0x8686b3d43415e3d4bfbf686450a2799b9249fd03', 'ktt-operator-06', '1inch-resolver'),
('0x86b1980c889ad981837a5e8b587ae82100fe8e04', 'ktt-operator-07', 'vt'),
('0x860d6604b813376645b996278464e9953f884705', 'ktt-operator-08', 'vt'),
('0x86e81996242d1dce7fc5f9b9d060bbd3410fc206', 'ktt-operator-09', 'vt');

DROP INDEX zerox_timestamp_idx;
DROP INDEX kyberswap_timestamp_idx;
DROP INDEX kyberswap_rfq_timestamp_idx;
DROP INDEX paraswap_timestamp_idx;
DROP INDEX hashflow_v3_timestamp_idx;
DROP INDEX oneinch_v6_timestamp_idx;
DROP INDEX uniswapx_timestamp_idx;
DROP INDEX bebop_timestamp_idx;
DROP INDEX zerox_v3_timestamp_idx;
DROP INDEX pancakeswap_timestamp_idx;

create index zerox_timestamp_idx
    on tradelogs_zerox (timestamp DESC);
create index kyberswap_timestamp_idx
    on tradelogs_kyberswap (timestamp DESC);
create index kyberswap_rfq_timestamp_idx
    on tradelogs_kyberswap_rfq (timestamp DESC);
create index paraswap_timestamp_idx
    on tradelogs_paraswap (timestamp DESC);
create index hashflow_v3_timestamp_idx
    on tradelogs_hashflow_v3 (timestamp DESC);
create index oneinch_v6_timestamp_idx
    on tradelogs_oneinch_v6 (timestamp DESC);
create index uniswapx_timestamp_idx
    on tradelogs_uniswapx (timestamp DESC);
create index bebop_timestamp_idx
    on tradelogs_bebop (timestamp DESC);
create index zerox_v3_timestamp_idx
    on tradelogs_zerox_v3 (timestamp DESC);
create index pancakeswap_timestamp_idx
    on tradelogs_pancakeswap (timestamp DESC);
