-- migrate:up
alter table users
    add deleted_at timestamp default null;


-- migrate:down

