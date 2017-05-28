
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied

create table user
(
	id int auto_increment primary key,
	username varchar(25) null,
	password varchar(25) not null,
	token varchar(50) null,
	email varchar(25) null,
	constraint user_id_uindex unique (id),
	constraint user_username_uindex unique (username),
	constraint user_token_uindex unique (token),
	constraint user_email_uindex unique (email)
);



-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE user;
