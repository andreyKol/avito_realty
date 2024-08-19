DROP TABLE IF EXISTS flats;

CREATE TABLE IF NOT EXISTS flats
(
    id      BIGSERIAL PRIMARY KEY,
    house_id BIGINT NOT NULL REFERENCES houses(id) ON DELETE CASCADE,
    price   INT NOT NULL,
    rooms   INT NOT NULL,
    status  TEXT
);
