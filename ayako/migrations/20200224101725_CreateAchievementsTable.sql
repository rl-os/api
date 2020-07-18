-- migrate:up
create table achievements
(
    id serial not null,
    name varchar not null,
    description varchar not null,
    enabled bool default true not null,
    grouping varchar not null,
    image varchar,
    mode varchar default 'osu',
    quest_instructions varchar,
    slug varchar not null
);

create unique index achievements_id_uindex
    on achievements (id);

create unique index achievements_slug_uindex
    on achievements (slug);

alter table achievements
    add constraint achievements_pk
        primary key (id);


-- migrate:down
drop index achievements_id_uindex;
drop index achievements_slug_uindex;
drop table achievements;
