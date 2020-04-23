-- migrate:up
create table beatmap_set
(
    id serial not null,
    last_checked timestamptz default current_timestamp not null,
    title varchar default '' not null,
    artist varchar default '' not null,
    play_count int default 0 not null,
    favourite_count int default 0 not null,
    has_favourited bool default false,
    submitted_date timestamptz default current_timestamp not null,
    last_updated timestamptz default current_timestamp not null,
    ranked_date timestamptz default null,
    creator varchar default '' not null,
    user_id int not null,
    bpm int default 150 not null,
    source varchar not null,
    covers json default json_build_object() not null,
    preview_url varchar not null,
    tags varchar not null,
    video bool default false not null,
    storyboard bool default false not null,
    ranked int default 0 not null,
    status varchar not null,
    is_scoreable bool default true not null,
    discussion_enabled bool default true not null,
    discussion_locked bool default false not null,
    can_be_hyped bool default true not null,
    availability json default json_build_object() not null,
    hype json default json_build_object() not null,
    nominations json default json_build_object() not null,
    legacy_thread_url varchar default '' not null,
    description json default json_build_object('description', '') not null,
    genre json default json_build_object('id', 1, 'name', 'None'),
    language json default json_build_object('id', 1, 'name', 'None') not null,
    "user" json default json_build_object() not null
);

comment on column beatmap_set.has_favourited is 'TODO THIS';

comment on column beatmap_set.user_id is 'user in original bancho';

create unique index beatmap_set_id_uindex
    on beatmap_set (id);

alter table beatmap_set
    add constraint beatmap_set_pk
        primary key (id);

-- migrate:down
drop table beatmap_set;
