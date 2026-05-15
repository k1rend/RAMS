package approve

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/k1rend/RAMS/repo"
)

type ApprovalService struct {
	repo *repo.Queries
}

func NewApprovalService(repo *repo.Queries) *ApprovalService {
	return &ApprovalService{repo: repo}
}

func (s *ApprovalService) ListPendingApprovals(ctx context.Context, approverID int32) ([]repo.ListPendingApprovalsRow, error) {
	return s.repo.ListPendingApprovals(ctx, pgtype.Int4{Int32: approverID, Valid: true})
}

// ApproveRequest handles the approval of a request by an approver.
func (s *ApprovalService) ApproveStep(ctx context.Context, approveStepID int32, approverID int32) error {
	return nil
}
