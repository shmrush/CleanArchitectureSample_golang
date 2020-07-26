-- +migrate Up
create table users(
  id bigint(20) not null primary key auto_increment,
  name varchar(255) not null,
  email varchar(255) not null unique,
  encrypted_password varchar(255) not null,
  created_at datetime(6) not null,
  updated_at datetime(6) not null,
  deleted_at datetime(6),
  index idx_users_name(name)
);

-- +migrate Down
drop table users;
