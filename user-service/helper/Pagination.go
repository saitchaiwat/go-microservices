package helper

import (
	"fmt"
	"gorm.io/gorm"
	"math"
)

type Paginator struct {
	Limit      int         `json:"limit"`
	Page       int         `json:"page"`
	TotalRows  int64       `json:"totalRows"`
	TotalPages int         `json:"totalPages"`
	Rows       interface{} `json:"rows"`
}

func Paginate(db *gorm.DB, countSQL string, rawSQL string, limit int, page int, named map[string]interface{}, rows interface{}) (paginator Paginator, err error) {
	var totalRows int64

	if limit <= 0 {
		limit = 10
	}

	if page <= 0 {
		page = 1
	}

	page = page - 1

	if len(named) == 0 {
		db.Raw(countSQL).Scan(&totalRows)
	} else {
		db.Raw(countSQL, named).Scan(&totalRows)
	}

	named["limit"] = limit
	named["offset"] = page

	sql := fmt.Sprintf("%s LIMIT @limit OFFSET @offset", rawSQL)

	if err := db.Raw(sql, named).Scan(&rows).Error; err != nil {
		return paginator, err
	}

	totalPages := int(math.Ceil(float64(totalRows) / float64(limit)))

	paginator.Limit = limit
	paginator.Page = page
	paginator.TotalRows = totalRows
	paginator.TotalPages = totalPages
	paginator.Rows = rows
	return paginator, nil
}
