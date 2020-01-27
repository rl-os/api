-- migrate:up
create table user_channels
(
    id serial not null,
    user_id int not null
        constraint user_channels_users_id_fk
            references users
            on delete cascade,
    channel_id int not null
        constraint user_channels_channels_id_fk
            references channels
            on delete cascade
);

create unique index user_channels_id_uindex
    on user_channels (id);

create index user_channels_user_id_index
    on user_channels (user_id);

alter table user_channels
    add constraint user_channels_pk
        primary key (id);


-- migrate:down
drop index user_channels_user_id_index;
drop index user_channels_id_uindex;
drop table user_channels;
