-- name: ListMods :many
SELECT * FROM mods;

-- name: CreateMod :execresult
INSERT INTO mods (
    `name`,
    `description`,
    `image`
) VALUES (
    ?, ?, ? 
);

-- name: DeleteMod :exec
DELETE FROM mods WHERE id = ?;

-- name: CreateMember :execresult
INSERT INTO members (
    `name`,
    `image`
) VALUES (
    ?, ?
);

-- name: ListMembers :many
SELECT * FROM members;

-- name: DeleteMember :exec
DELETE FROM members WHERE id = ?;