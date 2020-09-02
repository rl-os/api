-- migrate:up
alter table users
    add support_expired_at timestamp default CURRENT_TIMESTAMP not null;

-- migrate:down

alter table users
    drop column support_expired_at;