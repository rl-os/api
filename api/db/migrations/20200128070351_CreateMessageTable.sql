-- migrate:up
create table message
(
    id serial not null,
    sender_id int not null
        constraint message_users_id_fk
            references users,
    channel_id int not null
        constraint message_channels_id_fk
            references channels,
    created_at timestamp default CURRENT_TIMESTAMP not null,
    content varchar not null,
    is_action bool default false not null
);

create index message_channel_id_index
    on message (channel_id);

create unique index message_id_uindex
    on message (id);

create index message_sender_id_index
    on message (sender_id);

alter table message
    add constraint message_pk
        primary key (id);

-- migrate:down
drop index message_channel_id_index;
drop index message_id_uindex;
drop index message_sender_id_index;
drop table message;

