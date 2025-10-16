-- +goose up
-- +goose statementbegin
CREATE TABLE mods_tags (
    id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    mod_id INT UNSIGNED NOT NULL,
    tag_id INT UNSIGNED NOT NULL,

    CONSTRAINT FOREIGN KEY (mod_id) REFERENCES mods (id) ON DELETE CASCADE,
    CONSTRAINT FOREIGN KEY (tag_id) REFERENCES tags (id) ON DELETE CASCADE
);
-- +goose statementend

-- +goose down
-- +goose statementbegin
DROP TABLE mods_tags;
-- +goose statementend