drop database blockchain_app

create database blockchain_app

create table user (
    id int not null serial primary key
    username varchar(50) not null
    passwordHash varchar(50) not null
)