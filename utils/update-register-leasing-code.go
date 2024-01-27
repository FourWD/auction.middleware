package utils

import (
	"fmt"
	"strconv"
	"time"

	"github.com/FourWD/auction.middleware/orm"
	"github.com/FourWD/middleware/common"
)

func UpdateRegisterLeasingCode(id string) (string, error) {

	type Code struct {
		RegisterLeasing orm.RegisterLeasing

		Financeid string `json:"financeid" query:"financeid"`
		Code      string `json:"code" query:"code"`
	}

	var verifyCode Code

	sql := fmt.Sprintf(`SELECT rl.running_no, f.label as code, rl.finance_id as financeid
	FROM register_leasings rl
	LEFT JOIN users u ON rl.user_id = u.id
	   LEFT JOIN finances f ON rl.finance_id = f.id
	WHERE rl.id = '%s'`, id)

	common.Database.Raw(sql).Scan(&verifyCode)
	common.Print(sql, verifyCode.Financeid)
	type Count struct {
		Count int64 `json:"count"`
	}

	var count Count
	year := time.Now().Year()
	month := time.Now().Month()

	queryBanktransfer := fmt.Sprintf(`SELECT COUNT(*) FROM register_leasings WHERE YEAR(created_at) = %d AND MONTH(created_at) = %d `, year, month)
	common.Database.Raw(queryBanktransfer).Scan(&count.Count)
	common.Print(queryBanktransfer, fmt.Sprintf("%d", count.Count))

	fyear := strconv.Itoa(year)
	fmonth := fmt.Sprintf("%02d", month)

	verifyLeasing := fmt.Sprintf("%s%s%s%s%03d", "FN", verifyCode.Financeid, fyear[2:4], fmonth, count.Count)

	updateBanktransfer := fmt.Sprintf(`UPDATE register_leasings SET code = '%s' WHERE id = '%s'`, verifyLeasing, id)
	checkupdate := common.Database.Exec(updateBanktransfer)

	if checkupdate.RowsAffected == 0 {
		errMsg := fmt.Sprintf("Error updating register leasing code for id %s", id)
		common.PrintError(errMsg, updateBanktransfer)
	}

	return verifyLeasing, nil
}
