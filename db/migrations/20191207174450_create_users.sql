-- migrate:up
create table user
(
    id            serial                              not null constraint user_pk primary key,
    username      varchar                             not null,
    email         varchar                             not null,
    password_hash varchar                             not null,
    last_visit    timestamp default CURRENT_TIMESTAMP not null,
    created_at    timestamp default CURRENT_TIMESTAMP not null,
);
alter table oauth_token
    owner to postgres;

create unique index user_username_uindex
    on user (username);

-- migrate:down

drop index user_id_uindex;
drop index user_username_uindex;
drop table user;