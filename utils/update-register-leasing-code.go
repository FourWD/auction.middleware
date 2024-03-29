package utils

import (
	"fmt"
	"strconv"
	"time"

	"github.com/FourWD/auction.middleware/orm"
	"github.com/FourWD/middleware/common"
)

func GenCode(userID string) (string, error) {

	type CustomerCode struct {
		orm.User
		UserTypeCode string `json:"user_type_code" query:"user_type_code"`
	}
	var user CustomerCode
	common.PrintError("GencodeUserID", userID)
	sql := fmt.Sprintf(`SELECT users.running_no, users.user_type_id, user_types.code as user_type_code
							FROM users
						LEFT JOIN user_types ON users.user_type_id = user_types.id
						WHERE users.id = '%s' `, userID)

	common.Database.Raw(sql).Debug().Scan(&user)
	common.Print("struct", common.StructToString(user))

	type Count struct {
		Count int64 `json:"count"`
	}

	var count Count
	year := time.Now().Year()
	month := time.Now().Month()
	// day := time.Now().Day()

	query := fmt.Sprintf(`SELECT COUNT(*) FROM users WHERE YEAR(created_at) = %d AND MONTH(created_at) = %d
	AND running_no <= %d AND user_type_id = '%s'`, year, month, user.RunningNo, user.UserTypeID)
	// query := fmt.Sprintf(`SELECT COUNT(*) FROM users WHERE YEAR(created_at) = %d AND MONTH(created_at) = %d
	//  AND user_type_id = '%s'`, year, month, user.UserTypeID)
	common.Database.Raw(query).Debug().Scan(&count.Count)
	common.Print(query, fmt.Sprintf("%d", count.Count))

	fyear := strconv.Itoa(year)
	fmonth := fmt.Sprintf("%02d", month)
	// fday := fmt.Sprintf("%02d", day)
	code := fmt.Sprintf("%s%s%s%04d", user.UserTypeCode, fyear[2:4], fmonth, count.Count)

	update := fmt.Sprintf(`UPDATE users set code = '%s' where id = '%s'`, code, userID)
	checkupdate := common.Database.Debug().Exec(update)
	if checkupdate.RowsAffected == 0 {
		common.PrintError("Error checkupdate", userID)
	}
	common.Print("check", code)
	return code, nil
}
func UpdateBanktrasferCode(userID string) (string, error) {
	type Code struct {
		orm.BankTransfer
		Code string `json:"code" query:"code"`
	}

	var verifyCode Code

	sql := fmt.Sprintf(`SELECT u.id,p.code, bt.running_no FROM users u 
							LEFT JOIN payment_types p ON u.payment_type_id = p.id
							LEFT JOIN bank_transfers bt ON u.id = bt.user_id
						WHERE u.id = '%s'`, userID)

	common.Database.Raw(sql).Scan(&verifyCode)
	common.Print(sql, "verifycode")
	type Count struct {
		Count int64 `json:"count"`
	}

	var count Count
	year := time.Now().Year()
	month := time.Now().Month()
	day := time.Now().Day()

	queryBanktransfer := fmt.Sprintf(`SELECT COUNT(*) FROM bank_transfers WHERE YEAR(created_at) = %d AND MONTH(created_at) = %d AND DAY(created_at) = %d `, year, month, day)
	common.Database.Raw(queryBanktransfer).Scan(&count.Count)
	common.Print(queryBanktransfer, fmt.Sprintf("%d", count.Count))

	fyear := strconv.Itoa(year)
	fmonth := fmt.Sprintf("%02d", month)
	fday := fmt.Sprintf("%02d", day)

	verifyBanktransfer := fmt.Sprintf("%s%s%s%s%03d", "CS", fyear[2:4], fmonth, fday, count.Count)

	// date := common.NilDate()

	updateBanktransfer := fmt.Sprintf(`UPDATE bank_transfers SET code = '%s' , updated_at = '%s' WHERE id = '%s'`, verifyBanktransfer, common.NilDate(), userID)

	checkUpdate := common.Database.Exec(updateBanktransfer)

	if checkUpdate.RowsAffected == 0 {
		common.PrintError("Error checkupdate", userID)
	}

	return verifyBanktransfer, nil
}
func UpdateRefundCode(userID string) (string, error) {
	type Code struct {
		orm.Refund
		Code string `json:"code" query:"code"`
	}

	var verifyCode Code

	sql := fmt.Sprintf(`SELECT u.id,p.code, rf.running_no FROM users u 
							LEFT JOIN payment_types p ON u.payment_type_id = p.id
							LEFT JOIN refunds rf ON u.id = rf.user_id
						WHERE u.id = '%s'`, userID)

	common.Database.Raw(sql).Scan(&verifyCode)

	type Count struct {
		Count int64 `json:"count"`
	}

	var count Count
	year := time.Now().Year()
	month := time.Now().Month()
	day := time.Now().Day()

	queryRefund := fmt.Sprintf(`SELECT COUNT(*) FROM refunds WHERE YEAR(created_at) = %d  AND MONTH(created_at) = %d AND DAY(created_at) = %d `, year, month, day)
	common.Database.Raw(queryRefund).Scan(&count.Count)
	common.Print(queryRefund, fmt.Sprintf("%d", count.Count))

	fyear := strconv.Itoa(year)
	fmonth := fmt.Sprintf("%02d", month)
	fday := fmt.Sprintf("%02d", day)

	verifyRefund := fmt.Sprintf("%s%s%s%s%03d", "RE", fyear[2:4], fmonth, fday, count.Count)
	updateRefund := fmt.Sprintf(`UPDATE refunds SET code = '%s' WHERE id = '%s'`, verifyRefund, userID)

	checkUpdate := common.Database.Exec(updateRefund)

	if checkUpdate.RowsAffected == 0 {
		common.PrintError("Error checkupdate", userID)
	}

	return verifyRefund, nil
}

func UpdateDeductCode(userID string) (string, error) {
	type Code struct {
		orm.DeductRight
		Code string `json:"code" query:"code"`
	}

	var verifyCode Code

	sql := fmt.Sprintf(`SELECT u.id,p.code, dr.running_no FROM users u 
							LEFT JOIN payment_types p ON u.payment_type_id = p.id
							LEFT JOIN deduct_rights dr ON u.id = dr.user_id
						WHERE u.id = '%s'`, userID)

	common.Database.Raw(sql).Scan(&verifyCode)

	type Count struct {
		Count int64 `json:"count"`
	}

	var count Count
	year := time.Now().Year()
	month := time.Now().Month()
	day := time.Now().Day()

	queryDeduct := fmt.Sprintf(`SELECT COUNT(*) FROM deduct_rights WHERE YEAR(created_at) = %d  AND MONTH(created_at) = %d AND DAY(created_at) = %d `, year, month, day)
	common.Database.Raw(queryDeduct).Scan(&count.Count)
	common.Print(queryDeduct, fmt.Sprintf("%d", count.Count))

	fyear := strconv.Itoa(year)
	fmonth := fmt.Sprintf("%02d", month)

	fday := fmt.Sprintf("%02d", day)

	verifyDeduct := fmt.Sprintf("%s%s%s%s%03d", "DE", fyear[2:4], fmonth, fday, count.Count)
	updateDeduct := fmt.Sprintf(`UPDATE deduct_rights SET code = '%s' WHERE id = '%s'`, verifyDeduct, userID)

	checkUpdate := common.Database.Exec(updateDeduct)

	if checkUpdate.RowsAffected == 0 {
		common.PrintError("Error checkupdate", userID)
	}

	return verifyDeduct, nil
}

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

	queryLeasing := fmt.Sprintf(`SELECT COUNT(*) FROM register_leasings WHERE YEAR(created_at) = %d AND MONTH(created_at) = %d `, year, month)
	common.Database.Raw(queryLeasing).Scan(&count.Count)
	common.Print(queryLeasing, fmt.Sprintf("%d", count.Count))

	fyear := strconv.Itoa(year)
	fmonth := fmt.Sprintf("%02d", month)

	verifyLeasing := fmt.Sprintf("%s%s%s%s%03d", "FN", verifyCode.Code, fyear[2:4], fmonth, count.Count)

	updateBanktransfer := fmt.Sprintf(`UPDATE register_leasings SET code = '%s' WHERE id = '%s'`, verifyLeasing, id)
	checkupdate := common.Database.Exec(updateBanktransfer)
	if checkupdate.RowsAffected == 0 {
		common.PrintError("Error checkupdate", id)
	}
	return verifyLeasing, nil
}
