-- migrate:up
create table beatmaps
(
    id serial not null,
    beatmapset_id int not null
        constraint beatmaps_beatmap_set_id_fk
            references beatmap_set,
    mode varchar default 'osu' not null,
    mode_int int default 0 not null,
    convert bool default false not null,
    difficulty_rating float default 1.0 not null,
    version varchar default '' not null,
    total_length int default 100 not null,
    hit_length int default 100,
    bpm int default 100 not null,
    cs int default 5 not null,
    drain int default 5 not null,
    accuracy int default 5 not null,
    ar int default 5 not null,
    playcount int default 0 not null,
    passcount int default 0 not null,
    count_circles int default 0 not null,
    count_sliders int default 0 not null,
    count_spinners int default 0 not null,
    count_total int default 0 not null,
    is_scoreable bool default true not null,
    last_updated timestamptz default current_timestamp not null,
    ranked int default 0 not null,
    status varchar default 'ranked' not null,
    url varchar default '' not null,
    deleted_at timestamptz default null,
    max_combo int default null
);

create unique index beatmaps_id_uindex
    on beatmaps (id);

alter table beatmaps
    add constraint beatmaps_pk
        primary key (id);

-- migrate:down
drop table beatmaps;
