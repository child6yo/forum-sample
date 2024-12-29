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
    user_id     int          references users (id) on delete cascade not null,
    title       varchar(255) not null,
    content     text         not null,
    cr_time     timestamp    not null,
    update      bool         not null,
    upd_time    timestamp    not null
);

CREATE TABLE threads
(
    id          serial       not null unique,
    user_id     int references users (id) on delete cascade not null,
    content     text         not null,
    answer_at   int          not null,
    cr_time     timestamp    not null,
    update      bool         not null,
    upd_time    timestamp    not null
);

CREATE TABLE post_threads
(
    id              serial  not null unique,
    post_id         int     references posts (id) on delete cascade not null,
    thread_id       int     references threads (id) on delete cascade not null
);

