-- migrate:up
alter table channels alter column icon drop not null;

-- migrate:down
alter table channels alter column icon set not null;
