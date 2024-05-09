package utils

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/FourWD/middleware/common"
	"github.com/jung-kurt/gofpdf"
)

// func GetPDFVehicle(c *fiber.Ctx) (string, error) {

// 	auctionID := c.Params("auction_id")
// 	query := c.Queries()
// 	vehicleIDList := query["vehicle_id"]
// 	vehicleArr := strings.Split(vehicleIDList, ",")

// 	GenPDFVehicle(auctionID, vehicleArr)

// 	path, err := common.UploadPdfToGoogle(pdf, "รายการประมูลรถยนต์", "auction", "fourwd-auction") //carlist
// 	if err != nil {
// 		return "", err
// 	}
// 	return path, nil
// }

// ==================================

func headerCarlist(pdf gofpdf.Pdf) {
	// Table header
	// pdf.MoveTo(15, 35)
	pdf.SetFont("THSarabun", "B", 12)
	pdf.CellFormat(5, 8, "#", "1", 0, "C", false, 0, "")
	pdf.CellFormat(80, 8, "รายละเอียดรถยนต์", "1", 0, "C", false, 0, "")
	pdf.CellFormat(15, 8, "เกรด", "1", 0, "C", false, 0, "")
	pdf.CellFormat(15, 8, "เลขไมล์", "1", 0, "C", false, 0, "")
	pdf.CellFormat(20, 8, "ลานจอด", "1", 0, "C", false, 0, "")
	pdf.CellFormat(20, 8, "ปีผลิต", "1", 0, "C", false, 0, "")
	pdf.CellFormat(20, 8, "วันหมดภาษี", "1", 0, "C", false, 0, "")
	pdf.CellFormat(25, 8, "ราคาตั้งประมูล", "1", 0, "C", false, 0, "")
	pdf.Ln(-1)

}

func GenPDFVehicle(auctionID string, vehicles []string) (string, error) { //carlist

	auctionDetail, err := prepareAuctionAndVehicleDetails(auctionID, vehicles)
	if err != nil {
		return "", err
	}

	filepathStr := "images/pdf/"
	// fileextention := ".pdf"
	// Create new PDF instance
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetMargins(5, 7, 5)
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
	pdf.Text(48, 15, "รายการประมูลรถยนต์ "+auctionDetail.Auction[0].RoundName)
	pdf.Ln(10)

	pdf.SetFont("THSarabun", "", 16)
	pdf.Text(10, 33, "รอบการประมูลของ "+auctionDetail.Auction[0].Name+" | "+auctionDetail.Auction[0].AuctionName)
	pdf.Ln(10)

	pdf.SetLineWidth(0.2) // Set line width
	pdf.Line(10, 25, 200, 25)
	pdf.Ln(10)
	pdf.SetLineWidth(0) // Set line width

	pdf.SetFont("THSarabun", "B", 18)
	pdf.Text(180, 33, "จำนวน "+fmt.Sprintf("%d", (len(auctionDetail.Vehicles)))+" คัน")

	// Table header
	// pdf.MoveTo(15, 35)
	headerCarlist(pdf)

	page := 1
	perpage := 8
	line := 1
	tableYz := 53
	//loop vehicle

	for i, vehicle := range auctionDetail.Vehicles {

		i++

		if page == 1 {
		} else {
			perpage = 9
		}
		fmt.Printf("index = %d, License = %s", i, vehicle.License)
		fmt.Printf(", page = %d, perpage = %d", page, perpage)
		fmt.Printf(", line = %d", line)
		fmt.Println("========================")

		// Set the fill color to light gray
		pdf.SetFillColor(240, 240, 240) // Light gray color for the cell background

		pdf.SetFont("THSarabun", "B", 9)
		pdf.CellFormat(5, 8, vehicle.VehicleNo, "1", 0, "C", true, 0, "")
		pdf.CellFormat(80, 8, vehicle.License+" "+vehicle.LicenseProvinceName+" | "+

			vehicle.Year+" "+vehicle.VehicleBrandName+" "+vehicle.VehicleModelName+" "+
			vehicle.VehicleSubModelName+" "+vehicle.VehicleColorName, "1", 0, "", true, 0, "")

		pdf.CellFormat(15, 8, vehicle.VehicleGradeID, "1", 0, "C", true, 0, "")

		mile, _ := strconv.ParseFloat(vehicle.Mile, 64)
		milecomma := common.FloatWithCommas(mile, 0)
		pdf.CellFormat(15, 8, milecomma, "1", 0, "C", true, 0, "")
		pdf.CellFormat(20, 8, vehicle.Label, "1", 0, "C", true, 0, "")
		DDMMYYYY := "02/01/2006"

		pdf.CellFormat(20, 8, vehicle.Year, "1", 0, "C", true, 0, "")
		// pdf.CellFormat(20, 8, vehicle.TaxExpireDate.Format(DDMMYYYY), "1", 0, "C", true, 0, "")

		var taxExpireDate string
		if !vehicle.TaxExpireDate.IsZero() && vehicle.TaxExpireDate.After(time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC)) {
			taxExpireDate = vehicle.TaxExpireDate.Format(DDMMYYYY)
		} else {
			taxExpireDate = ""
		}
		pdf.CellFormat(20, 8, taxExpireDate, "1", 0, "C", true, 0, "")

		// h := 50
		// numFloat := float64(h)
		price, _ := strconv.ParseFloat(vehicle.OpenPrice, 64)

		var openprice string
		if price == 0 {
			openprice = "รอกำหนดราคา"
		} else {
			openprice = common.FloatWithCommas(price, 2)
		}

		pdf.SetFont("THSarabun", "", 9)
		pdf.SetTextColor(0, 0, 0)

		pdf.SetFont("THSarabun", "B", 9)
		pdf.SetTextColor(255, 0, 0)
		pdf.CellFormat(25, 8, openprice, "1", 0, "C", true, 0, "")
		pdf.SetTextColor(0, 0, 0)
		pdf.SetFont("THSarabun", "", 9)

		pdf.Ln(-1)
		tableYz += 2

		if vehicle.GradeRemark != "" {

			pdf.SetFont("THSarabun", "", 9)
			// pdf.CellFormat(5, x, "", "1", 0, "C", false, 0, "")
			pdf.MultiCell(0, 5.25, vehicle.GradeRemark, "1", "L", false) // MultiCell for remark column
			pdf.SetFillColor(240, 240, 240)

		}

		line++
		newpage := false

		if i == 8 {
			newpage = true

		} else if i > 8 {
			temp := i - 8
			if temp%perpage == 0 {
				newpage = true
			}
		}
		if newpage {
			line = 1
			page++
			pdf.AddPage()
			headerCarlist(pdf)
			pdf.SetMargins(5, 5, 5)

			tableYz = 2

		}

	}
	pdf.SetFont("THSarabun", "", 10)
	pdf.SetTextColor(255, 0, 0)
	pdf.Cell(10, 10, "หมายเหตุ")
	pdf.SetTextColor(0, 0, 0)

	pdf.Cell(10, 10, "*รถที่ประมูลได้ราคารถยังไม่รวมภาษีมูลค่าเพิ่ม 7%")
	pdf.Ln(5)

	xx := pdf.GetX()
	yy := pdf.GetY()
	xx += 10
	pdf.SetXY(xx, yy)
	pdf.CellFormat(10, 10, "**บริษัทฯ ขอสงวนสิทธิ์ในการพิจารณายกเลิก หรือเปลี่ยนแปลงการประมูลขายรถบางรายการได้ตามความเหมาะสม โดยไม่ต้องแจ้งให้ทราบล่วงหน้า", "", 0, "L", false, 0, "")
	pdf.Ln(-1)

	fileName := generateFileNameVehicle(auctionDetail.Auction[0].Name, auctionDetail.Auction[0].AuctionName)
	path, err := common.UploadPdfToGoogle(pdf, fileName, "auction", "fourwd-auction")
	if err != nil {
		return "", err
	}
	// sanitizedPath := sanitizeFilename(path)

	return path, nil

}

func generateFileNameVehicle(auctionName, auctionDate string) string {
	auctionName = strings.ReplaceAll(auctionName, " ", "_")
	auctionDate = strings.ReplaceAll(auctionDate, "/", "")
	auctionDate = strings.ReplaceAll(auctionDate, ":", "")
	fileName := "รายการประมูลรถยนต์_" + auctionName + "_" + auctionDate
	return fileName
}

func prepareAuctionAndVehicleDetails(auctionID string, vehicleID []string) (AuctionVehicleDetail, error) {
	var auctionNames []AuctionList
	common.Database.Raw(`
							SELECT r.name as name,a.name as round_name, 
							CONCAT_WS(" ", a.name, CONCAT(TIME_FORMAT(Time(a.start_date), "%H.%i"), " น.", " - ", TIME_FORMAT(Time(a.end_date), "%H.%i"), " น.")) AS auction_name
									FROM auctions a
								LEFT JOIN rounds r on a.round_id = r.id WHERE a.id = ?  `, auctionID).Scan(&auctionNames)
	common.Print("", common.StructToString(auctionNames))

	// if len(auctionNames) < 2 {
	// 	return AuctionVehicleDetail{}, errors.New("insufficient auction names data")
	// }
	var vehicles []VehicleCarlist
	str := strings.Join(vehicleID, "','")
	str = fmt.Sprintf("'%s'", str)
	common.Print("", str)
	common.Database.Raw(`
						SELECT  av.auction_id,av.vehicle_no,av.vehicle_id,vgr.grade_remark,b.label, v.*, v.year_manufacturing as year, v.tax_expire_date,v.license_expire_date,av.open_price
						FROM vehicles v
	  						 LEFT JOIN auction_vehicles av ON v.id = av.vehicle_id
						LEFT JOIN auctions a on av.auction_id = a.id  AND av.auction_id = ?
						LEFT JOIN branches b ON v.branch_id = b.id
						LEFT JOIN vehicle_grade_remarks vgr on v.chassis_no = vgr.chassis_number
	  						 WHERE av.vehicle_id IN (?) and av.auction_id = ? order by vehicle_no asc
    `, auctionID, vehicleID, auctionID).Debug().Scan(&vehicles)
	fmt.Print("lennn", len(vehicles))
	common.Print("", common.StructToString(vehicles))

	auctionVehicleDetail := AuctionVehicleDetail{
		Auction: auctionNames,
		// AuctionName: auctionNames[1],
		Vehicles: vehicles,
	}
	common.Print("", common.StructToString(auctionVehicleDetail))

	return auctionVehicleDetail, nil
}
