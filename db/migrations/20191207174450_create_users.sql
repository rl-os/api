-- migrate:up
create table users
(
    id            serial                              not null constraint users_pk primary key,
    username      varchar                             not null,
    email         varchar                             not null,
    password_hash varchar                             not null,
    last_visit    timestamp default CURRENT_TIMESTAMP not null,
    created_at    timestamp default CURRENT_TIMESTAMP not null
);
alter table users
    owner to postgres;

create unique index users_id_uindex
    on users (id);

create unique index users_username_uindex
    on users (username);

-- migrate:down

drop index users_id_uindex;
drop index users_username_uindex;
drop table users;