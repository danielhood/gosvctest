drop database if exists test;

create database test;

use test;

create table test1 (
  id int not null primary key,
  test_value nvarchar(200) not null
);

insert into test1 (id, test_value) values (1, 'test value 1');

insert into test1 (id, test_value) values (2, 'test value 2');

insert into test1 (id, test_value) values (3, 'test value 3');
