DROP TABLE IF EXISTS users;

CREATE TABLE IF NOT EXISTS users
(
    id           BIGSERIAL PRIMARY KEY,
    email        TEXT NOT NULL,
    password_enc TEXT NOT NULL,
    user_type    TEXT NOT NULL
);