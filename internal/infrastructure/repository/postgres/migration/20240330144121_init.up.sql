CREATE TABLE users
(
    id         TEXT PRIMARY KEY,
    username   TEXT NOT NULL,
    last_name  TEXT NOT NULL,
    first_name TEXT NOT NULL
);

CREATE TABLE databases
(
    id   TEXT PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE diagrams
(
    id          TEXT PRIMARY KEY,
    database_id TEXT    NOT NULL REFERENCES databases (id),
    name        TEXT    NOT NULL,
    is_public   BOOLEAN NOT NULL,
    comment     TEXT
);

CREATE TABLE access_rights
(
    id          TEXT PRIMARY KEY,
    user_id     TEXT NOT NULL REFERENCES users (id),
    diagram_id  TEXT NOT NULL REFERENCES diagrams (id),
    access_type TEXT NOT NULL
);

CREATE TABLE data_types
(
    id          TEXT PRIMARY KEY,
    database_id TEXT NOT NULL REFERENCES databases (id),
    name        TEXT NOT NULL
);

CREATE TABLE tables
(
    id            TEXT PRIMARY KEY,
    diagram_id    TEXT  NOT NULL REFERENCES diagrams (id),
    color_rgb     TEXT  NOT NULL,
    logical_name  TEXT  NOT NULL,
    physical_name TEXT  NOT NULL,
    position_x    FLOAT NOT NULL,
    position_y    FLOAT NOT NULL,
    comment       TEXT  NOT NULL
);

CREATE TABLE rows
(
    id             TEXT PRIMARY KEY,
    table_id       TEXT    NOT NULL REFERENCES tables (id),
    data_type_id   TEXT    NOT NULL REFERENCES data_types (id),
    color_rgb      TEXT    NOT NULL,
    logical_name   TEXT    NOT NULL,
    physical_name  TEXT    NOT NULL,
    is_primary_key BOOLEAN NOT NULL,
    is_unique      BOOLEAN NOT NULL,
    is_nullable    BOOLEAN NOT NULL,
    default_value  TEXT,
    comment        TEXT
);

CREATE TABLE relations
(
    first_row_id             TEXT NOT NULL REFERENCES rows (id),
    first_row_relation_type  TEXT NOT NULL,
    second_row_id            TEXT NOT NULL REFERENCES rows (id),
    second_row_relation_type TEXT NOT NULL
);
