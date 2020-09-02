-- migrate:up
create table favouritemaps
(
    id serial not null,
    beatmapset_id int not null,
    user_id int not null,
    created_at timestamptz default current_timestamp not null
);

create unique index favouritemaps_id_uindex
    on favouritemaps (id);

create unique index favouritemaps_user_record_index
    on favouritemaps (user_id, beatmapset_id);

alter table favouritemaps
    add constraint favouritemaps_pk
        primary key (id);

-- migrate:down
drop table favouritemaps;
