-- name: CreateRequest :one
insert into requests (applicant_id, resource_id, access_type, access_reason, status)
values ($1, $2, $3, $4, 'pending') returning *;

-- -- name: ListRequests :many
-- select r.id, u.first_name || ' ' || u.last_name as applicant_name, d.name as department_name, res.name as resource_name, r.access_type, r.access_reason, r.status, r.created_at
-- from requests r
-- join users u on r.applicant_id = u.id
-- join departments d on r.department_id = d.id
-- join resources res on r.resource_id = res.id;

-- name: ListUserRequests :many
select r.id, res.id as resource_id, res.name as resource_name, r.access_type, r.access_reason, r.status, r.comments, r.created_at
from requests r
join resources res on r.resource_id = res.id
where r.applicant_id = $1;

-- name: GetRequestByID :one
select r.id, r.applicant_id, res.id as resource_id, res.name as resource_name, r.access_type, r.access_reason, r.status, r.comments, r.created_at
from requests r
join resources res on r.resource_id = res.id
where r.id = $1;

-- name: UpdateRequestStatus :exec
update requests
set status = $1
where id = $2;

-- name: DeleteRequest :exec
delete from requests
where id = $1;

-- name: CreateApprovalStep :exec
insert into approval_workflow (request_id, approver_id, step_order, status)
values ($1, $2, $3, $4);


