-- name: CreateUser :one
insert into users (username, email, password_hash, first_name, last_name, department_id, manager_id, job_title)
values ($1, $2, $3, $4, $5, $6, $7, $8) returning *;

-- name: GiveUserRole :exec
insert into user_roles (user_id, role_id) values ($1, $2);

-- name: GetUserByUsername :one
select * from users where username = $1;

-- name: GetUserRoles :many
select r.code from roles r
join user_roles ur on r.id = ur.role_id
where ur.user_id = $1;

-- name: GetUserInfoByID :one
select u.first_name || ' ' || u.last_name, u.email, u.job_title, d.name as department_name
from users u
left join departments d on u.department_id = d.id
where u.id = $1;

-- name: GetManagerID :one
select manager_id from users where id = $1;

-- name: GetSecurityEmployeesID :many
select user_id from user_roles where role_id = (select id from roles where code = 'security');