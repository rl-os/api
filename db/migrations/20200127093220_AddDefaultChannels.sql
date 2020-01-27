-- migrate:up
INSERT INTO channels (name, description, type, icon) VALUES ('#osu', 'classic osu channel', 'PUBLIC', DEFAULT);
INSERT INTO channels (name, description, type, icon) VALUES ('#announce', 'announces', 'PUBLIC', DEFAULT);
INSERT INTO channels (name, description, type, icon) VALUES ('#developers', 'developers world', 'PUBLIC', DEFAULT);
INSERT INTO channels (name, description, type, icon) VALUES ('#english', 'English communication channel', 'PUBLIC', DEFAULT);
INSERT INTO channels (name, description, type, icon) VALUES ('#russia', 'Канал для общения на русском', 'PUBLIC', DEFAULT);
INSERT INTO channels (name, description, type, icon) VALUES ('#german', 'Deutscher Kanal', 'PUBLIC', DEFAULT);

-- migrate:down

DELETE  FROM channels WHERE type = 'PUBLIC';