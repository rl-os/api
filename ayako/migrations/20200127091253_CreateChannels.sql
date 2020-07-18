-- migrate:up
create table channels
(
    id serial not null,
    name varchar not null,
    description varchar not null,
    type varchar not null,
    icon varchar not null
);

create unique index channels_id_index
    on channels (id);

create index table_name_name_index
    on channels (name);

alter table channels
    add constraint channels_pk
        primary key (id);

-- migrate:down

drop table channels;