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
	sql := `SELECT users.running_no, users.user_type_id, user_types.code as user_type_code
							FROM users
						LEFT JOIN user_types ON users.user_type_id = user_types.id
						WHERE users.id = ?`

	common.Database.Raw(sql, userID).Scan(&user)
	common.Print("struct", common.StructToString(user))

	type Count struct {
		Count int64 `json:"count"`
	}

	var count Count
	year := time.Now().Year()
	month := time.Now().Month()
	// day := time.Now().Day()

	query := `SELECT COUNT(*) 
	FROM users 
	WHERE YEAR(created_at) = ? AND MONTH(created_at) = ?
	AND running_no <= ? AND user_type_id = ?`
	// query := fmt.Sprintf(`SELECT COUNT(*) FROM users WHERE YEAR(created_at) = %d AND MONTH(created_at) = %d
	//  AND user_type_id = '%s'`, year, month, user.UserTypeID)
	common.Database.Raw(query, year, month, user.RunningNo, user.UserTypeID).Scan(&count.Count)
	common.Print(query, fmt.Sprintf("%d", count.Count))

	fyear := strconv.Itoa(year)
	fmonth := fmt.Sprintf("%02d", month)
	// fday := fmt.Sprintf("%02d", day)
	code := fmt.Sprintf("%s%s%s%04d", user.UserTypeCode, fyear[2:4], fmonth, count.Count)

	update := `UPDATE users set code = ? where id = ?`
	checkupdate := common.Database.Exec(update, code, userID)
	if checkupdate.RowsAffected == 0 {
		common.PrintError("Error checkupdate", userID)
	}
	common.Print("check", code)
	return code, nil
}
