package utils

import (
	"errors"

	"github.com/FourWD/middleware/common"
)

type UserFirebase struct {
	ID              string `json:"id" query:"id" firestore:"id"`
	RightDeposit    int    `json:"right_deposit" query:"right_deposit" firestore:"right_deposit"`
	PaymentTypeID   string `json:"payment_type_id" query:"payment_type_id" firestore:"payment_type_id" `
	UserDisplayName string `json:"user_display_name" query:"user_display_name" firestore:"user_display_name"`
}

func SyncUser(userID string) error {

	if userID == "" {
		return errors.New("wrong user")
	}

	user := getUserMy(userID)

	if _, err := common.FirebaseClient.Collection("users").Doc(userID).Set(common.FirebaseCtx, &user); err != nil {
		return err
	}

	return nil
}

func getUserMy(userID string) UserFirebase {
	sql := `SELECT u.id , u.payment_type_id , u.right_deposit,
	IF(u.user_type_id = '01',CONCAT_WS(" ", u.firstname,u.lastname),up.company_name) as user_display_name
	FROM user u
	LEFT JOIN user_profiles up ON u.id = up.id
	WHERE u.id = ?`
	var user = new(UserFirebase)
	common.Database.Raw(sql, userID).Scan(&user)
	return *user
}
