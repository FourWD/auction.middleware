package utils

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/FourWD/auction.middleware/orm"
	"github.com/FourWD/middleware/common"
)

func UpdateRefundCode(id string) (string, error) {
	var refund orm.Refund
	common.Database.First(&refund, id)

	// sql := `SELECT * refunds WHERE id = ?`
	// common.Database.Raw(sql, userID).Scan(&refund)

	type Count struct {
		Count int64 `json:"count"`
	}
	var count Count
	year := time.Now().Year()
	month := time.Now().Month()
	day := time.Now().Day()

	sql := `SELECT COUNT(*) 
	FROM refunds 
	WHERE YEAR(created_at) = ? 
	AND MONTH(created_at) = ? AND DAY(created_at) = ? AND running_no <= ? `
	common.Database.Raw(sql, year, month, day, refund.RunningNo).Scan(&count.Count)
	// common.Print(sql, fmt.Sprintf("%d", count.Count))

	fyear := strconv.Itoa(year)
	fmonth := fmt.Sprintf("%02d", month)
	fday := fmt.Sprintf("%02d", day)
	code := fmt.Sprintf("%s%s%s%s%03d", "RE", fyear[2:4], fmonth, fday, count.Count)

	updateSql := `UPDATE refunds SET code = ? WHERE id = ?`
	if err := common.Database.Exec(updateSql, code, id).Error; err != nil {
		log.Println(err)
	}
	return code, nil
}
