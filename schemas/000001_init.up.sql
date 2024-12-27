CREATE TABLE users
(
    id            serial       not null unique,
    email         varchar(255) not null unique,
    username      varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE posts
(
    id          serial       not null unique,
    user_id     int references users (id) on delete cascade not null,
    title       varchar(255) not null,
    content     text         not null,
    cr_date     timestamp    not null
);
