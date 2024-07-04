BEGIN;

create table delegation(
    id integer primary key not null,
	amount      integer not null,
	sender_addr  varchar(255) not null,
	block_height integer not null,
	ts   timestamp not null,
	year        integer not null
);

COMMIT;