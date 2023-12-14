package psql

import (
	"clean-architecture/domain/entity"
	"clean-architecture/domain/orm"
	etPS "clean-architecture/repository/psql/entity"
	corm "clean-architecture/repository/psql/orm"

	"github.com/jinzhu/gorm"
	gormV2 "gorm.io/gorm"
)

type PsqlRepositories struct {
	DbV2   *gormV2.DB
	Orm    orm.Repository
	Entity entity.Repository
}

func NewPsqlRepositories(db *gorm.DB, dbV2 *gormV2.DB) *PsqlRepositories {
	Orm := corm.New(db)
	return &PsqlRepositories{
		DbV2:   dbV2,
		Orm:    Orm,
		Entity: etPS.New(Orm, dbV2),
	}
}
