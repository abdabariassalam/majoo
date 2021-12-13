package service

import (
	"github.com/abdabariassalam/majoo/config"
	"github.com/abdabariassalam/majoo/internal/repository"
)

type service struct {
	cfg  config.Config
	repo repository.Repository
}

type Service interface {
	VerifyToken(reqToken string) (*TokenClaims, error)
	Login(userName, password string) (string, error)
	MerchantOmzetReport(userID int64, month, year, page, perPage int) (*MerchantOmzetReportResponse, error)
	OutletOmzetReport(userID int64, month, year, page, perPage int) (*OtletOmzetReportResponse, error)
}

func New(cfg config.Config, repo repository.Repository) Service {
	return service{
		cfg:  cfg,
		repo: repo,
	}
}
