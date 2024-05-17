package utils

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/FourWD/middleware/common"
)

func GenZipVehicleImage(auctionID string, userID string) (string, error) {

	mainPath := common.DateString() + "_" + common.MD5(auctionID + userID)[:4]
	vehicles := prepareData(auctionID, userID)
	if len(vehicles) < 1 {
		return "no data", errors.New("no data")
	}
	// println(mainPath)

	_, errs := os.MkdirTemp("/tmp", mainPath)
	if errs != nil {
		common.Print("error make main dir", errs.Error())
		return "", errs
	}
	var file_name = ""
	for _, v := range vehicles {
		vehiclePath := fmt.Sprintf("/tmp/" + mainPath + "/" + v.License)
		errss := os.MkdirAll(vehiclePath, 0777)
		if errss != nil {
			common.PrintError("error make sub dir", errss.Error())
			return "", errss
		}
		file_name = v.License + "_" + strings.ReplaceAll(v.ImageName, "/", "") + ".jpg"
		err := common.DownloadFile(v.ImagePath, vehiclePath+"/"+file_name, "auction", "fourwd-auction")
		if err != nil {
			// println("download " + file_name + " failed") // COMMENT WORKING PATH
			return "", err
		}

		// println("download " + file_name + " work") // COMMENT WORKING PATH
		// download v.image path to this folder rename to imagename
	}
	// zip folder mainpath
	zipPath, errs := common.ZipFile(mainPath, mainPath)
	if errs != nil {
		return "", errs
	}
	// call fn upload file to google drive
	resultPath, errss := common.UploadFileToGoogle(zipPath, "auction", "fourwd-auction")
	if errss != nil {
		return "", errss
	}
	// common.Print("zip path", resultPath)
	return resultPath, nil
}

type vehicleImage struct {
	License   string `json:"license"`
	ImageName string `json:"image_name"`
	ImagePath string `json:"image_path"`
}

func prepareData(auctionID string, userID string) []vehicleImage {
	var vehicleList []vehicleImage

	sql := `SELECT v.license , vm.name as image_name , vem.image_path 
	FROM template_download_zip_vehicle_images dm
	LEFT JOIN template_vehicle_images vm ON dm.vehicle_image_id = vm.id
	LEFT JOIN vehicle_images vem ON vm.id = vem.template_vehicle_image_id
	LEFT JOIN vehicles v ON vem.vehicle_id = v.id
	WHERE vem.vehicle_id IN (SELECT v.id FROM vehicles v
	LEFT JOIN auction_vehicles av ON v.id = av.vehicle_id 
	WHERE av.auction_id = ? AND av.winner_user_id = ? AND av.is_win = 1) AND vem.is_delete = 0 
	ORDER BY vehicle_id , dm.row_order`

	common.Database.Raw(sql, auctionID, userID).Scan(&vehicleList)

	return vehicleList
}
