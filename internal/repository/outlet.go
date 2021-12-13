package repository

import (
	"time"

	"github.com/abdabariassalam/majoo/internal/models"
)

func (r repository) OutletOmzetReport(userID int64, month, year, page, perPage int) (report []*models.OutletMonthReport, totalRow int64, err error) {
	offset := (page - 1) * perPage
	firstOfMonth := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Local)
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)
	dbResp := r.db.Table("Transactions").Select("r.tgl, r.merchant_name, o.outlet_name, COALESCE(SUM(Transactions.bill_total),0) as omzet").Joins(`
        right join (SELECT dr.merchant_id, dr.user_id, dr.merchant_name, dr.tgl
        from (SELECT ADDDATE(?, row-1) as tgl, m.user_id, m.merchant_name, m.id as merchant_id 
        from (SELECT @row := @row + 1 as row
        FROM (select 0 union all select 1 union all select 3 union all select 4) t,
        (select 0 union all select 1 union all select 3 union all select 4 union all select 5 union all select 6 union all select 7 union all select 8) t2, (SELECT @row:=0) r) n
        join Merchants m 
        where n.row <= ?) dr
        ORDER by dr.merchant_id, dr.tgl) r ON r.merchant_id = Transactions.merchant_id AND r.tgl = date(Transactions.created_at)
        inner join Outlets o on r.merchant_id = o.merchant_id 
        WHERE r.user_id = ?
        GROUP by r.tgl, r.user_id, r.merchant_id, r.merchant_name, o.outlet_name`, firstOfMonth.Format("2006-01-02"), int(lastOfMonth.Sub(firstOfMonth).Hours()/24), userID)
	dbResp.Count(&totalRow)
	dbResp.Limit(perPage).Offset(offset)
	rows, err := dbResp.Rows()
	if err != nil {
		return nil, 0, err
	}
	for rows.Next() {
		temp := &models.OutletMonthReport{}
		err = rows.Scan(
			&temp.Date,
			&temp.MerchantName,
			&temp.OtletName,
			&temp.Omzet)
		if err != nil {
			return nil, 0, err
		}
		report = append(report, temp)
		// do something
	}
	return report, totalRow, err
}
