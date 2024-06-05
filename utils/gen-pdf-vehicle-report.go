package utils

import (
	"strconv"
	"time"

	"github.com/FourWD/middleware/common"
	"github.com/jung-kurt/gofpdf"
)

// func GenPDFVehicleReport(auctionID string, vehicleID string) (string, error) { //ใบรายงานผลการประมูล สรรพากร
// 	pdf := gofpdf.New("P", "mm", "A4", "")

// 	GenPDFReport(pdf, auctionID, vehicleID)

// 	path, err := common.UploadPdfToGoogle(pdf, "ใบประกาศผลประมูล", "auction", "fourwd-auction")
// 	if err != nil {
// 		return "", err
// 	}
// 	return path, nil
// }

//	type SummaryReport struct {
//		StartDateAuction string `json:"start_date_auction"`
//			SourceName  string `json:"source_name"`
//			AuctionName string `json:"auction_name"`
//			TotalCar    int    `json:"total_car"`
//			TotalPrice  int    `json:"total_price"`
//		}
type VehicleReport struct {
	VehicleID       string `json:"vehicle_id"`
	VehicleNo       string `json:"vehicle_no"`
	Receipt         string `json:"receipt"`
	VehicleGradeID  string `json:"vehicle_grade_id"`
	VehicleGearName string `json:"vehicle_gear_name"`
	CN              string `json:"cn"`

	VehicleBrandName      string `json:"vehicle_brand_name"`
	VehicleBrandNameTh    string `json:"vehicle_brand_name_th"`
	VehicleSubModelNameTh string `json:"vehicle_sub_model_name_th"`
	VehicleSubModelName   string `json:"vehicle_sub_model_name"`
	VehicleModelName      string `json:"vehicle_model_name"`
	VehicleModelNameTh    string `json:"vehicle_model_name_th"`
	VehicleColorName      string `json:"vehicle_color_name"`
	Years                 string `json:"years"`
	YearRegister          string `json:"year_register"`
	ImagePreviewPath      string `json:"image_preview_path"`
	OpenPrice             string `json:"open_price"`
	ClosePrice            string `json:"close_price"`
	TotalCar              string `json:"total_car"`

	SKU                 string `json:"sku"`
	EngineSize          string `json:"engine_size"`
	EngineNo            string `json:"engine_no"`
	EngineCapacity      string `json:"engine_capacity"`
	VehicleFuelTypeName string `json:"vehicle_fuel_type_name"`
	Mile                string `json:"mile"`
	ChassisNo           string `json:"chassis_no"`
	VehicleTypeName     string `json:"vehicle_type_name"`
	VehicleSubTypeName  string `json:"vehicle_sub_type_name"`
	SourceName          string `json:"source_name"`
	SourceID            string `json:"source_id"`
	Branch              string `json:"branch"`
	BranchLabel         string `json:"branch_label"`
	License             string `json:"license"`
	LicenseProvinceName string `json:"license_province_name"`

	LicenseReceiveDate string    `json:"license_receive_date"`
	StartDateAuction   time.Time `json:"start_date_auction"`

	EndDateAuction time.Time `json:"end_date_auction"`
}

// func headertableReport(pdf gofpdf.Pdf, tabley int) { //มีราคาเปิด
// Define table headers
// pdf.SetFillColor(196, 16, 22)
// pdf.SetTextColor(255, 255, 255) // ตั้งค่าสีข้อความเป็นสีขาว
// pdf.SetDrawColor(0, 0, 0)
// pdf.SetLineWidth(0.2)
// pdf.SetFont("Sarabun", "B", 16)
// tableX := 0
// tableY := tabley
// pdf.SetXY(float64(tableX), float64(tableY))
// cellWidth := 20.0
// cellHeight := 10.0
// // border := "10"
// pdf.CellFormat(10, 18, "เลขรถ", "1", 0, "C", true, 0, "")
// pdf.CellFormat(15, 18, "ยี่ห้อรถ", "1", 0, "C", true, 0, "")
// // pdf.CellFormat(10, 8, "ช่อง", "1", 0, "C", true, 0, "")
// pdf.CellFormat(60, 18, "รุ่นรถ", "1", 0, "C", true, 0, "")
// pdf.CellFormat(15, 18, "ปีผลิต", "1", 0, "C", true, 0, "")
// pdf.CellFormat(10, 18, "เกรด", "1", 0, "C", true, 0, "")
// pdf.CellFormat(15, 18, "สี", "1", 0, "C", true, 0, "")
// pdf.CellFormat(25, 18, "ทะเบียนรถ", "1", 0, "C", true, 0, "")
// pdf.CellFormat(20, 18, "ราคาเปิด", "1", 0, "C", true, 0, "")
// // กำหนดพื้นหลังสีเหมือนกันสำหรับเซลล์ทั้งหมด
// pdf.SetFillColor(196, 16, 22)
// pdf.SetDrawColor(0, 0, 0)
// // แสดงข้อความในเซลล์ MultiCell และเซลล์อื่น ๆ
// pdf.SetFont("Sarabun", "B", 12)
// pdf.CellFormat(40, 8, "ราคาประมูลสูงสุด", "1", 0, "C", true, 0, "")
// pdf.SetFont("Sarabun", "B", 10)
// pdf.SetXY(170, 78)
// pdf.MultiCell(cellWidth, cellHeight/2, "ไม่รวมภาษี\nมูลค่าเพิ่ม", "1", "C", true)
// pdf.SetXY(190, 78)
// pdf.MultiCell(cellWidth, cellHeight/2, "รวมภาษี\nมูลค่าเพิ่ม", "1", "C", true)
// // pdf.CellFormat(20, 8, "รวมภาษีมูลค่าเพิ่ม", "1", 0, "C", true, 0, "")
// // กำหนดสีข้อความและพื้นหลังเป็นค่าเริ่มต้น
// pdf.SetTextColor(0, 0, 0)
// pdf.SetFillColor(255, 255, 255)
// }

func headertableReport(pdf gofpdf.Pdf, tabley int, y int) {
	// Define table headers
	pdf.SetFillColor(196, 16, 22)
	pdf.SetTextColor(255, 255, 255) // ตั้งค่าสีข้อความเป็นสีขาว
	pdf.SetDrawColor(0, 0, 0)
	pdf.SetLineWidth(0.2)
	pdf.SetFont("Sarabun", "B", 16)
	tableX := 0
	tableY := tabley
	pdf.SetXY(float64(tableX), float64(tableY))
	cellWidth := 20.0
	cellHeight := 10.0
	// border := "10"
	pdf.CellFormat(10, 18, "เลขรถ", "1", 0, "C", true, 0, "")
	pdf.CellFormat(15, 18, "ยี่ห้อรถ", "1", 0, "C", true, 0, "")
	// pdf.CellFormat(10, 8, "ช่อง", "1", 0, "C", true, 0, "")
	pdf.CellFormat(55, 18, "รุ่นรถ", "1", 0, "C", true, 0, "")
	pdf.CellFormat(15, 18, "ปีผลิต", "1", 0, "C", true, 0, "")
	pdf.CellFormat(10, 18, "เกรด", "1", 0, "C", true, 0, "")
	pdf.CellFormat(15, 18, "สี", "1", 0, "C", true, 0, "")
	pdf.CellFormat(30, 18, "ทะเบียนรถ", "1", 0, "C", true, 0, "")
	pdf.CellFormat(20, 18, "ราคาเปิด", "1", 0, "C", true, 0, "")

	// กำหนดพื้นหลังสีเหมือนกันสำหรับเซลล์ทั้งหมด
	pdf.SetFillColor(196, 16, 22)
	pdf.SetDrawColor(0, 0, 0)

	// แสดงข้อความในเซลล์ MultiCell และเซลล์อื่น ๆ
	pdf.SetFont("Sarabun", "B", 12)

	pdf.CellFormat(40, 8, "ราคาประมูลสูงสุด", "1", 0, "C", true, 0, "")
	pdf.SetFont("Sarabun", "B", 10)
	pdf.SetXY(170, float64(y))
	pdf.MultiCell(cellWidth, cellHeight/2, "ไม่รวมภาษี\nมูลค่าเพิ่ม", "1", "C", true)
	pdf.SetXY(190, float64(y))
	pdf.MultiCell(cellWidth, cellHeight/2, "รวมภาษี\nมูลค่าเพิ่ม", "1", "C", true)
	// pdf.CellFormat(20, 8, "รวมภาษีมูลค่าเพิ่ม", "1", 0, "C", true, 0, "")

	// กำหนดสีข้อความและพื้นหลังเป็นค่าเริ่มต้น
	pdf.SetTextColor(0, 0, 0)
	pdf.SetFillColor(255, 255, 255)
}

func GenPDFReport(auctionID string, vehicleID string) (string, error) { //ใบรายงานผลการประมูล สรรพากร

	vehicles, _ := prepareVehicleResult(auctionID, vehicleID)

	filepathStr := "images/pdf/"
	// fileextention := ".pdf"
	pdf := gofpdf.New("P", "mm", "A4", "")

	for _, v := range vehicles {

		pdf.AddPage()
		pdf.AddUTF8Font("Sarabun", "", "fonts/THSarabun.ttf")
		pdf.AddUTF8Font("Sarabun", "B", "fonts/THSarabunBold.ttf")

		x := 8.00
		y := 10.00
		width := 32.00
		height := 18.00

		imagePath := filepathStr + "simple.png"
		pdf.Image(imagePath, x, y, width, height, false, "", 0, "")

		//header
		pdf.SetTextColor(0, 0, 0)
		pdf.SetFont("Sarabun", "B", 10)
		pdf.Text(45, 10, "บริษัท.โอมาคาเสะคาร์ จำกัด")
		pdf.Text(45, 15, "00002 เลขที่11/11 หมู่ที่8 ถนนประเสริฐมนูกิจ แขวงนวลจันทร์เขตบึงกุ่ม กรุงเทพมหานคร 10240 โทร 098-828-1111")
		pdf.Text(45, 20, "00003 เลขที่90 ถนนงามวงศ์วาน แขวงลาดยาว เขตจตุจักร กรุงเทพมหานคร 10900 โทร 02-102-6481")
		pdf.Text(45, 25, "00004 เลขที่898 หมู่ที่1 ตำบลกลางดง อำเภอปากช่อง จังหวัดนครราชสมีา 30320 โทร 044-002-555")
		pdf.Text(45, 30, "00005 เลขที่22 หมู่ที่18 ตำบลบางแม่นาง อำเภอบางใหญ่จังหวัดนนทบุรี 11140 โทร 02-103-1920")
		pdf.Text(45, 35, "00006 เลขที่99/2 หมู่ที่1 ตำบลธงชัยเหนือ อำเภอปักธงชัย จังหวัดนครราชสีมา 30150 โทร 044-022-772")
		pdf.Text(45, 40, "00007 เลขที่99 หมู่ที่1 ตำบลหนองชุมพลเหนือ อำเภอเขาย้อย จังหวัดเพชรบุรี 76140 โทร 032-890-151")
		pdf.Text(45, 45, "เลขประจำตัวผู้เสียภาษี 0105562094168")

		pdf.SetTextColor(0, 0, 0)
		pdf.SetFont("Sarabun", "B", 20)
		pdf.Text(80, 55, "ใบรายงานผลการประมูล")

		//text detail+22
		pdf.SetFont("Sarabun", "B", 13)
		pdf.SetXY(100, 60)
		pdf.SetFillColor(255, 255, 255)
		pdf.SetDrawColor(0, 0, 0)
		// pdf.Text(120, 80, "ระบุสถานที่ประมูล ")
		pdf.CellFormat(35, 8, "ระบุสถานที่ประมูล", "0", 0, "L", true, 0, "")
		pdf.CellFormat(50, 8, "OMAKASECAR ON-LINE AUCTION", "0", 0, "L", true, 0, "")

		pdf.SetFont("Sarabun", "B", 13)
		pdf.SetXY(100, 68)
		pdf.SetFillColor(255, 255, 255)
		pdf.SetDrawColor(0, 0, 0)
		// pdf.Text(120, 80, "ระบุสถานที่ประมูล ")

		pdf.CellFormat(35, 8, "วันที่", "0", 0, "L", true, 0, "")
		// formattedEndDate := v.EndDateAuction.Format("02 มกราคม 2006")
		currentDate := v.EndDateAuction
		formattedEndDate := strconv.Itoa(currentDate.Day()) + " " + getThaiMonth(currentDate.Month().String()) + " " + strconv.Itoa(currentDate.Year()+543)
		pdf.CellFormat(50, 8, formattedEndDate, "0", 0, "L", true, 0, "")

		pdf.SetFont("Sarabun", "B", 13)
		pdf.SetXY(100, 76)
		pdf.SetFillColor(255, 255, 255)
		pdf.SetDrawColor(0, 0, 0)
		// pdf.Text(120, 80, "ระบุสถานที่ประมูล ")
		pdf.CellFormat(35, 8, "สถานที่ออก ", "0", 0, "L", true, 0, "")
		pdf.CellFormat(50, 8, "สำนักงานใหญ่", "0", 0, "L", true, 0, "")

		pdf.SetFont("Sarabun", "B", 15)
		pdf.SetXY(10, 86)
		pdf.SetFillColor(255, 255, 255)
		pdf.SetDrawColor(0, 0, 0)
		// pdf.Text(120, 80, "ระบุสถานที่ประมูล ")
		pdf.MultiCell(190, 12/2, "         ตามที่ "+v.SourceName+" ได้มอบหมายให้บริษัท โอมาคาเสะ คาร์ จำกัด\nทำการขายทอดตลาดรถยนต์และอุปกรณ์เมื่อวันที่ "+formattedEndDate+" ผลการประมูลทอดตลาดปรากฏ ดังนี้", "0", "L", true)

		// pdf.CellFormat(100, 10, "ตามที่"+v.SourceName+" ได้มอบหมายให้บริษัท โอมาคาเสะ คาร์ จำกัด", "0", 0, "L", true, 0, "")
		// pdf.CellFormat(100, 10, "ทำการขายทอดตลาดรถยนต์และอุปกรณ์เมื่อวันที่ "+formattedEndDate+" ผลการประมูลทอดตลาดปรากฏ ดังนี้", "0", 0, "L", true, 0, "")

		y += 5.00
		pdf.Ln(5)

		// Table
		pdf.SetXY(10, 100)
		pdf.SetFillColor(255, 255, 255)
		pdf.SetTextColor(0, 0, 0)
		pdf.SetFont("Sarabun", "B", 15)
		pdf.CellFormat(15, 8, "ลำดับที่", "1", 0, "C", true, 0, "")
		pdf.CellFormat(35, 8, "เลขที่ใบรับรถ", "1", 0, "C", true, 0, "")
		pdf.CellFormat(65, 8, "รายละเอียด", "1", 0, "C", true, 0, "")
		pdf.CellFormat(25, 8, "ผลประมูล", "1", 0, "C", true, 0, "")
		pdf.CellFormat(35, 8, "ราคาที่ประมูล", "1", 0, "C", true, 0, "")
		pdf.CellFormat(15, 8, "ค่าอื่นๆ", "1", 0, "C", true, 0, "")

		pdf.Ln(-1)
		//detail
		pdf.SetFont("Sarabun", "B", 15)

		pdf.Text(16, 113, v.VehicleNo)
		pdf.CellFormat(15, 60, "", "1", 0, "TC", false, 0, "")
		pdf.Text(29, 113, v.Receipt)
		pdf.CellFormat(35, 60, "", "1", 0, "TC", false, 0, "")

		var test1 string
		if v.License != "" && v.LicenseProvinceName != "" {
			test1 = "ทะเบียน : " + v.License + " " + v.LicenseProvinceName
		} else {
			test1 = "ทะเบียน : " + v.ChassisNo
		}
		// test1 := "ทะเบียน : " + v.License + " " + v.LicenseProvinceName
		test2 := "ยี่ห้อ : " + v.VehicleBrandName + " รุ่น : " + v.VehicleModelName
		// test3 := "รุ่น : " + v.VehicleModelName
		test4 := "รุ่นย่อย : " + v.VehicleSubModelName
		// test5 := "สี : " + v.VehicleColorName
		test6 := "ปี : " + v.YearRegister + " เกียร์ : " + v.VehicleGearName
		// test7 := "เกียร์ : " + v.VehicleGearName
		test8 := "เลขเครื่อง : " + v.EngineNo
		test9 := "เลขตัวถัง : " + v.ChassisNo
		test10 := "เลขที่สัญญา  : " + v.CN

		pdf.MultiCell(70, 7, test1+"\n"+test2+"\n"+test4+"\n"+test6+"\n"+test8+"\n"+test9+"\n"+test10+"\n", "0", "L", false)

		closePriceFloat, _ := strconv.ParseFloat(v.ClosePrice, 64)
		closePriceComma := common.FloatWithCommas(closePriceFloat, 2)

		pdf.SetXY(125, 108) // ตั้งตำแหน่ง X และ Y ใหม่
		pdf.Text(132, 113, "ขายได้")
		pdf.CellFormat(25, 60, "", "1", 0, "TC", false, 0, "")
		// pdf.Text(154, 113, closePriceComma)
		pdf.CellFormat(35, 60, closePriceComma, "1", 0, "TC", false, 0, "")
		pdf.CellFormat(15, 60, "", "1", 0, "TC", false, 0, "")

		pdf.Ln(-1)

		pdf.SetFont("Sarabun", "B", 15)
		pdf.CellFormat(15, 8, "", "1", 0, "", false, 0, "")
		pdf.CellFormat(35, 8, "", "1", 0, "", false, 0, "")
		pdf.CellFormat(90, 8, "รวมเงิน", "1", 0, "R", false, 0, "")

		pdf.CellFormat(35, 8, closePriceComma, "1", 0, "C", false, 0, "")
		pdf.CellFormat(15, 8, "", "1", 0, "", false, 0, "")
		// End of Table

		pdf.Ln(-1)

		// แปลง TotalPrice เป็น float
		closeprice, _ := strconv.ParseFloat(v.ClosePrice, 64)
		// imcvat := float64(float64(closeprice) * 1.07)
		prevat := float64(float64(closeprice) * 0.07)

		// formattedVatPrice := common.FloatWithCommas(imcvat, 2)
		formattedPrePrice := common.FloatWithCommas(prevat, 2)

		formattedCloseprice := common.FloatWithCommas(closeprice, 2)

		// แสดงผลลัพธ์ที่ได้ใน PDF

		pdf.SetFont("Sarabun", "B", 16)
		pdf.SetXY(15, 177)
		pdf.SetFillColor(255, 255, 255)
		pdf.SetDrawColor(0, 0, 0)
		// pdf.Text(120, 80, "ระบุสถานที่ประมูล ")
		pdf.CellFormat(35, 6, "ประมูลขายทอดตลาดได้รวม 1 คัน เป็นเงิน", "0", 0, "L", true, 0, "")

		pdf.SetXY(130, 177)
		pdf.SetFillColor(255, 255, 255)
		pdf.SetDrawColor(0, 0, 0)
		pdf.CellFormat(50, 6, formattedCloseprice+" บาท", "0", 0, "R", true, 0, "")

		pdf.SetXY(15, 184)
		pdf.SetFillColor(255, 255, 255)
		pdf.SetDrawColor(0, 0, 0)
		pdf.CellFormat(35, 6, "ภาษีมูลค่าเพิ่ม 7 %", "0", 0, "L", true, 0, "")

		pdf.SetXY(130, 184)
		pdf.SetFillColor(255, 255, 255)
		pdf.SetDrawColor(0, 0, 0)
		pdf.CellFormat(50, 6, formattedPrePrice+" บาท", "0", 0, "R", true, 0, "")

		pdf.SetXY(15, 191)
		pdf.SetFillColor(255, 255, 255)
		pdf.SetDrawColor(0, 0, 0)
		pdf.CellFormat(35, 8, "บริษัทฯ หักเงินเพื่อชำระเป็น", "0", 0, "L", true, 0, "")

		pdf.SetXY(25, 197)
		pdf.SetFillColor(255, 255, 255)
		pdf.SetDrawColor(0, 0, 0)
		pdf.CellFormat(35, 8, "ภาษีมูลค่าเพิ่ม (ที่ทางบริษัทฯนำส่ง ภ.พ.36 ให้)", "0", 0, "L", true, 0, "")

		pdf.SetXY(130, 199)
		pdf.SetFillColor(255, 255, 255)
		pdf.SetDrawColor(0, 0, 0)
		pdf.CellFormat(50, 4, formattedPrePrice+" บาท", "0", 0, "R", true, 0, "")

		pdf.SetXY(25, 205)
		pdf.SetFillColor(255, 255, 255)
		pdf.SetDrawColor(0, 0, 0)
		pdf.CellFormat(35, 4, "ค่าอื่นๆ", "0", 0, "L", true, 0, "")

		pdf.SetXY(25, 211)
		pdf.SetFillColor(255, 255, 255)
		pdf.SetDrawColor(0, 0, 0)
		pdf.CellFormat(35, 4, "รวมเป็นเงินที่หัก", "0", 0, "L", true, 0, "")

		pdf.SetXY(130, 211)
		pdf.SetFillColor(255, 255, 255)
		pdf.SetDrawColor(0, 0, 0)
		pdf.CellFormat(50, 4, formattedPrePrice+" บาท", "0", 0, "R", true, 0, "")

		pdf.SetXY(15, 219)
		pdf.SetFillColor(255, 255, 255)
		pdf.SetDrawColor(0, 0, 0)
		pdf.CellFormat(35, 4, "คงเหลือเงินส่งคืนจำนวน", "0", 0, "L", true, 0, "")

		pdf.SetXY(130, 219)
		pdf.SetFillColor(255, 255, 255)
		pdf.SetDrawColor(0, 0, 0)
		pdf.CellFormat(50, 4, formattedCloseprice+" บาท", "0", 0, "R", true, 0, "")

		pdf.SetXY(130, 232)
		pdf.SetFillColor(255, 255, 255)
		pdf.SetDrawColor(0, 0, 0)
		pdf.CellFormat(35, 5, "ขอแสดงความนับถือ", "0", 0, "C", true, 0, "")

		pdf.SetXY(130, 260)
		pdf.SetFillColor(255, 255, 255)
		pdf.SetDrawColor(0, 0, 0)
		pdf.CellFormat(35, 5, "ลงชื่อ....................................", "0", 0, "C", true, 0, "")

		pdf.SetXY(130, 268)
		pdf.SetFillColor(255, 255, 255)
		pdf.SetDrawColor(0, 0, 0)
		pdf.CellFormat(35, 5, "บริษัท โอมาคาเสะ คาร์ จำกัด", "0", 0, "C", true, 0, "")

	}
	pdf.Ln(-1)

	path, err := common.UploadPdfToGoogle(pdf, vehicles[0].SourceName, "auction", "fourwd-auction")
	if err != nil {
		return "", err
	}
	return path, nil

}
