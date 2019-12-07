-- migrate:up
INSERT INTO users (username, email, password_hash) VALUES ('genesis', 'none@google.com', 'none');

-- migrate:down

DELETE  FROM users WHERE username = 'genesis' AND email = 'none@google.com' AND password_hash = 'none';