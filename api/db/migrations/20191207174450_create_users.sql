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
    avatar_url      varchar   default 'https://301222.selcdn.ru/akasi/avatars/1.png' not null,
    country_code    varchar   default '-'               not null,
    default_group   varchar   default 'osu'             not null,
    can_moderate    bool      default false             not null,
    interests       varchar   default                   null,
    occupation      varchar   default ''                not null,
    title           varchar   default                   null,
    location        varchar   default                   null,
    twitter         varchar   default                   null,
    lastfm          varchar   default                   null,
    skype           varchar   default                   null,
    website         varchar   default                   null,
    discord         varchar   default                   null,
    playstyle       varchar   array default '{}'        not null,
    playmode        varchar   default ''                not null,
    cover_url       varchar   default 'https://301222.selcdn.ru/akasi/bg/1.jpg' not null,
    max_blocks      int       default 50                not null,
    max_friends     int       default 100               not null
);
alter table users
    owner to postgres;

create unique index users_id_uindex
    on users (id);

create unique index users_username_uindex
    on users (username);

create unique index users_email_uindex
    on users (email);

-- migrate:down

drop index users_id_uindex;
drop index users_username_uindex;
drop table users;