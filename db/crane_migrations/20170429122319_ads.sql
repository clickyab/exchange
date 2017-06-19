
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
create table ads
(
  id int not null auto_increment
    primary key,
  target tinyint not null,
  height int not null,
  width int not null,
  active INT default 0 not null,
  user_id int not null,
  url text not null,
  src varchar(255) null,
  attribute text,
  created_at timestamp default CURRENT_TIMESTAMP not null,
  updated_at timestamp default CURRENT_TIMESTAMP not null,
  constraint ads_users_id_fk
    foreign key (user_id) references users (id)
)
;

create index ads_users_id_fk
  on ads (user_id)
;

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE ads;
