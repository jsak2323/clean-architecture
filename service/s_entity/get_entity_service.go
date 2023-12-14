package s_entity

import (
	"clean-architecture/domain/entity"
	"context"
)

func (s *EntityService) GetEntityService(ctx context.Context, id string) (res entity.Entity, err error) {
	err = s.psqlRepo.Entity.GetEntity(ctx, "")
	if err != nil {
		return res, err
	}
	return res, err
}
