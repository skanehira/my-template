create database test;
go

use test;
go

create table users(
    id int identity(1,1) primary key,
    name nvarchar(32),
    birthday datetime
);
go

insert into users values ("gorilla", GETDATE());
go
