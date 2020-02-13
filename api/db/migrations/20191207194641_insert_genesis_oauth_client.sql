-- migrate:up
INSERT INTO oauth_client (id, user_id, name, secret, redirect) VALUES (5, 1, 'genesis', 'FGc9GAtyHzeQDshWP5Ah7dega8hJACAJpQtw6OXk', 'https://example.com');

-- migrate:down

DELETE  FROM oauth_client WHERE id = 5;