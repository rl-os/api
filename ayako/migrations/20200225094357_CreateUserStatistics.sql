-- migrate:up
create table user_statistics
(
    id serial not null,
    user_id int not null
        constraint user_statistics_users_id_fk
            references users
            on delete cascade,
    level_current int default 1 not null,
    level_progress int default 0 not null,
    pp float default 0.0 not null,
    ranked_score int default 0 not null,
    hit_accuracy float default 0.0 not null,
    play_count int default 0 not null,
    play_time int default 0 not null,
    total_score int default 0 not null,
    total_hits int default 0 not null,
    maximum_combo int default 0 not null,
    replays_watched_by_others int default 0 not null,
    is_ranked boolean default true not null,
    grade_counts_ss int default 0 not null,
    grade_counts_ssh int default 0 not null,
    grade_counts_s int default 0 not null,
    grade_counts_sh int default 0 not null,
    grade_counts_a int default 0 not null
);

create unique index user_statistics_id_uindex
    on user_statistics (id);

create index user_statistics_user_id_is_ranked_index
    on user_statistics (user_id, is_ranked);

alter table user_statistics
    add constraint user_statistics_pk
        primary key (id);

insert into user_statistics (user_id) values (1);
insert into user_statistics (user_id) values (2);

-- migrate:down
drop index user_statistics_id_uindex;
drop index user_statistics_user_id_is_ranked_index;
drop table user_statistics;
