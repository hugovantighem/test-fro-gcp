BEGIN;

create table delegations(
    id bigint primary key not null,
	amount      bigint not null,
	sender_addr  varchar(255) not null,
	block_height integer not null,
	ts   timestamp not null,
	year        integer not null
);

COMMIT;