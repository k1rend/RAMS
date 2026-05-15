-- name: ListPendingApprovals :many
select aw.id, r.id as request_id, u.first_name || ' ' || u.last_name as applicant_name, res.name as resource_name, r.access_type, r.access_reason, aw.status, aw.created_at
from approval_workflow aw
join requests r on aw.request_id = r.id
join users u on r.applicant_id = u.id
join resources res on r.resource_id = res.id
where aw.approver_id = $1 and aw.status = 'pending'
order by aw.created_at;

-- name: UpdateApprovalStatus :exec
update approval_workflow
set status = $1
where id = $2;