package utils

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/FourWD/auction.middleware/orm"
	"github.com/FourWD/middleware/common"
)

func UpdateDeductCode(id string) (string, error) {
	var deduct orm.DeductRight
	common.Database.First(&deduct, id)

	// sql := `SELECT * FROM deduct_rights WHERE id = ?`
	// common.Database.Raw(sql, id).Scan(&deduct)

	type Count struct {
		Count int64 `json:"count"`
	}
	var count Count
	year := time.Now().Year()
	month := time.Now().Month()
	day := time.Now().Day()

	sql := `SELECT COUNT(*) 
	FROM deduct_rights 
	WHERE YEAR(created_at) = ?  
	AND MONTH(created_at) = ? AND DAY(created_at) = ? AND running_no <= ?`
	common.Database.Raw(sql, year, month, day, deduct.RunningNo).Scan(&count.Count)
	// common.Print(sql, fmt.Sprintf("%d", count.Count))

	fyear := strconv.Itoa(year)
	fmonth := fmt.Sprintf("%02d", month)
	fday := fmt.Sprintf("%02d", day)
	code := fmt.Sprintf("%s%s%s%s%03d", "DE", fyear[2:4], fmonth, fday, count.Count)

	updateSql := `UPDATE deduct_rights SET code = ? WHERE id = ?`
	if err := common.Database.Exec(updateSql, code, id).Error; err != nil {
		log.Println(err)
	}
	return code, nil
}
