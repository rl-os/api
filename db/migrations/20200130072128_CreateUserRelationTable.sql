-- migrate:up
create table user_relation
(
    id serial not null,
    user_id int not null,
    target_id int not null,
    friend bool default false not null,
    constraint user_relation_users_id_id_fk
        foreign key (user_id) references users (id)
            on delete cascade,
    constraint user_relation_target_id_id_fk
        foreign key (target_id) references users (id)
            on delete cascade
);

create unique index user_relation_id_uindex
    on user_relation (id);

create index user_relation_user_id_target_id_index
    on user_relation (user_id, target_id);

alter table user_relation
    add constraint user_relation_pk
        primary key (id);

-- migrate:down
drop index user_relation_id_uindex;
drop index user_relation_user_id_target_id_index;
drop table user_relation;
