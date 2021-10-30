CREATE TABLE "episode"
(
    id           TEXT PRIMARY KEY,
    publication  TEXT,
    series       INTEGER,
    episode      INTEGER,
    release_date TIMESTAMP,
    metadata     JSON,
    contributors JSON
);

CREATE TABLE "dialog"
(
    id           TEXT PRIMARY KEY,
    episode_id   TEXT REFERENCES episode (id),
    pos          INTEGER NOT NULL,
    offset       INTEGER NULL,
    type         TEXT    NOT NULL,
    actor        TEXT,
    content      TEXT,
    metadata     JSON,
    notable      BOOLEAN
);
