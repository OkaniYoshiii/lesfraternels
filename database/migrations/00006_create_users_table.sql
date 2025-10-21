-- +goose up
-- +goose statementbegin
CREATE TABLE users (
    id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL
);
-- +goose statementend

-- +goose down
-- +goose statementbegin
DROP TABLE users;
-- +goose statementend
