-- migrate:up
alter table channels
    add created_at timestamp default current_timestamp not null;

-- migrate:down
alter table channels
    drop column created_at;
