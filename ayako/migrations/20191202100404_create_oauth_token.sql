-- migrate:up
create table oauth_token
(
    id            serial                              not null
        constraint oauth_token_pk
            primary key,
    user_id       integer                             not null,
    access_token  varchar                             not null,
    refresh_token varchar                             not null,
    scopes        varchar                             not null,
    revoked       boolean   default false             not null,
    expires_at    timestamp                           not null,
    created_at    timestamp default CURRENT_TIMESTAMP not null
);
alter table oauth_token
    owner to postgres;

create unique index oauth_token_id_uindex
    on oauth_token (id);

-- migrate:down

drop index oauth_token_id_uindex;
drop table oauth_token;