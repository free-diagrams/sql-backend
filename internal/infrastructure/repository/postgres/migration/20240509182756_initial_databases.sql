-- +goose Up
-- +goose StatementBegin
INSERT INTO databases(id, name)
VALUES ('947b7426-c6a1-4cc0-a3e3-f9c2aae91a9a', 'PostgreSQL');

INSERT INTO databases(id, name)
VALUES ('3ee32547-bff2-47c6-bb00-e33a8dcae5ab', 'Oracle');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE
FROM databases
WHERE id = '947b7426-c6a1-4cc0-a3e3-f9c2aae91a9a';

DELETE
FROM databases
WHERE id = '3ee32547-bff2-47c6-bb00-e33a8dcae5ab';
-- +goose StatementEnd
