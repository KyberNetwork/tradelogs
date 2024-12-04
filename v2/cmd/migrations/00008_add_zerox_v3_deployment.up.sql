create table zerox_v3_deployment
(
    block_number     bigint not null,
    contract_type    int    not null,
    contract_address text   not null,
    primary key (contract_type, contract_address)
);

insert into zerox_v3_deployment (block_number, contract_type, contract_address) values
    (19921333, 2, '0xecf4248a682ffc676f4c596214cd6a4b463d8d2e'),
    (19921333, 3, '0x6d837c5f39d609b996ad4ad45f744ad2df93fb57'),
    (20027105, 2, '0xbcd5e096c749bf7dd8cdb5ae90a2186866b67d77'),
    (20027105, 3, '0x4ee2d0b91084821b96365f93da61cf86e10c42bb'),
    (20064192, 2, '0x7f6cee965959295cc64d0e6c00d99d6532d8e86b'),
    (20064192, 3, '0x7c39a136ea20b3483e402ea031c1f3c019bab24b'),
    (20412702, 2, '0x07e594aa718bb872b526e93eed830a8d2a6a1071'),
    (20412702, 3, '0x25b81ce58ab0c4877d25a96ad644491ceab81048'),
    (20514759, 2, '0x2c4b05349418ef279184f07590e61af27cf3a86b'),
    (20514759, 3, '0xae11b95c8ebb5247548c279a00120b0acadc7451'),
    (20616430, 2, '0x70bf6634ee8cb27d04478f184b9b8bb13e5f4710'),
    (20616430, 3, '0x12d737470fb3ec6c3deec9b518100bec9d520144');