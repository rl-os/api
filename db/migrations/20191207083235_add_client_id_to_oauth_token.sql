-- migrate:up
alter table oauth_token
    add client_id int not null default 0;

-- migrate:down

alter table oauth_token
    drop column client_id;
