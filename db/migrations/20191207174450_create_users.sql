-- migrate:up
create table users
(
    id              serial                              not null constraint users_pk primary key,
    username        varchar                             not null,
    email           varchar                             not null,
    password_hash   varchar                             not null,
    last_visit      timestamp default CURRENT_TIMESTAMP not null,
    created_at      timestamp default CURRENT_TIMESTAMP not null,
    is_bot          bool      default false             not null,
    is_active       bool      default true              not null,
    is_supporter    bool      default false             not null,
    has_supported   bool      default false             not null,
    support_level   int       default 0                 not null,
    pm_friends_only bool      default false             not null,
    avatar_url      varchar   default ''                not null,
    country_code    varchar   default ''                not null,
    default_group   varchar   default 'osu'             not null
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