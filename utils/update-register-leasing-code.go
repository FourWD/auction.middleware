package utils

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/FourWD/auction.middleware/orm"
	"github.com/FourWD/middleware/common"
)

func UpdateRegisterLeasingCode(id string) (string, error) {
	var leasing orm.RegisterLeasing
	// common.Database.First(&leasing, id)
	sqlID := `SELECT * FROM register_leasings WHERE id = ?`
	common.Database.Raw(sqlID, id).Scan(&leasing)

	type Count struct {
		Count int64 `json:"count"`
	}
	var count Count
	year := time.Now().Year()
	month := time.Now().Month()

	sql := `SELECT COUNT(*) 
	FROM register_leasings 
	WHERE YEAR(created_at) = ? AND MONTH(created_at) = ? AND running_no <= ? `
	common.Database.Raw(sql, year, month, leasing.RunningNo).Scan(&count.Count)
	// common.Print(sql, fmt.Sprintf("%d", count.Count))

	fyear := strconv.Itoa(year)
	fmonth := fmt.Sprintf("%02d", month)
	code := fmt.Sprintf("%s%s%s%s%03d", "FN", leasing.Code, fyear[2:4], fmonth, count.Count)

	updateSql := `UPDATE register_leasings SET code = ? WHERE id = ?`
	if err := common.Database.Exec(updateSql, code, id).Error; err != nil {
		log.Println(err)
	}
	return code, nil
}
