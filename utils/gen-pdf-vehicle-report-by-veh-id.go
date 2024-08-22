package utils

import (
	"strconv"

	"github.com/FourWD/middleware/common"
	"github.com/jung-kurt/gofpdf"
)

func GenPDFReports(auctionID string, vehicleID []string) (string, error) { //ใบรายงานผลการประมูล สรรพากร

	vehicles, summary := prepareVehicleResults(auctionID, vehicleID)

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
		pdf.CellFormat(30, 8, "เลขที่ใบรับรถ", "1", 0, "C", true, 0, "")
		pdf.CellFormat(80, 8, "รายละเอียด", "1", 0, "C", true, 0, "")
		pdf.CellFormat(20, 8, "ผลประมูล", "1", 0, "C", true, 0, "")
		pdf.CellFormat(30, 8, "ราคาที่ประมูล", "1", 0, "C", true, 0, "")
		pdf.CellFormat(15, 8, "ค่าอื่นๆ", "1", 0, "C", true, 0, "")

		pdf.Ln(-1)
		//detail
		pdf.SetFont("Sarabun", "B", 15)

		pdf.Text(16, 113, v.VehicleNo)
		pdf.CellFormat(15, 60, "", "1", 0, "TC", false, 0, "")
		pdf.Text(27, 113, v.Receipt)
		pdf.CellFormat(30, 60, "", "1", 0, "TC", false, 0, "")

		var test1 string
		if v.License != "" && v.LicenseProvinceName != "" {
			test1 = "ทะเบียน : " + v.License + " " + v.LicenseProvinceName
		} else {
			test1 = "ทะเบียน : " + v.ChassisNo
		}
		test2 := "ยี่ห้อ : " + v.VehicleBrandName + " รุ่น : " + v.VehicleModelName
		test4 := "รุ่นย่อย : " + v.VehicleSubModelName
		test6 := "ปี : " + v.YearRegister + " เกียร์ : " + v.VehicleGearName
		test8 := "เลขเครื่อง : " + v.EngineNo
		test9 := "เลขตัวถัง : " + v.ChassisNo
		test10 := "เลขที่สัญญา  : " + v.CN

		pdf.MultiCell(80, 7, test1+"\n"+test2+"\n"+test4+"\n"+test6+"\n"+test8+"\n"+test9+"\n"+test10+"\n", "0", "L", false)

		closePriceFloat, _ := strconv.ParseFloat(v.ClosePrice, 64)
		closePriceComma := common.FloatWithCommas(closePriceFloat, 2)

		pdf.SetXY(135, 108) // ตั้งตำแหน่ง X และ Y ใหม่
		pdf.Text(140, 113, "ขายได้")
		pdf.CellFormat(20, 60, "", "1", 0, "TC", false, 0, "")
		// pdf.Text(154, 113, closePriceComma)
		pdf.CellFormat(30, 60, closePriceComma, "1", 0, "TC", false, 0, "")
		pdf.CellFormat(15, 60, "", "1", 0, "TC", false, 0, "")
		pdf.Ln(-1)

		pdf.SetFont("Sarabun", "B", 15)
		pdf.CellFormat(15, 8, "", "1", 0, "", false, 0, "")
		pdf.CellFormat(30, 8, "", "1", 0, "", false, 0, "")
		pdf.CellFormat(100, 8, "รวมเงิน", "1", 0, "R", false, 0, "")

		pdf.CellFormat(30, 8, closePriceComma, "1", 0, "C", false, 0, "")
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

	var fileName string
	if len(vehicleID) == 0 {
		fileName = summary.AuctionName
	} else {
		for _, data := range vehicles {
			fileName = summary.AuctionName + "_" + data.License
		}
	}

	// localPath := "images/pdf" // localprint
	// path, err := SavePdf(pdf, fileName, localPath)
	// if err != nil {
	// 	return "", err
	// }

	path, err := common.UploadPdfToGoogle(pdf, fileName, "auction", "fourwd-auction")
	if err != nil {
		return "", err
	}
	return path, nil

}

func prepareVehicleResults(auctionID string, vehicleID []string) ([]VehicleReport, Summary) {

	var vehicles []VehicleReport
	if len(vehicleID) != 0 {
		common.Database.Raw(`SELECT v.engine_no,v.chassis_no,v.contract_number as cn,v.vehicle_auction_receipt as receipt, DATE(a.start_date) as start_date_auction,DATE(a.end_date) as end_date_auction,b.text_color,b.background_color,s.name as source_name,vg.color_code as color_code,v.id as vehicle_id, v.license_receive_date,(r.crp-(r.crp*7/100)) AS crp_p_vat,	r.crp as crp,
	v.vehicle_grade_id, v.vehicle_gear_name, v.vehicle_brand_name,vb.name_en as vehicle_brand_name_th,v.vehicle_color_name,
	v.vehicle_model_name, av.open_price,av.close_price as close_price, v.year_manufacturing as years,v.year_register, vm.name_en as vehicle_model_name_th,v.vehicle_sub_model_name,v.license,
	  v.license_province_id, v.license_province_name, v.mile AS mile, b.label as branch_label,v.remark,
COALESCE(v.vehicle_grade_id, '') AS vehicle_grade_id,
COALESCE(v.vehicle_grade_value, '') AS vehicle_grade_value,
COALESCE(av.vehicle_no, 0) AS vehicle_no,
vi.image_path AS image_preview_path
FROM auction_vehicles AS av
LEFT JOIN auctions AS a ON av.auction_id = a.id
LEFT JOIN vehicles AS v ON  v.id = av.vehicle_id
LEFT JOIN vehicle_grades AS vg ON v.vehicle_grade_id = vg.id
LEFT JOIN sources AS s ON v.source_id = s.id
LEFT JOIN (
	SELECT DISTINCT auction_id , chassis_no , license , crp , redbook
FROM redbooks
 ) r ON av.auction_id = r.auction_id AND v.chassis_no = r.chassis_no AND v.license = r.license
LEFT JOIN branches b ON v.branch_id = b.id
LEFT JOIN vehicle_brands vb ON v.vehicle_brand_id = vb.id
LEFT JOIN vehicle_models vm ON v.vehicle_model_id = vm.id
LEFT JOIN vehicle_images AS vi ON av.vehicle_id = vi.vehicle_id AND vi.template_vehicle_image_id = '4b8fe630-a2b6-4470-9e86-3825c169a8f5' AND vi.is_delete = 0
						WHERE av.auction_id = ? AND av.vehicle_id IN ? AND av.is_win = true
							ORDER BY av.vehicle_no ASC`, auctionID, vehicleID).Scan(&vehicles)
	}
	// common.Print(auctionID, "auctionID")
	// common.Print("vehicle", common.StructToString(vehicles))

	summary := new(Summary)
	common.Database.Raw(`
	SELECT DATE(a.start_date) as start_date_auction,COUNT(*) AS total_car, 
    SUM(av.close_price) AS total_price, 
    MAX(a.name) AS auction_name
    FROM
        auction_vehicles av
        LEFT JOIN vehicles v on av.vehicle_id = v.id
        LEFT JOIN auctions a ON av.auction_id = a.id 
	where
	  av.auction_id = ? AND av.is_win = true  GROUP BY a.start_date`, auctionID).Scan(&summary)
	// common.Print(auctionID, "auctionID")

	return vehicles, *summary
}
