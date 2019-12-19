-- migrate:up
create table user_details
(
    id serial not null,
    user_id int not null
        constraint user_details_users_id_fk
            references users
            on delete cascade,
    can_moderate bool default false not null,
    interests varchar default null,
    occupation varchar default '' not null,
    title varchar default null,
    location varchar default null,
    twitter varchar default null,
    lastfm varchar default null,
    skype varchar default null,
    website varchar default null,
    discord varchar default null,
    playstyle varchar array default '{}',
    playmode varchar default '' not null,
    cover_url varchar default '',
    max_blocks int default 50 not null,
    max_friends int default 100
);

create unique index user_details_id_uindex
    on user_details (id);

create unique index user_details_user_id_uindex
    on user_details (user_id);

alter table user_details
    add constraint user_details_pk
        primary key (id);

-- insert genesis user details
INSERT INTO user_details ("id", "user_id", "can_moderate", "interests", "occupation", "title", "location", "twitter", "lastfm", "skype", "website", "discord", "playstyle", "playmode", "cover_url", "max_blocks", "max_friends") VALUES (DEFAULT, 1, DEFAULT, null, DEFAULT, null, null, null, null, null, null, null, DEFAULT, DEFAULT, DEFAULT, DEFAULT, DEFAULT);


-- migrate:down
drop index user_details_id_uindex;
drop index user_details_user_id_uindex;
drop table user_details;