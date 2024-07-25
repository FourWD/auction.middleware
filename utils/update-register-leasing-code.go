package utils

import (
	"fmt"
	"strconv"
	"time"

	"github.com/FourWD/middleware/common"
)

func UpdateRegisterLeasingCode(id string) (string, error) {

	type Code struct {
		// RegisterLeasing orm.RegisterLeasing

		Financeid string `json:"financeid" query:"financeid"`
		Code      string `json:"code" query:"code"`
		RunningNo int    `json:"running_no" query:"running_no"`
	}

	var verifyCode Code

	sql := `SELECT rl.running_no, f.label as code, rl.finance_id as financeid
	FROM register_leasings rl
	LEFT JOIN users u ON rl.user_id = u.id
	LEFT JOIN finances f ON rl.finance_id = f.id
	WHERE rl.id = ?`

	common.Database.Raw(sql, id).Scan(&verifyCode)
	//common.Print(sql, verifyCode.Financeid)
	type Count struct {
		Count int64 `json:"count"`
	}

	var count Count
	year := time.Now().Year()
	month := time.Now().Month()

	println(verifyCode.RunningNo)
	queryLeasing := `SELECT COUNT(*) 
	FROM register_leasings 
	WHERE YEAR(created_at) = ? AND MONTH(created_at) = ? AND runinng_no <= ? `
	common.Database.Raw(queryLeasing, year, month, verifyCode.RunningNo).Scan(&count.Count)
	common.Print(queryLeasing, fmt.Sprintf("%d", count.Count))

	fyear := strconv.Itoa(year)
	fmonth := fmt.Sprintf("%02d", month)

	verifyLeasing := fmt.Sprintf("%s%s%s%s%03d", "FN", verifyCode.Code, fyear[2:4], fmonth, count.Count)

	updateBanktransfer := `UPDATE register_leasings 
	SET code = ? WHERE id = ?`
	checkupdate := common.Database.Exec(updateBanktransfer, verifyLeasing, id)
	if checkupdate.RowsAffected == 0 {
		common.PrintError("Error checkupdate", id)
	}
	return verifyLeasing, nil
}
