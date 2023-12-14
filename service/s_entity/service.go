package s_entity

import (
	"clean-architecture/domain/entity"
	"clean-architecture/repository/psql"
	"context"

	"github.com/go-redis/redis"
	"github.com/go-resty/resty/v2"
	"github.com/labstack/echo/v4"
	"github.com/minio/minio-go/v7"
)

type EntityServiceInterface interface {
	GetEntityService(ctx context.Context, id string) (res entity.Entity, err error)
}

type EntityService struct {
	validator   echo.Validator
	redisClient *redis.Client
	samuelRedis *redis.Client
	psqlRepo    *psql.PsqlRepositories
	srClient    *resty.Client
	minio       *minio.Client
	samuelUrl   string
	samuelUrl2  string
}

func NewEntityService(validator echo.Validator, redisClient *redis.Client, samuelRedis *redis.Client, psqlRepo *psql.PsqlRepositories, srClient *resty.Client, minio *minio.Client, samuelUrl, samuelUrl2 string) *EntityService {
	return &EntityService{
		validator:   validator,
		redisClient: redisClient,
		samuelRedis: samuelRedis,
		psqlRepo:    psqlRepo,
		srClient:    srClient,
		minio:       minio,
		samuelUrl:   samuelUrl,
		samuelUrl2:  samuelUrl2,
	}
}
