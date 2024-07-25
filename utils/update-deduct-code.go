package utils

import (
	"fmt"
	"strconv"
	"time"

	"github.com/FourWD/auction.middleware/orm"
	"github.com/FourWD/middleware/common"
)

func UpdateDeductCode(userID string) (string, error) {
	type Code struct {
		orm.DeductRight
		Code string `json:"code" query:"code"`
	}

	var verifyCode Code

	sql := `SELECT u.id,p.code, dr.running_no FROM users u 
	LEFT JOIN payment_types p ON u.payment_type_id = p.id
	LEFT JOIN deduct_rights dr ON u.id = dr.user_id
	WHERE u.id = ?`
	common.Database.Raw(sql, userID).Scan(&verifyCode)

	type Count struct {
		Count int64 `json:"count"`
	}

	var count Count
	year := time.Now().Year()
	month := time.Now().Month()
	day := time.Now().Day()

	queryDeduct := `SELECT COUNT(*) 
	FROM deduct_rights 
	WHERE YEAR(created_at) = ?  
	AND MONTH(created_at) = ? AND DAY(created_at) = ? AND runinng_no <= ?`
	common.Database.Raw(queryDeduct, year, month, day, verifyCode.RunningNo).Scan(&count.Count)
	common.Print(queryDeduct, fmt.Sprintf("%d", count.Count))

	fyear := strconv.Itoa(year)
	fmonth := fmt.Sprintf("%02d", month)

	fday := fmt.Sprintf("%02d", day)

	verifyDeduct := fmt.Sprintf("%s%s%s%s%03d", "DE", fyear[2:4], fmonth, fday, count.Count)
	updateDeduct := `UPDATE deduct_rights 
	SET code = ? WHERE id = ?`

	checkUpdate := common.Database.Exec(updateDeduct, verifyDeduct, userID)

	if checkUpdate.RowsAffected == 0 {
		common.PrintError("Error checkupdate", userID)
	}

	return verifyDeduct, nil
}
