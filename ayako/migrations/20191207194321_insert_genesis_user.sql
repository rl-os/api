-- migrate:up
INSERT INTO users (username, email, password_hash) VALUES ('genesis', 'none@google.com', 'none');
INSERT INTO users (username, email, password_hash, avatar_url, is_bot) VALUES ('chore', 'none@yandex.ru', 'none', 'https://301222.selcdn.ru/akasi/avatars/2.png', true);
ALTER SEQUENCE users_id_seq RESTART WITH 100;
-- migrate:down

DELETE  FROM users WHERE username = 'genesis' AND email = 'none@google.com' AND password_hash = 'none';
DELETE  FROM users WHERE username = 'chore' AND email = 'none@yandex.ru' AND password_hash = 'none' AND is_bot = true;
ALTER SEQUENCE users_id_seq RESTART WITH 1;