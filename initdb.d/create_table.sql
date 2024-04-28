create table users (
    id int primary key,
    user_name varchar(255),
    email varchar(255) unique not null,
    password varchar(255) not null
);