package utils

import (
	"fmt"
	"strconv"
	"time"

	"github.com/FourWD/auction.middleware/orm"
	"github.com/FourWD/middleware/common"
)

func UpdateRefundCode(userID string) (string, error) {
	type Code struct {
		orm.Refund
		Code string `json:"code" query:"code"`
	}

	var verifyCode Code

	sql := `SELECT u.id,p.code, rf.running_no 
	FROM users u 
	LEFT JOIN payment_types p ON u.payment_type_id = p.id
	LEFT JOIN refunds rf ON u.id = rf.user_id
	WHERE u.id = ?`
	common.Database.Raw(sql, userID).Scan(&verifyCode)

	type Count struct {
		Count int64 `json:"count"`
	}

	var count Count
	year := time.Now().Year()
	month := time.Now().Month()
	day := time.Now().Day()

	queryRefund := `SELECT COUNT(*) 
	FROM refunds 
	WHERE YEAR(created_at) = ? 
	AND MONTH(created_at) = ? AND DAY(created_at) = ? AND runinng_no <= ? `
	common.Database.Raw(queryRefund, year, month, day, verifyCode.RunningNo).Scan(&count.Count)
	common.Print(queryRefund, fmt.Sprintf("%d", count.Count))

	fyear := strconv.Itoa(year)
	fmonth := fmt.Sprintf("%02d", month)
	fday := fmt.Sprintf("%02d", day)

	verifyRefund := fmt.Sprintf("%s%s%s%s%03d", "RE", fyear[2:4], fmonth, fday, count.Count)
	updateRefund := `UPDATE refunds SET code = ? WHERE id = ?`

	checkUpdate := common.Database.Exec(updateRefund, verifyRefund, userID)

	if checkUpdate.RowsAffected == 0 {
		common.PrintError("Error checkupdate", userID)
	}

	return verifyRefund, nil
}
