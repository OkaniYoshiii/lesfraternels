-- +goose up
-- +goose statementbegin
CREATE TABLE mods (
    id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    description TEXT(5000) NOT NULL,
    image VARCHAR(255)
);

CREATE TABLE members (
    id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    image VARCHAR(255) NOT NULL UNIQUE
);
-- +goose statementend

-- +goose down
-- +goose statementbegin
DROP TABLE mods;
DROP TABLE members;
-- +goose statementend