DROP TABLE IF EXISTS subscriptions;

CREATE TABLE IF NOT EXISTS subscriptions
(
    email            TEXT NOT NULL,
    house_id         BIGINT NOT NULL REFERENCES houses(id) ON DELETE CASCADE,
    created_at       TIMESTAMP NOT NULL
);
