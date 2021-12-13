package repository

import (
	"github.com/abdabariassalam/majoo/config"
	"github.com/abdabariassalam/majoo/internal/models"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

type Repository interface {
	FindByNameAndPassword(userName, password string) (*models.User, error)
	MerchantOmzetReport(userID int64, month, year, page, perPage int) (report []*models.MerchantMonthReport, totalRow int64, err error)
	OutletOmzetReport(userID int64, month, year, page, perPage int) (report []*models.OutletMonthReport, totalRow int64, err error)
}

type Args struct {
	DB  *gorm.DB
	Cfg *config.Config
}

func New(a *Args) Repository {
	return repository{
		db: a.DB,
	}
}
