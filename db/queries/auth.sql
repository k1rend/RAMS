-- name: CreateUser :one
insert into users (username, email, password_hash, first_name, last_name, department_id)
values ($1, $2, $3, $4, $5, $6) returning *;

-- name: GiveUserRole :exec
insert into user_roles (user_id, role_id) values ($1, $2);

-- name: GetUserByUsername :one
select * from users where username = $1;

-- name: GetUserRoles :many
select r.code from roles r
join user_roles ur on r.id = ur.role_id
where ur.user_id = $1;