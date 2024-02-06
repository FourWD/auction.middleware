package utils

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/FourWD/middleware/common"
	"github.com/jung-kurt/gofpdf"
)

type AuctionVehicleDetail struct {
	AuctionName string
	Name        string
	Vehicles    []VehicleCarlist
}
type VehicleCarlist struct {
	// Vehicle              []orm.Vehicle
	VehicleAuctionRecipt string `json:"vehicle_auction_receipt"`
	VehicleGradeID       string `json:"vehicle_grade_id"`
	TotalCar             string `json:"total_car"`
	VehicleBrandName     string `json:"vehicle_brand_name"`
	VehicleSubModelName  string `json:"vehicle_model_name"`
	VehicleModelName     string `json:"vehicle_sub_model_name"`
	VehicleColorName     string `json:"vehicle_color_name"`
	YearManuFacturing    string `json:"year_manufacturing"`
	ImagePreviewPath     string `json:"image_preview_path"`
	OpenPrice            string `json:"open_price"`
	TaxExprieDate        string `json:"tax_expire_date"`
	Mile                 string `json:"mile"`
	Remark               string `json:"remark"`
	License              string `json:"license"`
	Label                string `json:"label"`
	ChassisNo            string `json:"chassis_no"`

	LicenseProvinceName string `json:"license_province_name"`
}

func GenPDFVehicle(auctionID string, userID string, vehicles []string) (string, error) {
	auctionDetail, err := prepareAuctionAndVehicleDetails(auctionID, vehicles)
	if err != nil {
		return "", err
	}

	filepathStr := "images/pdf/"
	// fileextention := ".pdf"
	// Create new PDF instance
	pdf := gofpdf.New("P", "mm", "A4", "")

	// Add page to PDF
	pdf.AddPage()
	pdf.AddUTF8Font("THSarabun", "", "fonts/THSarabun.ttf")
	pdf.AddUTF8Font("THSarabun", "B", "fonts/THSarabunBold.ttf")
	x := 8.00
	y := 4.00
	width := 30.00
	height := 18.00

	imagePath := filepathStr + "simple.png"
	pdf.Image(imagePath, x, y, width, height, false, "", 0, "")
	// Add auction details to PDF
	pdf.SetFont("THSarabun", "B", 24)
	pdf.Text(48, 15, "รายการประมูลรถยนต์ "+auctionDetail.AuctionName)
	pdf.Ln(10)

	pdf.SetFont("THSarabun", "", 16)
	pdf.Text(10, 33, "รอบการประมูลของ บริษัท ตรีเพชรอีซูซุลิสซิ่ง จำกัด | "+auctionDetail.AuctionName)
	pdf.Ln(10)

	pdf.SetLineWidth(0.2) // Set line width
	pdf.Line(10, 25, 200, 25)
	pdf.Ln(10)
	pdf.SetLineWidth(0) // Set line width

	pdf.SetFont("THSarabun", "B", 18)
	pdf.Text(180, 33, "จำนวน "+fmt.Sprintf("%d", (len(auctionDetail.Vehicles)))+" คัน")
	// Table header
	// pdf.MoveTo(15, 35)
	pdf.SetFont("THSarabun", "B", 12)
	pdf.CellFormat(5, 8, "#", "1", 0, "C", false, 0, "")
	pdf.CellFormat(80, 8, "รายละเอียดรถยนต์", "1", 0, "C", false, 0, "")
	pdf.CellFormat(15, 8, "เกรด", "1", 0, "C", false, 0, "")
	pdf.CellFormat(20, 8, "เลขไมล์", "1", 0, "C", false, 0, "")
	pdf.CellFormat(25, 8, "ลานจอด", "1", 0, "C", false, 0, "")
	pdf.CellFormat(20, 8, "วันหมดภาษี", "1", 0, "C", false, 0, "")
	pdf.CellFormat(25, 8, "ราคาตั้งประมูล", "1", 0, "C", false, 0, "")
	pdf.Ln(-1)

	//loop vehicle
	pdf.SetFont("THSarabun", "", 12)
	counter := 1
	for _, vehicle := range auctionDetail.Vehicles {
		// Set the fill color to light gray
		pdf.SetFillColor(240, 240, 240) // Light gray color for the cell background

		pdf.CellFormat(5, 8, strconv.Itoa(counter), "1", 0, "C", true, 0, "")

		pdf.CellFormat(80, 8, vehicle.License+" "+vehicle.LicenseProvinceName+" | "+
			vehicle.YearManuFacturing+" "+vehicle.VehicleBrandName+" "+vehicle.VehicleModelName+" "+
			vehicle.VehicleSubModelName+" "+vehicle.VehicleColorName, "1", 0, "", true, 0, "")

		pdf.CellFormat(15, 8, vehicle.VehicleGradeID, "1", 0, "C", true, 0, "")
		pdf.CellFormat(20, 8, vehicle.Mile, "1", 0, "C", true, 0, "")
		pdf.CellFormat(25, 8, vehicle.Label, "1", 0, "C", true, 0, "")
		pdf.CellFormat(20, 8, vehicle.TaxExprieDate, "1", 0, "C", true, 0, "")
		pdf.SetFont("THSarabun", "B", 12)
		pdf.SetTextColor(255, 0, 0)
		pdf.CellFormat(25, 8, vehicle.OpenPrice, "1", 0, "C", true, 0, "")
		pdf.SetTextColor(0, 0, 0)
		pdf.Ln(-1)
		pdf.CellFormat(5, 8, "", "1", 0, "C", false, 0, "")
		pdf.MultiCell(0, 6, vehicle.Remark, "1", "L", false) // MultiCell for remark column

		pdf.SetFillColor(240, 240, 240)

		counter++

	}
	pdf.Ln(2)

	pdf.SetFont("THSarabun", "", 10)

	pdf.SetTextColor(255, 0, 0)
	pdf.Cell(10, 10, "หมายเหตุ")
	pdf.SetTextColor(0, 0, 0)

	pdf.Cell(10, 10, "*รถที่ประมูลได้ราคารถยังไม่รวมภาษีมูลค่าเพิ่ม 7%")
	pdf.Ln(5)

	pdf.Cell(10, 10, "**บริษัทฯ ขอสงวนสิทธิ์ในการพิจารณายกเลิก หรือเปลี่ยนแปลงการประมูลขายรถบางรายการได้ตามความเหมาะสม โดยไม่ต้องแจ้งให้ทราบล่วงหน้า")
	pdf.Ln(10)

	// filedestination := filepathStr + "รายการประมูลรถยนต์" + fileextention

	// Output PDF to file
	// err = pdf.OutputFileAndClose(filedestination)
	// if err != nil {
	// 	return "", err
	// }]

	path, err := common.UploadPdfToGoogle(pdf, "pdf", "auction", "fourwd-auction")
	if err != nil {
		return "", err
	}
	return path, nil
}

func prepareAuctionAndVehicleDetails(auctionID string, vehicleID []string) (AuctionVehicleDetail, error) {
	var auctionNames []string
	common.Database.Raw(`
        SELECT 
            CONCAT_WS(" ", a.name, CONCAT(TIME_FORMAT(Time(a.start_date), "%H.%i"), " น.", " - ", TIME_FORMAT(Time(a.end_date), "%H.%i"), " น.")) AS auction_name
        FROM 
            auctions a
        WHERE 
            a.id = ?
    `, auctionID).Scan(&auctionNames)

	var vehicles []VehicleCarlist
	str := strings.Join(vehicleID, "','")
	str = fmt.Sprintf("'%s'", str)
	common.Print("", str)
	common.Database.Raw(`
	SELECT b.label, v.*, v.tax_expire_date,av.open_price
			FROM vehicles v
       		LEFT JOIN auction_vehicles av ON v.id = av.vehicle_id
			LEFT JOIN branches b ON v.branch_id = b.id
        WHERE av.vehicle_id IN (?)
    `, vehicleID).Debug().Scan(&vehicles)
	fmt.Print("lennn", len(vehicles))
	common.Print("", common.StructToString(vehicles))

	auctionVehicleDetail := AuctionVehicleDetail{
		AuctionName: auctionNames[0],
		Vehicles:    vehicles,
	}
	common.Print("", common.StructToString(auctionVehicleDetail))

	return auctionVehicleDetail, nil
}

// func prepareAuctionDetail(auctionID string) []string {
// 	var auctionNames []string
// 	common.Database.Raw(`
// 		SELECT
// 			CONCAT_WS(" ", a.name, CONCAT(TIME_FORMAT(Time(a.start_date), "%H.%i"), " น.", " - ", TIME_FORMAT(Time(a.end_date), "%H.%i"), " น.")) AS auction_name
// 		FROM
// 			auctions a
// 		WHERE
// 			a.id = ?
// 	`, auctionID).Scan(&auctionNames)
// 	return auctionNames
// }

// func prepareVehicleList(vehicleID []string) []VehicleCarlist {

// 	var vehicles []VehicleCarlist

// 	str := strings.Join(vehicleID, "','")
// 	str = fmt.Sprintf("'%s'", str)

// 	common.Print("", str)

// 	common.Database.Raw(`SELECT * FROM vehicles v
// 	left join auction_vehicles av on v.id = av.vehicle_id
// 	where av.auction_id IN (?)
// 	`, vehicleID).Scan(&vehicles)

// 	return vehicles
// }
