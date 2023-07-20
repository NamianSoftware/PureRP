CREATE TABLE users
(
    id            serial primary key,
    first_name    varchar(255) not null,
    last_name     varchar(255) not null,
    email         varchar(320) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE characters
(
    id         serial primary key,
    first_name varchar(255) not null,
    last_name  varchar(255) not null
);

CREATE TABLE users_characters
(
    id           serial primary key,
    user_id      int references users (id) on delete cascade unique,
    character_id int references characters (id) on delete cascade
);