package helpers

import "github.com/LUSHDigital/microservice-core-golang/pagination"

type Pagination struct {
	PerPage     int
	CurrentPage int
	TotalRecord int
	TotalPage   int
	LastPage    int
}

func GeneratePagination(page, perPage, totalRow int) (Pagination, error) {
	pagination, err := pagination.NewPaginator(perPage, page, totalRow)
	return Pagination{
		PerPage:     pagination.GetPerPage(),
		CurrentPage: pagination.GetPage(),
		TotalRecord: totalRow,
		TotalPage:   pagination.GetLastPage(),
		LastPage:    pagination.GetLastPage(),
	}, err
}
