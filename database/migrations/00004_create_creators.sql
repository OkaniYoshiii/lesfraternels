-- +goose up
-- +goose statementbegin
CREATE TABLE creators (
    id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE
);
-- +goose statementend

-- +goose down
-- +goose statementbegin
DROP TABLE creators;
-- +goose statementend