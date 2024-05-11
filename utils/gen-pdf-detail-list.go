package utils

import (
	"strconv"
	"strings"

	"github.com/FourWD/middleware/common"
	"github.com/jung-kurt/gofpdf"
	"github.com/jung-kurt/gofpdf/contrib/httpimg"
)

// func GenPDFDetailList(auctionID string) (string, error) { //list รายการราคารถ
// 	pdf := gofpdf.New("P", "mm", "A4", "")

// 	// auctionID := c.Params("auction_id")
// 	GenPDFVehicleDetail(pdf, auctionID)

// 	path, err := common.UploadPdfToGoogle(pdf, "ใบลิสต์รายการราคารถ", "auction", "fourwd-auction")
// 	if err != nil {
// 		return "", err
// 	}
// 	return path, nil

// }

type VehicleSummaryDetail struct {
	RoundID              string `json:"round_id"`
	ID                   string `json:"id"`
	CodeAuction          string `json:"code_auction"`
	ClosePrice           string `json:"close_price"`
	CarImage             string `json:"car_image"`
	LicenseImage         string `json:"license_image"`
	VehicleAuctionRecipt string `json:"vehicle_auction_receipt"`
	VehicleGradeID       string `json:"vehicle_grade_id"`
	VehicleNo            string `json:"vehicle_no"`
	TotalCar             string `json:"total_car"`
	VehicleBrandName     string `json:"vehicle_brand_name"`
	VehicleSubModelName  string `json:"vehicle_model_name"`
	VehicleModelName     string `json:"vehicle_sub_model_name"`
	VehicleColorName     string `json:"vehicle_color_name"`
	SourceName           string `json:"source_name"`

	ColorCode string `json:"color_code"`

	Years               string `json:"years"`
	YearRegister        string `json:"year_register"`
	ImagePreviewPath    string `json:"image_preview_path"`
	OpenPrice           string `json:"open_price"`
	AuctionName         string `json:"auction_name"`
	TaxExprieDate       string `json:"tax_expire_date"`
	Mile                string `json:"mile"`
	ChassisNo           string `json:"chassis_no"`
	EngineSize          string `json:"engine_size"`
	EngineNo            string `json:"engine_no"`
	VehicleGearName     string `json:"vehicle_gear_Name"`
	Remark              string `json:"remark"`
	License             string `json:"license"`
	BranchLabel         string `json:"branch_label"`
	CRPPVat             string `json:"crp_p_vat"`
	CRP                 string `json:"crp"`
	LicenseProvinceName string `json:"license_province_name"`
	BackgroundColor     string `json:"background_color"`
	TextColor           string `json:"text_color"`
}

func headertable(pdf gofpdf.Pdf, tabley int) {
	// Define table headers
	pdf.SetFillColor(196, 16, 22)
	pdf.SetTextColor(255, 255, 255)
	pdf.SetDrawColor(0, 0, 0)
	pdf.SetLineWidth(0.2)
	pdf.SetFont("Sarabun", "B", 12)
	tableX := 0
	tableY := tabley
	pdf.SetXY(float64(tableX), float64(tableY))

	// border := "10"
	pdf.CellFormat(10, 8, "ลำดับ", "1", 0, "C", true, 0, "")
	pdf.CellFormat(20, 8, "สถานที่จอด", "1", 0, "C", true, 0, "")
	pdf.CellFormat(10, 8, "ช่อง", "1", 0, "C", true, 0, "")
	// pdf.CellFormat(10, 8, "ช่อง", "1", 0, "C", true, 0, "")
	pdf.CellFormat(66, 8, "รุ่นรถ", "1", 0, "C", true, 0, "")
	pdf.CellFormat(20, 8, "สี", "1", 0, "C", true, 0, "")
	pdf.CellFormat(10, 8, "รุ่นปี", "1", 0, "C", true, 0, "")
	pdf.CellFormat(15, 8, "เกียร์", "1", 0, "C", true, 0, "")
	pdf.CellFormat(10, 8, "เกรด", "1", 0, "C", true, 0, "")
	pdf.CellFormat(20, 8, "เลขไมล์", "1", 0, "C", true, 0, "")
	pdf.CellFormat(30, 8, "ทะเบียนรถ", "1", 0, "C", true, 0, "")
	pdf.CellFormat(25, 8, "ราคาประมูลเริ่มต้น", "1", 0, "C", true, 0, "")
	pdf.CellFormat(12, 8, "%CRP", "1", 0, "C", true, 0, "")
	pdf.CellFormat(35, 8, "ราคาป้ายแดง (ไม่รวม vat)", "1", 0, "C", true, 0, "")
	pdf.CellFormat(15, 8, "หมายเหตุ", "1", 0, "C", true, 0, "")
	pdf.SetTextColor(0, 0, 0)
	pdf.SetFillColor(255, 255, 255)

}

func GenPDFVehicleDetail(auctionID string) (string, error) {
	vehicles := prepareDetailList(auctionID)

	// filepathStr := "images/pdf/"
	// fileextention := ".pdf"
	pdf := gofpdf.New("L", "mm", "A4", "")

	pdf.AddPage()
	pdf.AddUTF8Font("Sarabun", "", "fonts/THSarabun.ttf")
	pdf.AddUTF8Font("Sarabun", "B", "fonts/THSarabunBold.ttf")

	// header := filepathStr + "top-bar-detail.jpg"
	// pdf.Image(header, 0, 0, 297, 45, false, "", 0, "")
	header := "https://storage.googleapis.com/fourwd-auction/app/pdf_resource/top-list.jpg"
	httpimg.Register(pdf, header, "")
	pdf.Image(header, 0, 0, 297, 53, false, "", 0, "")

	headertable(pdf, 53)

	pdf.SetFont("Sarabun", "B", 12)
	tableYz := 60.5
	page := 1
	counter := make(map[string]int)

	perpage := 16
	line := 1

	for i, v := range vehicles {

		i++

		if page == 1 {
		} else {
			perpage = 22
		}
		// fmt.Printf("index = %d, License = %s", i, v.License)
		// fmt.Printf(", page = %d, perpage = %d", page, perpage)
		// fmt.Printf(", line = %d", line)
		// fmt.Println("========================")

		tableX := 0
		pdf.SetXY(float64(tableX), float64(tableYz))

		pdf.CellFormat(10, 8, v.VehicleNo, "1", 0, "C", true, 0, "")
		pdf.CellFormat(20, 8, v.BranchLabel, "1", 0, "C", true, 0, "")
		if _, ok := counter[v.BranchLabel]; !ok {
			counter[v.BranchLabel] = 1
		} else {
			counter[v.BranchLabel]++
			// if counter[v.BranchLabel] > 7 {
			// 	counter[v.BranchLabel] = 1
			// }
		}
		pdf.CellFormat(10, 8, strconv.Itoa(counter[v.BranchLabel]), "1", 0, "C", true, 0, "")
		pdf.CellFormat(66, 8, v.VehicleBrandName+" "+v.VehicleModelName+" "+v.VehicleSubModelName, "1", 0, "L", true, 0, "")
		pdf.CellFormat(20, 8, v.VehicleColorName, "1", 0, "C", true, 0, "")
		pdf.CellFormat(10, 8, v.Years, "1", 0, "C", true, 0, "")
		pdf.CellFormat(15, 8, v.VehicleGearName, "1", 0, "C", true, 0, "")
		pdf.CellFormat(10, 8, v.VehicleGradeID, "1", 0, "C", true, 0, "")
		mile, _ := strconv.ParseFloat(v.Mile, 64)
		mileComma := common.FloatWithCommas(mile, 0)

		pdf.CellFormat(20, 8, mileComma, "1", 0, "C", true, 0, "")
		pdf.CellFormat(30, 8, v.License, "1", 0, "C", true, 0, "")
		openPriceFloat, _ := strconv.ParseFloat(v.OpenPrice, 64)

		var openprice string
		if openPriceFloat == 0 {
			openprice = "รอกำหนดราคา"
		} else {
			openprice = common.FloatWithCommas(openPriceFloat, 0)
		}
		pdf.CellFormat(25, 8, openprice, "1", 0, "R", true, 0, "")

		crpFloat, _ := strconv.ParseFloat(v.CRP, 64)
		crp := crpFloat * (100.00 / 107.00)
		formattedCRP := common.FloatWithCommas(crp, 0)

		var formattedCRPPer string
		if crp != 0 {
			crpPer := (openPriceFloat / crp) * 100
			formattedCRPPer = common.FloatWithCommas(crpPer, 0)
		} else {
			formattedCRPPer = "0"
		}

		pdf.CellFormat(12, 8, formattedCRPPer+"%", "1", 0, "R", true, 0, "")
		pdf.CellFormat(25, 8, formattedCRP, "1", 0, "R", true, 0, "")
		pdf.CellFormat(20, 8, "", "1", 0, "C", true, 0, "")

		pdf.Ln(-1)
		tableYz += 8

		line++
		newpage := false

		if i == 16 {
			newpage = true

		} else if i > 16 {
			temp := i - 16
			if temp%perpage == 0 {
				newpage = true
			}
		}
		if newpage {
			line = 1
			page++
			if i < len(vehicles) {
				pdf.AddPage()
				tableYz = 8
				headertable(pdf, 0)
			}
		}
		headerdown := "https://storage.googleapis.com/fourwd-auction/app/pdf_resource/down-list.jpg"
		httpimg.Register(pdf, headerdown, "")
		// headerdown := filepathStr + "down-bar-detail.jpg"
		pdf.Image(headerdown, 0, 197, 297, 12, false, "", 0, "")

	}
	pdf.Ln(-1)

	fileName := generateFileNameList(vehicles[0].AuctionName)

	path, err := common.UploadPdfToGoogle(pdf, fileName, "auction", "fourwd-auction")
	if err != nil {
		return "", err
	}
	return path, nil

	// filedestination := filepathStr + "ใบราคา List TEST" + fileextention
	// err := pdf.OutputFileAndClose(filedestination)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// }
	// return filedestination, nil
}

func generateFileNameList(auctionName string) string {
	auctionName = strings.ReplaceAll(auctionName, " ", "_")
	fileName := "ใบ List_" + auctionName
	return fileName
}
func prepareDetailList(auctionID string) []VehicleSummaryDetail {

	var vs []VehicleSummaryDetail
	common.Database.Raw(`SELECT b.text_color,b.background_color,s.name as source_name,vg.color_code as color_code,v.id as vehicle_id, v.license_receive_date,(r.crp-(r.crp*7/100)) AS crp_p_vat,	r.crp as crp,
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
							WHERE av.auction_id = ? AND v.id != '' AND av.vehicle_no != 0
								ORDER BY av.vehicle_no ASC`, auctionID).Scan(&vs)
	// 	common.Database.Raw(`   SELECT *,a.round_id  FROM auction_vehicles av
	//     left join auctions a on av.auction_id = a.id
	//     left join vehicles v on av.vehicle_id = v.id
	//     where a.round_id = ?
	// `, vehicleID).Scan(&vs)
	common.Print(auctionID, "auctionID")
	common.Print("v", common.StructToString(vs))

	return vs
}
