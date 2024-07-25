package utils

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/FourWD/auction.middleware/orm"
	"github.com/FourWD/middleware/common"
)

func UpdateUserCode(id string) (string, error) {
	type User struct {
		orm.User
		UserTypeCode string `json:"user_type_code" query:"user_type_code"`
	}
	var user User
	// common.Database.First(&user, id)
	sqlID := `SELECT * FROM users WHERE id = ?`
	common.Database.Raw(sqlID, id).Debug().Scan(&user)

	type Count struct {
		Count int64 `json:"count"`
	}
	var count Count
	year := time.Now().Year()
	month := time.Now().Month()

	sql := `SELECT COUNT(*) 
	FROM users 
	WHERE YEAR(created_at) = ? AND MONTH(created_at) = ?
	AND user_type_id = ? AND running_no <= ?`
	common.Database.Raw(sql, year, month, user.UserTypeID, user.RunningNo).Debug().Scan(&count.Count)
	// common.Print(sql, fmt.Sprintf("%d", count.Count))

	fyear := strconv.Itoa(year)
	fmonth := fmt.Sprintf("%02d", month)
	code := fmt.Sprintf("%s%s%s%04d", user.UserTypeCode, fyear[2:4], fmonth, count.Count)

	updateSql := `UPDATE users set code = ? where id = ?`
	if err := common.Database.Debug().Exec(updateSql, code, id).Error; err != nil {
		log.Println(err)
	}
	return code, nil
}
