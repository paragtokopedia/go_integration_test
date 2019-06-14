create schema if not exists test;

create table if not exists test.user (
  id int primary key AUTO_INCREMENT,
  name varchar(10)
)