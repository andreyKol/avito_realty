DROP TABLE IF EXISTS houses;

CREATE TABLE IF NOT EXISTS houses
(
    id               BIGSERIAL PRIMARY KEY,
    address          TEXT NOT NULL,
    year             INT2 NOT NULL,
    developer        TEXT,
    created_at       TIMESTAMP NOT NULL,
    last_flat_added_at TIMESTAMP
);
