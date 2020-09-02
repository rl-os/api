-- migrate:up
create table oauth_client
(
    id            serial                              not null
        constraint oauth_client_pk
            primary key,
    user_id       integer                             not null,
    name  varchar                             not null,
    secret  varchar                             not null,
    redirect varchar                             not null,
    revoked       boolean default false                            not null,
    created_at    timestamp default CURRENT_TIMESTAMP not null
);
alter table oauth_client
    owner to postgres;

create unique index oauth_client_id_uindex
    on oauth_client (id);

-- migrate:down

drop index oauth_client_id_uindex;
drop table oauth_client;