package request

import (
	"context"
	"math/rand"

	"github.com/k1rend/RAMS/repo"

	"github.com/jackc/pgx/v5/pgtype"
)

type RequestService struct {
	Repo *repo.Queries
}

func NewRequestService(repo *repo.Queries) *RequestService {
	return &RequestService{
		Repo: repo,
	}
}

func (s *RequestService) CreateRequest(
	ctx context.Context,
	applicantID int32,
	resourceID int32,
	accessType string,
	accessReason string,
) (repo.Request, error) {

	req, err := s.Repo.CreateRequest(ctx, repo.CreateRequestParams{
		ApplicantID:  pgtype.Int4{Int32: applicantID, Valid: true},
		ResourceID:   pgtype.Int4{Int32: resourceID, Valid: true},
		AccessType:   accessType,
		AccessReason: accessReason,
	})
	if err != nil {
		return req, err
	}

	// first approval step is always the applicant's manager
	mgrID, err := s.Repo.GetManagerID(ctx, applicantID)
	if err != nil {
		return req, err
	}
	if mgrID.Valid {
		err = s.Repo.CreateApprovalStep(ctx, repo.CreateApprovalStepParams{
			RequestID:  pgtype.Int4{Int32: req.ID, Valid: true},
			ApproverID: mgrID,
			StepOrder:  1,
			Status:     "pending",
		})
		if err != nil {
			return req, err
		}
	}

	// second approval step is the resource owner
	ownerID, err := s.Repo.GetResourceOwnerID(ctx, resourceID)
	if err != nil {
		return req, err
	}
	if ownerID.Valid {
		err = s.Repo.CreateApprovalStep(ctx, repo.CreateApprovalStepParams{
			RequestID:  pgtype.Int4{Int32: req.ID, Valid: true},
			ApproverID: ownerID,
			StepOrder:  2,
			Status:     "waiting",
		})
		if err != nil {
			return req, err
		}
	}

	// third approval step is the security team
	securityIDs, err := s.Repo.GetSecurityEmployeesID(ctx) // TODO: make it constant in constructor
	if err != nil {
		return req, err
	}
	// TODO: give the request to all security employees until first approval, then remove from others
	secID := rand.Intn(len(securityIDs)) // randomly assign to one security employee
	err = s.Repo.CreateApprovalStep(ctx, repo.CreateApprovalStepParams{
		RequestID:  pgtype.Int4{Int32: req.ID, Valid: true},
		ApproverID: pgtype.Int4{Int32: int32(secID), Valid: true},
		StepOrder:  3,
		Status:     "waiting",
	})
	if err != nil {
		return req, err
	}

	return req, nil
}

func (s *RequestService) ListUserRequests(ctx context.Context, userID int32) ([]repo.ListUserRequestsRow, error) {
	return s.Repo.ListUserRequests(ctx, pgtype.Int4{Int32: userID, Valid: true})
}

func (s *RequestService) DeleteRequest(ctx context.Context, requestID int32) error {
	return s.Repo.DeleteRequest(ctx, requestID)
}
