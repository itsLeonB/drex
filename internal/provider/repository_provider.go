package provider

import (
	"github.com/itsLeonB/drex/internal/entity"
	"github.com/itsLeonB/drex/internal/repository"
	"github.com/itsLeonB/go-crud"
	"gorm.io/gorm"
)

type Repositories struct {
	DebtTransaction repository.DebtTransactionRepository
	TransferMethod  repository.TransferMethodRepository
}

func ProvideRepositories(gormDB *gorm.DB) *Repositories {
	if gormDB == nil {
		panic("gormDB cannot be nil")
	}

	return &Repositories{
		DebtTransaction: repository.NewDebtTransactionRepository(gormDB),
		TransferMethod:  crud.NewRepository[entity.TransferMethod](gormDB),
	}
}
