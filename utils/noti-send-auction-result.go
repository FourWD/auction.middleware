package utils

import (
	"fmt"
	"time"

	midOrm "github.com/FourWD/middleware/orm"
	"github.com/google/uuid"

	"github.com/FourWD/middleware/common"
)

// func NotiSendAuctionResult(auctionID string) error {
// 	title := "ประกาศผล"
// 	body := "กดที่นี่เพื่อดูผลประมูล"

// 	data := map[string]string{
// 		"auction_id": auctionID,
// 		"event_code": "R0001",
// 	}
// 	common.PrintError(common.StructToString(data), "tdasdasdasdas")

// 	topics := fmt.Sprintf(`JOIN_AUCTION_%s`, auctionID)
// 	common.PrintError(topics, "topicstopicstopicstopicstopicstopicstopicstopics")

// 	if errSendToSubscriber := common.SendMessageToSubscriber(topics, title, body, data); errSendToSubscriber != nil {
// 		common.PrintError("errSendToSubscriber", errSendToSubscriber.Error())
// 	}
// 	topicid := ""
// 	sql := `SELECT id FROM notification_topics WHERE name = ?`
// 	common.Database.Raw(sql, topics).Scan(&topicid)
// 	common.PrintError(topicid, "topicid")
// 	common.PrintError(sql, "sqlqqqqqqqqqqqqqqqqqqqqqqqqq")

// 	notificationresult := midOrm.Notification{
// 		ID:                    uuid.NewString(),
// 		ToNotificationTopicID: topicid,
// 		NotificationTypeID:    "02",
// 		Message:               body,
// 		Title:                 title,
// 		ShowDate:              time.Now(),
// 	}
// 	common.PrintError(common.StructToString(notificationresult), "notificationresultnotificationresultnotificationresultnotificationresultnotificationresultnotificationresultnotificationresultnotificationresult")

// 	if err := common.Database.Debug().Create(&notificationresult).Error; err != nil {
// 		return fmt.Errorf("failed to insert notificationresult: %v", err)
// 	}

//		return nil
//	}
func NotiSendAuctionResult(auctionID string, userID string, notificationToken string) error {
	title := "ประกาศผล"

	type AuctionDetails struct {
		AuctionID    string    `json:"auction_id"`
		AuctionName  string    `json:"auction_name"`
		RoundName    string    `json:"round_name"`
		CountVehicle int       `json:"count_vehicle"`
		WinnerUserID string    `json:"winner_user_id"`
		AuctionDate  time.Time `json:"auction_date"`
	}
	var auctionDetails AuctionDetails

	sql := `SELECT 
				a.id as auction_id, 
				a.name as auction_name, 
				r.name as round_name, 
				(SELECT COUNT(*) FROM auction_vehicles av2 WHERE av2.auction_id = a.id) as count_vehicle, 
				av.winner_user_id,
				a.end_date
			FROM 
				auctions a
			LEFT JOIN 
				auction_vehicles av ON a.id = av.auction_id 
			LEFT JOIN 
				rounds r ON a.round_id = r.id 
			WHERE 
				a.id = ?
			GROUP BY 
				a.id, a.name, r.name, av.winner_user_id, a.end_date`

	if err := common.Database.Raw(sql, auctionID).Scan(&auctionDetails).Error; err != nil {
		common.PrintError("Error retrieving auction details", err.Error())
		return fmt.Errorf("failed to retrieve auction details: %v", err)
	}

	common.PrintError(auctionDetails.AuctionName, "Auction details")

	// Fetch all users who participated in the auction from auction_vehicle_users table
	type Participant struct {
		UserID string
	}
	var participants []Participant
	sql = `SELECT user_id FROM auction_vehicle_users WHERE auction_id = ?`
	if err := common.Database.Raw(sql, auctionID).Scan(&participants).Error; err != nil {
		common.PrintError("Error retrieving auction participants", err.Error())
		return fmt.Errorf("failed to retrieve auction participants: %v", err)
	}

	// Loop through all participants and send notifications
	for _, participant := range participants {
		var body string
		if auctionDetails.WinnerUserID == participant.UserID {
			body = fmt.Sprintf("🎉 ยินดีด้วย คุณชนะการประมูล\nรอบการประมูลวันที่ %s รอบ %s จํานวนรถ %d คัน",
				auctionDetails.AuctionDate.Format("2 มกราคม 2006"),
				auctionDetails.RoundName,
				auctionDetails.CountVehicle)
		} else {
			body = fmt.Sprintf("🙏 Omakase Car Auction ขอขอบคุณ\nท่านไม่มีรถที่ชนะการประมูลในรอบการประมูลวันที่ %s รอบ %s พบกันใหม่ในการประมูลครั้งหน้านะคะ",
				auctionDetails.AuctionDate.Format("2 มกราคม 2006"),
				auctionDetails.RoundName)
		}

		data := map[string]string{
			"auction_id": auctionID,
			"event_code": "R0001",
		}
		common.PrintError(common.StructToString(data), "Notification data")

		// Construct topic name
		topics := fmt.Sprintf(`JOIN_AUCTION_%s`, auctionID)
		common.PrintError(topics, "Notification topic")

		if errSendToSubscriber := common.SendMessageToUser(notificationToken, title, body, data); errSendToSubscriber != nil {
			common.PrintError("Error sending to subscriber", errSendToSubscriber.Error())
			return fmt.Errorf("failed to send message to subscriber: %v", errSendToSubscriber)
		}

		var topicID string
		sql = `SELECT id FROM notification_topics WHERE name = ?`
		if err := common.Database.Raw(sql, topics).Scan(&topicID).Error; err != nil {
			common.PrintError("Error retrieving topic ID", err.Error())
			return fmt.Errorf("failed to retrieve topic ID: %v", err)
		}
		common.PrintError(topicID, "Retrieved topic ID")
		common.PrintError(sql, "SQL query")

		notificationResult := midOrm.Notification{
			ID:                    uuid.NewString(),
			ToNotificationTopicID: topicID,
			NotificationTypeID:    "02",
			Message:               body,
			Title:                 title,
			ShowDate:              time.Now(),
		}
		common.PrintError(common.StructToString(notificationResult), "Notification result")

		if err := common.Database.Debug().Create(&notificationResult).Error; err != nil {
			common.PrintError("Error inserting notification result", err.Error())
			return fmt.Errorf("failed to insert notification result: %v", err)
		}
	}

	return nil
}
