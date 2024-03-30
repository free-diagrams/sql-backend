CREATE TABLE users
(
    id   SERIAL PRIMARY KEY,
    uid  INTEGER UNIQUE NOT NULL,
    name TEXT           NOT NULL
);

CREATE TABLE diagrams
(
    id         SERIAL PRIMARY KEY,
    name       TEXT      NOT NULL,
    owner_id   INTEGER   NOT NULL REFERENCES users (id),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    comment    TEXT
);

CREATE TABLE colors
(
    id       SERIAL PRIMARY KEY,
    name     TEXT UNIQUE NOT NULL,
    hex_code TEXT        NOT NULL
);

CREATE TABLE tables
(
    id            SERIAL PRIMARY KEY,
    logical_name  TEXT    NOT NULL,
    physical_name TEXT    NOT NULL,
    diagram_id    INTEGER NOT NULL REFERENCES diagrams (id),
    position_x    FLOAT   NOT NULL,
    position_y    FLOAT   NOT NULL,
    color_id      INTEGER NOT NULL REFERENCES colors (id),
    comment       TEXT    NOT NULL
);

CREATE TABLE databases
(
    id   SERIAL PRIMARY KEY,
    name TEXT UNIQUE NOT NULL
);

CREATE TABLE data_types
(
    id   SERIAL PRIMARY KEY,
    name TEXT UNIQUE NOT NULL
);

CREATE TABLE databases_data_types
(
    id           SERIAL PRIMARY KEY,
    database_id  INTEGER NOT NULL REFERENCES databases (id),
    data_type_id INTEGER NOT NULL REFERENCES data_types (id)
);

CREATE TABLE rows
(
    id                    SERIAL PRIMARY KEY,
    logical_name          TEXT    NOT NULL,
    physical_name         TEXT    NOT NULL,
    table_id              INTEGER NOT NULL REFERENCES tables (id),
    is_primary_key        BOOLEAN,
    is_unique             BOOLEAN,
    is_nullable           BOOLEAN,
    default_value         TEXT,
    database_data_type_id INTEGER NOT NULL REFERENCES databases_data_types (id),
    comment               TEXT
);

CREATE TYPE RELATION_TYPE AS ENUM ('one', 'many');

CREATE TABLE relations
(
    id                       SERIAL PRIMARY KEY,
    first_row_id             INTEGER       NOT NULL REFERENCES rows (id),
    first_row_relation_type  RELATION_TYPE NOT NULL,
    second_row_id            INTEGER       NOT NULL REFERENCES rows (id),
    second_row_relation_type RELATION_TYPE NOT NULL
);
