-- name: CreateResource :exec
insert into resources (name, description, resource, resource_type, owner_id)
values ($1, $2, $3, $4, $5);

-- name: ListResources :many
select r.id, r.name, r.description, r.resource_type
from resources r;

-- name: GetResourceOwnerID :one
select owner_id from resources where id = $1;