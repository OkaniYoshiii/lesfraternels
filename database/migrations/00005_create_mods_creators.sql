-- +goose up
-- +goose statementbegin
CREATE TABLE mods_creators (
    id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    mod_id INT UNSIGNED NOT NULL,
    creator_id INT UNSIGNED NOT NULL,

    CONSTRAINT FOREIGN KEY (mod_id) REFERENCES mods (id) ON DELETE CASCADE,
    CONSTRAINT FOREIGN KEY (creator_id) REFERENCES creators (id) ON DELETE CASCADE
);
-- +goose statementend

-- +goose down
-- +goose statementbegin
DROP TABLE mods_creators;
-- +goose statementend