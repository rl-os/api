-- migrate:up
create function check_online(val timestamp without time zone) returns bool immutable as $$
BEGIN
    RETURN (val > (CURRENT_TIMESTAMP - (10 ||' minutes')::interval));
END;
$$
LANGUAGE PLPGSQL;


-- migrate:down
drop function check_online;
