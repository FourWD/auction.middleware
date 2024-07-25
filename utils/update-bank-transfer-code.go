package utils

import (
	"fmt"
	"strconv"
	"time"

	"github.com/FourWD/auction.middleware/orm"
	"github.com/FourWD/middleware/common"
)

func UpdateBanktrasferCode(userID string) (string, error) {
	type Code struct {
		orm.BankTransfer
		Code string `json:"code" query:"code"`
	}

	var verifyCode Code

	sql := `SELECT u.id,p.code, bt.running_no FROM users u 
							LEFT JOIN payment_types p ON u.payment_type_id = p.id
							LEFT JOIN bank_transfers bt ON u.id = bt.user_id
						WHERE u.id = ?`
	common.Database.Raw(sql, userID).Scan(&verifyCode)
	common.Print(sql, "verifycode")
	type Count struct {
		Count int64 `json:"count"`
	}

	var count Count
	year := time.Now().Year()
	month := time.Now().Month()
	day := time.Now().Day()

	queryBanktransfer := `SELECT COUNT(*) 
	FROM bank_transfers WHERE YEAR(created_at) = ? 
	AND MONTH(created_at) = ? AND DAY(created_at) = ? AND runinng_no <= ?`
	common.Database.Raw(queryBanktransfer, year, month, day, verifyCode.RunningNo).Scan(&count.Count)
	//common.Print(queryBanktransfer, fmt.Sprintf("%d", count.Count))

	fyear := strconv.Itoa(year)
	fmonth := fmt.Sprintf("%02d", month)
	fday := fmt.Sprintf("%02d", day)

	verifyBanktransfer := fmt.Sprintf("%s%s%s%s%03d", "CS", fyear[2:4], fmonth, fday, count.Count)

	// date := common.NilDate()

	updateBanktransfer := `UPDATE bank_transfers 
	SET code = ?, updated_at = '` + "1900-01-01 06:42:04.000" + `' WHERE id = ?`

	checkUpdate := common.Database.Exec(updateBanktransfer, verifyBanktransfer, userID)

	if checkUpdate.RowsAffected == 0 {
		common.PrintError("Error checkupdate", userID)
	}

	return verifyBanktransfer, nil
}
