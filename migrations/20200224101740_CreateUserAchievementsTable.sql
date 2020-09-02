-- migrate:up
create table user_achievements
(
    id serial not null,
    achievement_id int not null
        constraint user_achievements_achievements_fk
            references achievements (id)
            on delete cascade,
    user_id int not null
        constraint user_achievements_users_id_fk
            references users
            on delete cascade,
    created_at timestamp default CURRENT_TIMESTAMP
);

create unique index user_achievements_id_uindex
    on user_achievements (id);

create index user_achievements_user_id_index
    on user_achievements (user_id);

alter table user_achievements
    add constraint user_achievements_pk
        primary key (id);

-- migrate:down
drop index user_achievements_id_uindex;
drop index user_achievements_user_id_index;
drop table user_achievements;
