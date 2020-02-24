-- migrate:up
create table user_month_playcount
(
    id serial not null,
    user_id int not null
        constraint user_month_playcount_users_id_fk
            references users
            on delete cascade,
    playcount int not null,
    year_month varchar not null
);

comment on column user_month_playcount.year_month is '{year}-{month}-01';

create unique index user_month_playcount_id_uindex
    on user_month_playcount (id);

alter table user_month_playcount
    add constraint user_month_playcount_pk
        primary key (id);

-- migrate:down
drop index user_month_playcount_id_uindex;
drop table user_month_playcount;
