-- migrate:up
INSERT INTO users (id, username, email, password_hash) VALUES (1, 'genesis', 'none@google.com', 'none');
INSERT INTO users (id, username, email, password_hash, avatar_url, is_bot) VALUES (2, 'chore', 'none@yandex.ru', 'none', 'https://301222.selcdn.ru/akasi/avatars/2.png', true);

-- migrate:down

DELETE  FROM users WHERE username = 'genesis' AND email = 'none@google.com' AND password_hash = 'none';
DELETE  FROM users WHERE username = 'chore' AND email = 'none@yandex.ru' AND password_hash = 'none' AND is_bot = true;