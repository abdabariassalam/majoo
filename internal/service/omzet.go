package service

import (
	"github.com/abdabariassalam/majoo/internal/helpers"
	"github.com/abdabariassalam/majoo/internal/models"
)

type MerchantOmzetReportResponse struct {
	Report     []*models.MerchantMonthReport
	Pagination helpers.Pagination
}

func (s service) MerchantOmzetReport(userID int64, month, year, page, perPage int) (*MerchantOmzetReportResponse, error) {
	report, totalRow, err := s.repo.MerchantOmzetReport(userID, month, year, page, perPage)
	if err != nil {
		return nil, err
	}
	pagination, err := helpers.GeneratePagination(page, perPage, int(totalRow))
	return &MerchantOmzetReportResponse{
		Report:     report,
		Pagination: pagination,
	}, err
}

type OtletOmzetReportResponse struct {
	Report     []*models.OutletMonthReport
	Pagination helpers.Pagination
}

func (s service) OutletOmzetReport(userID int64, month, year, page, perPage int) (*OtletOmzetReportResponse, error) {
	report, totalRow, err := s.repo.OutletOmzetReport(userID, month, year, page, perPage)
	if err != nil {
		return nil, err
	}
	pagination, err := helpers.GeneratePagination(page, perPage, int(totalRow))
	return &OtletOmzetReportResponse{
		Report:     report,
		Pagination: pagination,
	}, err
}
