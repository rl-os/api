-- migrate:up
create function check_online(val timestamp without time zone) returns bool immutable as $$
BEGIN
    RETURN (val > (CURRENT_TIMESTAMP - (10 ||' minutes')::interval));
END;
$$
LANGUAGE PLPGSQL;

alter table users
    add is_online bool GENERATED ALWAYS AS(check_online(last_visit)) STORED;

-- migrate:down
drop function check_online;

alter table users
    drop column is_online;
