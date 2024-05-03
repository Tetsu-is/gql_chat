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

/* insert demo datas */
insert into messages(id, user_id, user_name, body, created_at) values(1, 1, 'user1', 'hello1', '2021-01-01 00:00:00'),
(2, 1, 'user1', 'world2', '2021-01-01 00:00:01'),
(3, 2, 'user2', 'hello3', '2021-01-01 00:00:02'),
(4, 2, 'user2', 'world4', '2021-01-01 00:00:03');