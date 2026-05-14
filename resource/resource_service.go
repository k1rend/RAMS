package resource

import (
	"context"

	"github.com/k1rend/RAMS/repo"
	"github.com/jackc/pgx/v5/pgtype"
)

type ResourceService struct {
	repo *repo.Queries
}

func NewResourceService(repo *repo.Queries) *ResourceService {
	return &ResourceService{repo: repo}
}

func (s *ResourceService) CreateResource(ctx context.Context, name, description, resource, resourceType string, ownerID int32) error {
	return s.repo.CreateResource(ctx, repo.CreateResourceParams{
		Name:         name,
		Description:  description,
		Resource:     resource,
		ResourceType: resourceType,
		OwnerID:      pgtype.Int4{Int32: ownerID, Valid: true},
	})
}

func (s *ResourceService) ListResources(ctx context.Context) ([]repo.ListResourcesRow, error) {
	return s.repo.ListResources(ctx)
}	