-- migrate:up
alter table beatmap_set
    drop column has_favourited;

-- migrate:down
alter table beatmap_set
    add has_favourited bool default false not null;
