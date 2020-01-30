-- migrate:up
alter table message drop constraint message_channels_id_fk;

alter table message
    add constraint message_channels_id_fk
        foreign key (channel_id) references channels
            on delete cascade;

-- migrate:down

