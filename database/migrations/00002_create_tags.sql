-- +goose up
-- +goose statementbegin
CREATE TABLE tags (
    id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE
);
-- +goose statementend

-- +goose down
-- +goose statementbegin
DROP TABLE tags;
-- +goose statementend