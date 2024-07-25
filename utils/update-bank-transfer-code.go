package utils

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/FourWD/auction.middleware/orm"
	"github.com/FourWD/middleware/common"
)

func UpdateBankTransferCode(id string) (string, error) {
	var bank orm.BankTransfer
	// common.Database.First(&bank, id)
	sqlID := `SELECT * FROM bank_transfers WHERE id = ?`
	common.Database.Raw(sqlID, id).Scan(&bank)

	type Count struct {
		Count int64 `json:"count"`
	}
	var count Count
	year := time.Now().Year()
	month := time.Now().Month()
	day := time.Now().Day()

	sql := `SELECT COUNT(*) 
	FROM bank_transfers WHERE YEAR(created_at) = ? 
	AND MONTH(created_at) = ? AND DAY(created_at) = ? AND running_no <= ?`
	common.Database.Raw(sql, year, month, day, bank.RunningNo).Scan(&count.Count)

	fyear := strconv.Itoa(year)
	fmonth := fmt.Sprintf("%02d", month)
	fday := fmt.Sprintf("%02d", day)
	code := fmt.Sprintf("%s%s%s%s%03d", "CS", fyear[2:4], fmonth, fday, count.Count)

	updateSql := `UPDATE bank_transfers SET code = ? WHERE id = ?`
	if err := common.Database.Exec(updateSql, code, id).Error; err != nil {
		log.Println(err)
	}
	return code, nil
}
