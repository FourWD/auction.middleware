package utils

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/FourWD/middleware/common"
	"github.com/jung-kurt/gofpdf"
	"github.com/jung-kurt/gofpdf/contrib/httpimg"
)

// func GenPDFVehicleSKM(auctionID string) (string, error) {
// 	pdf := gofpdf.New("P", "mm", "A3", "")

// 	GenPDFDownloadVehicle(pdf, auctionID)

// 	path, err := common.UploadPdfToGoogle(pdf, "ใบ สคบ.", "auction", "fourwd-auction")
// 	if err != nil {
// 		return "", err
// 	}
// 	return path, nil
// }

// type Image struct {
// 	VehicleID string `json:"vehicle_id"`
// 	ImageName string `json:"image_name"`
// 	ImagePath string `json:"image_path"`
// }

type Vehicle struct {
	AdditionalCode       string `json:"additional_code"`
	VehicleID            string `json:"vehicle_id"`
	VehicleAuctionRecipt string `json:"vehicle_auction_receipt"`
	VehicleGradeID       string `json:"vehicle_grade_id"`
	VehicleBrandName     string `json:"vehicle_brand_name"`
	VehicleBrandNameTh   string `json:"vehicle_brand_name_th"`

	VehicleSubModelNameTh string `json:"vehicle_sub_model_name_th"`
	VehicleSubModelName   string `json:"vehicle_sub_model_name"`

	VehicleModelName   string `json:"vehicle_model_name"`
	VehicleModelNameTh string `json:"vehicle_model_name_th"`

	VehicleColorName string `json:"vehicle_color_name"`
	Years            string `json:"years"`
	YearRegister     string `json:"year_register"`
	ImagePreviewPath string `json:"image_preview_path"`
	OpenPrice        string `json:"open_price"`
	SKU              string `json:"sku"`
	EngineSize       string `json:"engine_size"`
	EngineNo         string `json:"engine_no"`

	EngineCapacity      string `json:"engine_capacity"`
	VehicleFuelTypeName string `json:"vehicle_fuel_type_name"`
	Mile                string `json:"mile"`
	ChassisNo           string `json:"chassis_no"`
	VehicleTypeName     string `json:"vehicle_type_name"`
	VehicleSubTypeName  string `json:"vehicle_sub_type_name"`
	SourceName          string `json:"source_name"`
	SourceID            string `json:"source_id"`

	Branch      string `json:"branch"`
	BranchLabel string `json:"branch_label"`

	License             string `json:"license"`
	LicenseProvinceName string `json:"license_province_name"`

	ImagePath []Image
}

func GenPDFDownloadVehicle(auctionID string) (string, error) {
	vehicles := prepareVehicleSKM(auctionID)

	filepathStr := "images/pdf/"
	type Qr struct {
		AuctionID string `json:"auction_id"`
		VehicleID string `json:"vehicle_id"`
	}
	// fileextention := ".pdf"
	pdf := gofpdf.New("P", "mm", "A3", "")
	for _, v := range vehicles {

		pdf.AddPage()
		pdf.AddUTF8Font("Sarabun", "", "fonts/THSarabun.ttf")
		pdf.AddUTF8Font("Sarabun", "B", "fonts/THSarabunBold.ttf")

		x := 9.00
		y := 6.00
		width := 40.00
		height := 22.00

		imagePath := filepathStr + "simple.png"
		pdf.Image(imagePath, x, y, width, height, false, "", 0, "")

		// Set position for the text
		topHandle := filepathStr + "REDB.png"
		pdf.Image(topHandle, 60, 34, 275, 3, false, "", 0, "")

		pdf.SetTextColor(0, 0, 0)
		pdf.SetFont("Sarabun", "B", 48)
		pdf.Text(60, 17, "ใบรายละเอียดรถยนต์")

		pdf.SetFont("Sarabun", "B", 26)
		pdf.Text(60, 30, "ชนิดสินค้ารถยนต์ที่ใช้แล้ว")
		pdf.SetFont("Sarabun", "B", 24)
		pdf.SetTextColor(169, 169, 169)
		pdf.Text(126, 30, "ดูข้อมูลและรูปภาพรถยนต์เพิ่มเติมได้ที่")
		pdf.SetTextColor(0, 0, 0)
		pdf.Text(216, 30, "www.omakasecar.com")

		// กำหนดสีของเส้น
		pdf.SetDrawColor(0, 0, 0)

		//grade
		pdf.SetFont("Sarabun", "", 28)
		pdf.Text(15, 53, "เกรด")
		pdf.SetFont("Sarabun", "B", 80)
		pdf.Text(15, 75, v.VehicleGradeID)

		pdf.SetLineWidth(1.5)
		pdf.SetDrawColor(0, 0, 0)
		pdf.Line(38, 48, 38, 84) // เส้นกั้นใหม่แนวตั้ง

		//gray4box+26

		//text box1
		gray4box1 := filepathStr + "gray4.png"
		pdf.Image(gray4box1, 90, 50, 200, 24, false, "", 0, "")
		//text box2
		gray4box2 := filepathStr + "gray4.png"
		pdf.Image(gray4box2, 90, 76, 200, 24, false, "", 0, "")
		//text box3
		gray4box3 := filepathStr + "gray4.png"
		pdf.Image(gray4box3, 90, 102, 200, 24, false, "", 0, "")
		// // //text box4
		gray4box4 := filepathStr + "gray4.png"
		pdf.Image(gray4box4, 90, 128, 99, 24, false, "", 0, "")
		// //text box5
		gray4box5 := filepathStr + "gray4.png"
		pdf.Image(gray4box5, 191, 128, 99, 24, false, "", 0, "")

		//text box6
		gray4box6 := filepathStr + "gray4.png"
		pdf.Image(gray4box6, 8, 154, 140, 24, false, "", 0, "")
		// text box7
		gray4box7 := filepathStr + "gray4.png"
		pdf.Image(gray4box7, 150, 154, 140, 24, false, "", 0, "")
		//text box8
		gray4box8 := filepathStr + "gray4.png"
		pdf.Image(gray4box8, 8, 180, 140, 24, false, "", 0, "")
		// // text box9
		gray4box9 := filepathStr + "gray4.png"
		pdf.Image(gray4box9, 150, 180, 140, 24, false, "", 0, "")
		// //text box10
		gray4box10 := filepathStr + "gray4.png"
		pdf.Image(gray4box10, 8, 206, 140, 24, false, "", 0, "")
		// // text box11
		gray4box11 := filepathStr + "gray4.png"
		pdf.Image(gray4box11, 150, 206, 140, 24, false, "", 0, "")

		// //text box12
		gray4box12 := filepathStr + "gray4.png"
		pdf.Image(gray4box12, 8, 232, 140, 24, false, "", 0, "")
		// // text box13
		gray4box13 := filepathStr + "gray4.png"
		pdf.Image(gray4box13, 150, 232, 140, 24, false, "", 0, "")

		//text box14
		gray4box14 := filepathStr + "gray4.png"
		pdf.Image(gray4box14, 8, 258, 140, 24, false, "", 0, "")
		// text box15
		gray4box15 := filepathStr + "gray4.png"
		pdf.Image(gray4box15, 150, 258, 140, 24, false, "", 0, "")

		//text box16
		gray4box16 := filepathStr + "gray4.png"
		pdf.Image(gray4box16, 8, 284, 140, 24, false, "", 0, "")
		// text box17
		gray4box17 := filepathStr + "gray4.png"
		pdf.Image(gray4box17, 150, 284, 140, 24, false, "", 0, "")

		//text boป18
		gray4box18 := filepathStr + "gray4.png"
		pdf.Image(gray4box18, 8, 310, 140, 24, false, "", 0, "")
		// text box19
		gray4box19 := filepathStr + "gray4.png"
		pdf.Image(gray4box19, 150, 310, 140, 24, false, "", 0, "")

		//text box18
		gray4box20 := filepathStr + "gray4.png"
		pdf.Image(gray4box20, 8, 336, 140, 24, false, "", 0, "")
		// text box19
		gray4box21 := filepathStr + "gray4.png"
		pdf.Image(gray4box21, 150, 336, 140, 24, false, "", 0, "")

		//text box22
		gray4box22 := filepathStr + "gray4.png"
		pdf.Image(gray4box22, 8, 362, 140, 24, false, "", 0, "")
		// text box23
		gray4box23 := filepathStr + "gray4.png"
		pdf.Image(gray4box23, 150, 362, 140, 24, false, "", 0, "")

		//text box24
		gray4box24 := filepathStr + "gray4.png"
		pdf.Image(gray4box24, 8, 388, 140, 24, false, "", 0, "")
		// text box25
		gray4box25 := filepathStr + "gray4.png"
		pdf.Image(gray4box25, 150, 388, 140, 24, false, "", 0, "")

		//text detail+22
		pdf.SetFont("Sarabun", "B", 24)
		pdf.Text(90, 46, "Vehicle code: "+v.SKU)

		pdf.SetFont("Sarabun", "", 18)
		pdf.Text(94, 58, "ยี่ห้อรถ")
		pdf.SetFont("Sarabun", "B", 50)
		pdf.Text(94, 72, v.VehicleBrandNameTh)

		pdf.SetFont("Sarabun", "", 18)
		pdf.Text(94, 84, "รุ่นรถ")
		pdf.SetFont("Sarabun", "B", 50)
		// pdf.Text(94, 98, v.VehicleModelNameTh+" "+v.VehicleSubModelName)
		pdf.Text(94, 98, v.VehicleModelNameTh)

		pdf.SetFont("Sarabun", "", 18)
		pdf.Text(94, 110, "ราคาตั้งประมูล ( ไม่รวม VAT )")
		pdf.SetFont("Sarabun", "B", 50)
		price, _ := strconv.ParseFloat(v.OpenPrice, 64)
		openprice := common.FloatWithCommas(price, 0)
		pdf.Text(94, 122, openprice+" ( ไม่รวม VAT )")
		pdf.Text(94, 122, " ")

		pdf.SetFont("Sarabun", "", 18)
		pdf.Text(94, 137, "ดาวน์")
		pdf.SetFont("Sarabun", "B", 50)
		pdf.Text(94, 148, "")

		pdf.SetFont("Sarabun", "", 18)
		pdf.Text(193, 137, "ผ่อน (บาท)")
		pdf.SetFont("Sarabun", "B", 50)
		pdf.Text(193, 148, "")

		// if v.ImagePreviewPath != "" {
		// 	httpimg.Register(pdf, v.ImagePreviewPath, "")
		// 	pdf.Image(v.ImagePreviewPath, 12, 95, 70, 45, false, "", 0, "")
		// }
		if v.ImagePreviewPath != "" {
			httpimg.Register(pdf, v.ImagePreviewPath, "")
			pdf.Image(v.ImagePreviewPath, 9, 90, 74, 0, false, "", 0, "")
		}
		//บน 25
		//ล่าง 26
		pdf.SetFont("Sarabun", "", 18)
		pdf.Text(12, 162, "รุ่นปี")
		pdf.SetFont("Sarabun", "B", 50)
		pdf.Text(12, 174, v.Years)

		pdf.SetFont("Sarabun", "", 18)
		pdf.Text(154, 162, "สี")
		pdf.SetFont("Sarabun", "B", 50)
		pdf.Text(154, 174, v.VehicleColorName)
		//
		pdf.SetFont("Sarabun", "", 18)
		pdf.Text(12, 187, "ยี่ห้อเครื่องยนต์")
		pdf.SetFont("Sarabun", "B", 50)
		pdf.Text(12, 200, v.VehicleBrandNameTh)

		pdf.SetFont("Sarabun", "", 18)
		pdf.Text(154, 187, "ขนาดหรือน้ำหนัก")
		pdf.SetFont("Sarabun", "B", 50)
		engine, _ := strconv.ParseFloat(v.EngineSize, 64)
		enginecomma := common.FloatWithCommas(engine, 0)
		pdf.Text(154, 200, enginecomma+" "+"ซีซี")
		//
		pdf.SetFont("Sarabun", "", 18)
		pdf.Text(12, 212, "ความจุกระบอกสูบ")
		pdf.SetFont("Sarabun", "B", 50)
		pdf.Text(12, 226, v.EngineCapacity)

		pdf.SetFont("Sarabun", "", 18)
		pdf.Text(154, 212, "ชนิดเชื้อเพลิง")
		pdf.SetFont("Sarabun", "B", 50)
		pdf.Text(154, 226, v.VehicleFuelTypeName)
		//
		pdf.SetFont("Sarabun", "", 18)
		pdf.Text(12, 238, "ระยะทางการใช้งาน")
		pdf.SetFont("Sarabun", "B", 50)
		mile, _ := strconv.ParseFloat(v.Mile, 64)
		milecomma := common.FloatWithCommas(mile, 0) + " " + "กม."
		pdf.Text(12, 252, milecomma)

		pdf.SetFont("Sarabun", "", 18)
		pdf.Text(154, 238, "วันที่จดทะเบียน")
		pdf.SetFont("Sarabun", "B", 50)
		pdf.Text(154, 252, v.YearRegister)
		//
		pdf.SetFont("Sarabun", "", 18)
		pdf.Text(12, 264, "เลขตัวรถ")
		pdf.SetFont("Sarabun", "B", 50)
		pdf.Text(12, 278, v.ChassisNo)

		pdf.SetFont("Sarabun", "", 18)
		pdf.Text(154, 264, "เลขเครื่องยนต์")
		pdf.SetFont("Sarabun", "B", 50)
		pdf.Text(154, 278, v.EngineNo)
		//
		pdf.SetFont("Sarabun", "", 18)
		pdf.Text(12, 290, "เลขทะเบียน")
		pdf.SetFont("Sarabun", "B", 50)
		pdf.Text(12, 304, v.License+" "+v.LicenseProvinceName)

		pdf.SetFont("Sarabun", "", 18)
		pdf.Text(154, 290, "ข้อมูลการประสบภัย (ถ้ามี)")
		pdf.SetFont("Sarabun", "B", 50)
		pdf.Text(154, 307, "ไม่มี")
		//

		pdf.SetFont("Sarabun", "", 18)
		pdf.Text(12, 316, "ชื่อประเภทหรือชนิดสินค้า")
		pdf.SetFont("Sarabun", "B", 50)
		pdf.Text(12, 332, v.VehicleSubTypeName)

		pdf.SetFont("Sarabun", "", 18)
		pdf.Text(154, 316, "ประเทศผู้ผลิต")
		pdf.SetFont("Sarabun", "B", 50)
		pdf.Text(154, 333, "ไม่มี")
		//
		pdf.SetFont("Sarabun", "", 18)
		pdf.Text(12, 342, "ชื่อ-นามสกุล")
		pdf.SetFont("Sarabun", "B", 50)
		pdf.Text(12, 356, v.SourceID)

		pdf.SetFont("Sarabun", "", 18)
		pdf.Text(154, 342, "ลำดับของเจ้าของรถยนต์")
		pdf.SetFont("Sarabun", "B", 50)
		pdf.Text(154, 359, "ไม่มี")
		//
		pdf.SetFont("Sarabun", "", 18)
		pdf.Text(12, 368, "ภาระผูกพันของรถยนต์ที่มีอยู่ในวันจำหน่าย (ถ้ามี)")
		pdf.SetFont("Sarabun", "B", 50)
		pdf.Text(12, 385, "ไม่มี")

		pdf.SetFont("Sarabun", "", 18)
		pdf.Text(154, 368, "สมุดคู่มือการบำรุงรักษารถยนต์ (ถ้ามี)")
		pdf.SetFont("Sarabun", "B", 50)
		pdf.Text(154, 385, "ไม่มี")
		//
		pdf.SetFont("Sarabun", "", 18)
		pdf.Text(12, 394, "ผู้ประกอบธุรกิจ")
		pdf.SetFont("Sarabun", "B", 50)
		pdf.Text(12, 408, "บจก.โอมาคาเสะ คาร์")

		pdf.SetFont("Sarabun", "", 18)
		pdf.Text(154, 394, "สถานที่ประกอบการ")
		pdf.SetFont("Sarabun", "B", 50)
		pdf.Text(154, 408, v.BranchLabel)

		if v.AdditionalCode == "06" {
			return "", errors.New("G001: Auction has ended, cannot generate PDF")
		}
		//qr
		qr := Qr{
			AuctionID: auctionID,
			VehicleID: v.VehicleID,
		}
		buf, _ := common.GenBufferQrPdf(common.StructToString(qr))
		pdf.RegisterImageReader(fmt.Sprintf("qr_%s", v.VehicleID), "jpg", &buf)
		xx, yy := 45.0, 46.0
		widthh, heightt := 40.0, 40.0
		pdf.Image(fmt.Sprintf("qr_%s", v.VehicleID), xx, yy, widthh, heightt, false, "", 0, "")
		common.Print("vehicle_iddddddddddddddddddddddd", common.StructToString(qr))

		// url := common.GenBufferQRPDF(common.StructToString(qr))
		pdf.SetTextColor(255, 0, 0)
		// pdf.Text(15, 275, url)
		pdf.SetTextColor(0, 0, 0)
		// webScan := filepathStr + url
		// pdf.Image(webScan, 45, 52, 40, 40, false, "", 0, "")

		// //down
		// pdf.SetTextColor(255, 255, 255)

		downHandle := filepathStr + "down.png"
		pdf.Image(downHandle, 0, 415, 300, 15, false, "", 0, "")

		y += 5.00
		pdf.Ln(5)
		//loop page

	}

	// filedesination := filepathStr + "ใบรายละเอียดรถยนต์ - สคบ." + fileextention
	// err := pdf.OutputFileAndClose(filedesination)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// }
	// return filedesination, nil
	path, err := common.UploadPdfToGoogle(pdf, "ใบสคบ.", "auction", "fourwd-auction")
	if err != nil {
		return "", err
	}
	return path, nil

}
func prepareVehicleSKM(auctionID string) []Vehicle { //sql skm
	var vehicles []Vehicle
	common.Database.Raw(`
        SELECT 
            CASE WHEN a.auction_status_id = '06' THEN 'G001' ELSE '' END AS additional_code,
            a.auction_status_id,
            v.id as vehicle_id, 
            v.license_receive_date,
            v.sku, 
            v.vehicle_grade_id, 
            v.vehicle_brand_name, 
            vm.name_en as vehicle_brand_name_th,
            vb.name_en as vehicle_brand_name_th,
            v.vehicle_model_name, 
            a.open_price,
            v.year_manufacturing as years, 
            v.vehicle_color_name,
            v.engine_size,
            engine_capacity,
            v.vehicle_fuel_type_name,
            v.period_of_use,
            v.chassis_no,
            v.engine_no,
            v.vehicle_type_name,
            v.vehicle_sub_type_name,
            v.year_register,
            vm.name_en as vehicle_model_name_th,
            v.vehicle_sub_model_name, 
            v.source_name,
            v.source_id,
            v.license,
            v.license_province_id,
            v.license_province_name,
            v.mile AS mile,
            b.label as branch_label,
            v.branch_name as branch,
            COALESCE(v.vehicle_grade_id, '') AS vehicle_grade_id,
            COALESCE(v.vehicle_grade_value, '') AS vehicle_grade_value,
            COALESCE(a.vehicle_no, 0) AS vehicle_no,
            vi.image_path AS image_preview_path ,
            a.bidding_step_1,
            a.bidding_step_2,
            a.bidding_step_3,
            auc.start_date,
            auc.end_date,
            p.price AS proxy
        FROM 
            auction_vehicles AS a
        LEFT JOIN 
            vehicles AS v ON a.vehicle_id = v.id
        LEFT JOIN 
            branches b ON v.branch_id = b.id
        LEFT JOIN 
            vehicle_brands vb ON v.vehicle_brand_id = vb.id
        LEFT JOIN 
            vehicle_models vm ON v.vehicle_model_id = vm.id
        LEFT JOIN 
            vehicle_images AS vi ON a.vehicle_id = vi.vehicle_id AND vi.template_vehicle_image_id = '4b8fe630-a2b6-4470-9e86-3825c169a8f5' 
            AND vi.is_delete = 0
        LEFT JOIN 
            auctions AS auc ON a.auction_id = auc.id
        LEFT JOIN 
            auction_vehicle_users AS a_fav ON a_fav.auction_id = auc.id AND a_fav.vehicle_id = v.id and a_fav.user_id = ''
        LEFT JOIN 
            proxies p ON a_fav.proxy_id = p.id AND p.id != ''
        WHERE 
            a.auction_id = ? AND v.id != '' AND a.vehicle_no != 0 
        ORDER BY 
            a.vehicle_no ASC`, auctionID).Scan(&vehicles)

	for i := range vehicles {
		if vehicles[i].AdditionalCode == "G001" {
			fmt.Println("Additional code is G001")
		} else {
			fmt.Println("Additional code is not G001")
		}
	}

	return vehicles
}
