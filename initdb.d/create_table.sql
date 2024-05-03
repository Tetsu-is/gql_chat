create table users (
    id int primary key,
    user_name varchar(255),
    email varchar(255) unique not null,
    password varchar(255) not null
);

create table messages {
    id int primary key,
    user_id int,
    user_name varchar(255) not null,
    body varchar(255),
    created_at timestamp,
    foreign key (user_id) references users(id)
};