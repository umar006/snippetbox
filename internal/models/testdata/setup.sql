CREATE TABLE snippets (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(100) NOT NULL,
    content TEXT NOT NULL,
    created DATETIME NOT NULL,
    expires DATETIME NOT NULL
);

CREATE INDEX idx_snippets_created ON snippets(created);

CREATE TABLE users (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    hashed_password CHAR(60) NOT NULL,
    created DATETIME NOT NULL
);

ALTER TABLE users ADD CONSTRAINT users_uc_email UNIQUE (email);

INSERT INTO users (name, email, hashed_password, created) VALUES (
    'Alice Jones',
    'alice@example.com',
    '$2a$12$ITTeFp.T37uZnh4Fck2FCe5dHtQDlPGYlobK.TwhIcQ/84LUROhv.',
    '2022-01-01 10:00:00'
);

INSERT INTO snippets (title, content, created, expires)
VALUES 
    ('cat1', 'pompomtut1', '2022-02-02 10:00:00', '2023-02-02 10:00:00'),
    ('cat2', 'pompomtut2', '2022-02-02 10:00:00', '2023-02-02 10:00:00'),
    ('cat3', 'pompomtut3', '2022-02-02 10:00:00', '2023-02-02 10:00:00');
