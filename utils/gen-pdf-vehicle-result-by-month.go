package utils

import (
	"fmt"
	"strconv"
	"time"

	"github.com/FourWD/middleware/common"
	"github.com/jung-kurt/gofpdf"
)

// func GenPDFVehicleResult(auctionID string) (string, error) { //ใบประกาศ ผลประมูล

// 	pdf := gofpdf.New("P", "mm", "A4", "")

// 	GenPDFResult(pdf, auctionID)

// 	path, err := common.UploadPdfToGoogle(pdf, "ใบประกาศผลประมูล", "auction", "fourwd-auction")
// 	if err != nil {
// 		return "", err
// 	}
// 	return path, nil
// }

func GenPDFResultByMonth(month string) (string, error) { //ใบประกาศ ผลประมูล

	auctionlist, _ := prepareAuctionListByMonth(month) // 2024-08

	filepathStr := "images/pdf/"
	// fileextention := ".pdf"
	pdf := gofpdf.New("P", "mm", "A4", "")

	// Loop By Auction
	for _, auction := range auctionlist {

		vehicles, _ := prepareResultByMonth(auction.AuctionID) // auction UUID by month

		pdf.AddPage()
		pdf.AddUTF8Font("Sarabun", "", "fonts/THSarabun.ttf")
		pdf.AddUTF8Font("Sarabun", "B", "fonts/THSarabunBold.ttf")

		headertableReport(pdf, 70, 78)

		header := filepathStr + "Head-2.png"
		pdf.Image(header, 0, 0, 210, 70, false, "", 0, "")

		pdf.SetFont("Sarabun", "B", 24)
		pdf.SetTextColor(196, 16, 22)

		pdf.SetXY(150, 20)
		pdf.SetFillColor(255, 255, 255)
		pdf.SetDrawColor(0, 0, 0)
		pdf.CellFormat(35, 8, "ผลการประมูลรถ", "0", 0, "C", true, 0, "")

		pdf.SetXY(150, 30)
		pdf.SetFillColor(255, 255, 255)
		pdf.SetDrawColor(0, 0, 0)
		pdf.CellFormat(35, 8, "รอบประมูลวันที่", "0", 0, "C", true, 0, "")

		// for _, v := range vehicles {

		// currentTime := time.Now()
		// thaiYear := currentTime.Year() + 543

		// formattedEndDate := v
		// Parse the string to time.Time
		// currentDate, err := time.Parse("2006-01", month)
		// if err != nil {
		// 	fmt.Println("Error parsing date:", err)
		// 	return "", err
		// }
		formattedEndDate := strconv.Itoa(auction.EndDateAuction.Day()) + " " + getThaiMonth(auction.EndDateAuction.Month().String()) + " " + strconv.Itoa(auction.EndDateAuction.Year()+543)
		// formattedStartDateThai := fmt.Sprintf("%s %d", formattedEndDate, thaiYear)
		fmt.Println(formattedEndDate)
		pdf.SetXY(150, 40)
		pdf.SetFillColor(255, 255, 255)
		pdf.SetDrawColor(0, 0, 0)
		pdf.CellFormat(35, 8, formattedEndDate, "0", 0, "C", true, 0, "")

		// }

		pdf.SetTextColor(0, 0, 0)

		page := 1
		// y := 75.0
		perpage := 31
		line := 1
		tableYz := 88.00

		for i, v := range vehicles {
			i++

			if page == 1 {
			} else {
				perpage = 43
			}
			fmt.Printf("index = %d, License = %s", i, v.License)
			fmt.Printf(", page = %d, perpage = %d", page, perpage)
			fmt.Printf(", line = %d", line)
			fmt.Println("========================")

			tableX := 0
			pdf.SetXY(float64(tableX), float64(tableYz))
			pdf.SetFont("Sarabun", "", 12)

			// Calculate the height for the MultiCell for VehicleBrandName
			startX, startY := pdf.GetX(), pdf.GetY()
			pdf.MultiCell(25, 6, v.VehicleBrandName, "", "C", true)
			endY := pdf.GetY()
			brandCellHeight := endY - startY

			// Set row height based on the tallest cell
			rowHeight := brandCellHeight
			if rowHeight < 6 {
				rowHeight = 6
			}

			// Move back to the start of the row
			pdf.SetXY(startX, startY)

			// Draw all cells with adjusted height
			pdf.CellFormat(10, rowHeight, v.VehicleNo, "1", 0, "C", true, 0, "")
			pdf.SetXY(startX+10, startY)
			pdf.MultiCell(25, 6, v.VehicleBrandName, "1", "C", true)
			pdf.SetXY(startX+35, startY)
			pdf.CellFormat(75, rowHeight, v.VehicleModelName+" "+v.VehicleSubModelName, "1", 0, "L", true, 0, "")
			pdf.CellFormat(10, rowHeight, v.YearRegister, "1", 0, "C", true, 0, "")
			pdf.CellFormat(10, rowHeight, v.VehicleGradeID, "1", 0, "C", true, 0, "")
			pdf.CellFormat(15, rowHeight, v.VehicleColorName, "1", 0, "C", true, 0, "")

			pdf.SetFont("Sarabun", "B", 12)
			pdf.CellFormat(30, rowHeight, v.License, "1", 0, "C", true, 0, "")

			closeprice, _ := strconv.ParseFloat(v.ClosePrice, 64)
			imcvat := float64(float64(closeprice) * 1.07)
			formattedVatPrice := common.FloatWithCommas(imcvat, 0)
			formattedCloseprice := common.FloatWithCommas(closeprice, 0)
			pdf.CellFormat(17.5, rowHeight, formattedCloseprice, "1", 0, "R", true, 0, "")
			pdf.SetFont("Sarabun", "", 12)
			pdf.CellFormat(17.5, rowHeight, formattedVatPrice, "1", 0, "R", true, 0, "")

			pdf.Ln(-1)
			tableYz += rowHeight
			line++
			newpage := false

			if i == 31 {
				newpage = true
			} else if i > 31 {
				temp := i - 31
				if temp%perpage == 0 {
					newpage = true
				}
			}
			if newpage {
				line = 1
				page++
				if i < len(vehicles) {
					pdf.AddPage()
					tableYz = 18
					headertableReport(pdf, 0, 8)
				}
			}

			headerdown := filepathStr + "Bottom.png"
			pdf.Image(headerdown, 0, 279, 210, 18, false, "", 0, "")
		}
		pdf.Ln(-1)
	}
	// filedesination := filepathStr + "ใบประกาศผล Result" + fileextention
	// err := pdf.OutputFileAndClose(filedesination)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// }
	// return filedesination, nil
	path, err := common.UploadPdfToGoogle(pdf, "ใบประกาศผลประมูล", "auction", "fourwd-auction")
	if err != nil {
		return "", err
	}
	return path, nil
}

func prepareAuctionListByMonth(month string) ([]VehicleReportAuction, Summary) {

	var auctions []VehicleReportAuction
	common.Database.Raw(`SELECT  DATE(a.start_date) as start_date_auction,DATE(a.end_date) as end_date_auction,av.auction_id AS auction_id
	FROM auction_vehicles AS av
	LEFT JOIN auctions AS a ON av.auction_id = a.id
	LEFT JOIN vehicles AS v ON  v.id = av.vehicle_id
	LEFT JOIN vehicle_grades AS vg ON v.vehicle_grade_id = vg.id
	LEFT JOIN sources AS s ON v.source_id = s.id
	LEFT JOIN ( SELECT DISTINCT auction_id , chassis_no , license , crp , redbook FROM redbooks ) r 
	ON av.auction_id = r.auction_id AND v.chassis_no = r.chassis_no AND v.license = r.license
	LEFT JOIN branches b ON v.branch_id = b.id
	LEFT JOIN vehicle_brands vb ON v.vehicle_brand_id = vb.id
	LEFT JOIN vehicle_models vm ON v.vehicle_model_id = vm.id
	LEFT JOIN vehicle_images AS vi ON av.vehicle_id = vi.vehicle_id AND vi.template_vehicle_image_id = '4b8fe630-a2b6-4470-9e86-3825c169a8f5' AND vi.is_delete = 0
	WHERE DATE_FORMAT(a.actual_end_date, '%Y-%m') = ? AND av.is_win = true
	GROUP BY av.auction_id ORDER BY av.vehicle_no ASC`, month).Scan(&auctions)
	common.Print(month, "month")
	common.Print("vehicle", common.StructToString(auctions))

	summary := new(Summary)
	common.Database.Raw(`
	SELECT MONTH(a.start_date) as start_date_auction,COUNT(*) AS total_car,
    SUM(av.close_price) AS total_price,
    CONCAT_WS(" ","รวมรอบเดือน ",MONTH(a.start_date) )AS auction_name
    FROM
        auction_vehicles av
        LEFT JOIN vehicles v on av.vehicle_id = v.id
        LEFT JOIN auctions a ON av.auction_id = a.id
	where
	DATE_FORMAT(a.actual_end_date, '%Y-%m') = ? AND av.is_win = true  GROUP BY MONTH(a.start_date) , auction_name`, month).Scan(&summary)
	common.Print("year month", month)

	return auctions, *summary
}

func prepareResultByMonth(auctionID string) ([]VehicleReport, Summary) {

	var vehicles []VehicleReport
	common.Database.Raw(`SELECT v.engine_no, v.contract_number as cn,v.vehicle_auction_receipt as receipt, DATE(a.start_date) as start_date_auction,DATE(a.end_date) as end_date_auction, b.text_color,b.background_color,s.name as source_name,vg.color_code as color_code,v.id as vehicle_id, v.license_receive_date,(r.crp-(r.crp*7/100)) AS crp_p_vat,	r.crp as crp,
	v.vehicle_grade_id, v.vehicle_gear_name, v.vehicle_brand_name,vb.name_en as vehicle_brand_name_th,v.vehicle_color_name,
	v.vehicle_model_name, av.open_price,av.close_price as close_price, v.year_manufacturing as years,v.year_register, vm.name_en as vehicle_model_name_th,v.vehicle_sub_model_name,v.license,
	v.license_province_id, v.license_province_name, v.mile AS mile, b.label as branch_label,v.remark,COALESCE(v.vehicle_grade_id, '') AS vehicle_grade_id,
	COALESCE(v.vehicle_grade_value, '') AS vehicle_grade_value,
	COALESCE(av.vehicle_no, 0) AS vehicle_no,
	vi.image_path AS image_preview_path ,av.auction_id AS auction_id
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
						WHERE av.auction_id = ? AND av.is_win = true
							ORDER BY av.vehicle_no ASC`, auctionID).Scan(&vehicles)
	common.Print(auctionID, "auctionID")
	common.Print("vehicle", common.StructToString(vehicles))

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
	  av.auction_id = ? AND av.is_win = true 
	  GROUP BY DATE(a.start_date), a.name`, auctionID).Scan(&summary)
	common.Print(auctionID, "auctionID")

	return vehicles, *summary
}

type VehicleReportAuction struct {
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
	AuctionID          string    `json:"auction_id"`
	StartDateAuction   time.Time `json:"start_date_auction"`
	EndDateAuction     time.Time `json:"end_date_auction"`
}

/*
func GenPDFResultByMonth(month string) (string, error) { //ใบประกาศ ผลประมูล
	vehicles, _ := prepareResultByMonth(month) // 2024-08

	filepathStr := "images/pdf/"
	// fileextention := ".pdf"
	pdf := gofpdf.New("P", "mm", "A4", "")

	pdf.AddPage()
	pdf.AddUTF8Font("Sarabun", "", "fonts/THSarabun.ttf")
	pdf.AddUTF8Font("Sarabun", "B", "fonts/THSarabunBold.ttf")

	headertableReport(pdf, 70, 78)

	header := filepathStr + "Head-2.png"
	pdf.Image(header, 0, 0, 210, 70, false, "", 0, "")

	pdf.SetFont("Sarabun", "B", 24)
	pdf.SetTextColor(196, 16, 22)

	pdf.SetXY(150, 20)
	pdf.SetFillColor(255, 255, 255)
	pdf.SetDrawColor(0, 0, 0)
	pdf.CellFormat(35, 8, "ผลการประมูลรถ", "0", 0, "C", true, 0, "")

	pdf.SetXY(150, 30)
	pdf.SetFillColor(255, 255, 255)
	pdf.SetDrawColor(0, 0, 0)
	pdf.CellFormat(35, 8, "รอบประมูลเดือน", "0", 0, "C", true, 0, "")

	// for _, v := range vehicles {

	// currentTime := time.Now()
	// thaiYear := currentTime.Year() + 543

	// formattedEndDate := v
	// Parse the string to time.Time
	currentDate, err := time.Parse("2006-01", month)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return "", err
	}
	formattedEndDate := getThaiMonth(currentDate.Month().String()) + " " + strconv.Itoa(currentDate.Year()+543)
	// formattedStartDateThai := fmt.Sprintf("%s %d", formattedEndDate, thaiYear)
	fmt.Println(formattedEndDate)
	pdf.SetXY(150, 40)
	pdf.SetFillColor(255, 255, 255)
	pdf.SetDrawColor(0, 0, 0)
	pdf.CellFormat(35, 8, formattedEndDate, "0", 0, "C", true, 0, "")

	// }

	pdf.SetTextColor(0, 0, 0)

	page := 1
	// y := 75.0
	perpage := 31
	line := 1
	tableYz := 88.00

	// for i, v := range vehicles {
	// 	i++

	// 	if page == 1 {
	// 	} else {
	// 		perpage = 28
	// 	}
	// 	fmt.Printf("index = %d, License = %s", i, v.License)
	// 	fmt.Printf(", page = %d, perpage = %d", page, perpage)
	// 	fmt.Printf(", line = %d", line)
	// 	fmt.Println("========================")

	// 	tableX := 0
	// 	pdf.SetXY(float64(tableX), float64(tableYz))
	// 	pdf.SetFont("Sarabun", "", 12)
	// 	pdf.CellFormat(10, 8, v.VehicleNo, "1", 0, "C", true, 0, "")
	// 	pdf.CellFormat(15, 8, v.VehicleBrandName, "1", 0, "C", true, 0, "")
	// 	// pdf.CellFormat(10, 8, strconv.Itoa(counter), "1", 0, "C", true, 0, "")
	// 	pdf.CellFormat(75, 8, v.VehicleModelName+" "+v.VehicleSubModelName, "1", 0, "L", true, 0, "")
	// 	pdf.CellFormat(15, 8, v.YearRegister, "1", 0, "C", true, 0, "")
	// 	pdf.CellFormat(10, 8, v.VehicleGradeID, "1", 0, "C", true, 0, "")

	// 	pdf.CellFormat(15, 8, v.VehicleColorName, "1", 0, "C", true, 0, "")

	// 	pdf.SetFont("Sarabun", "B", 12)
	// 	// openprice, _ := strconv.ParseFloat(v.OpenPrice, 64)
	// 	// formattedOpenPrice := common.FloatWithCommas(openprice, 0)

	// 	pdf.CellFormat(30, 8, v.License, "1", 0, "C", true, 0, "")
	// 	// pdf.CellFormat(20, 8, formattedOpenPrice, "1", 0, "R", true, 0, "")

	// 	closeprice, _ := strconv.ParseFloat(v.ClosePrice, 64)
	// 	imcvat := float64(float64(closeprice) * 1.07)

	// 	formattedVatPrice := common.FloatWithCommas(imcvat, 0)
	// 	formattedCloseprice := common.FloatWithCommas(closeprice, 0)
	// 	// pdf.CellFormat(20, 8, v.ClosePrice, "1", 0, "C", true, 0, "")
	// 	pdf.CellFormat(20, 8, formattedCloseprice, "1", 0, "R", true, 0, "")
	// 	pdf.SetFont("Sarabun", "", 12)

	// 	pdf.CellFormat(20, 8, formattedVatPrice, "1", 0, "R", true, 0, "")

	// 	pdf.Ln(-1)

	// 	tableYz += 8

	// 	line++
	// 	newpage := false

	// 	if i == 23 {
	// 		newpage = true

	// 	} else if i > 23 {
	// 		temp := i - 23
	// 		if temp%perpage == 0 {
	// 			newpage = true
	// 		}
	// 	}
	// 	if newpage {
	// 		line = 1
	// 		page++
	// 		if i < len(vehicles) {
	// 			pdf.AddPage()
	// 			tableYz = 18
	// 			headertableReport(pdf, 0, 8)

	// 		}
	// 	}

	// 	headerdown := filepathStr + "Bottom.png"
	// 	pdf.Image(headerdown, 0, 279, 210, 18, false, "", 0, "")
	// }
	// pdf.Ln(-1)
	for i, v := range vehicles {
		i++

		if page == 1 {
		} else {
			perpage = 43
		}
		fmt.Printf("index = %d, License = %s", i, v.License)
		fmt.Printf(", page = %d, perpage = %d", page, perpage)
		fmt.Printf(", line = %d", line)
		fmt.Println("========================")

		tableX := 0
		pdf.SetXY(float64(tableX), float64(tableYz))
		pdf.SetFont("Sarabun", "", 12)

		// Calculate the height for the MultiCell for VehicleBrandName
		startX, startY := pdf.GetX(), pdf.GetY()
		pdf.MultiCell(25, 6, v.VehicleBrandName, "", "C", true)
		endY := pdf.GetY()
		brandCellHeight := endY - startY

		// Set row height based on the tallest cell
		rowHeight := brandCellHeight
		if rowHeight < 6 {
			rowHeight = 6
		}

		// Move back to the start of the row
		pdf.SetXY(startX, startY)

		// Draw all cells with adjusted height
		pdf.CellFormat(10, rowHeight, v.VehicleNo, "1", 0, "C", true, 0, "")
		pdf.SetXY(startX+10, startY)
		pdf.MultiCell(25, 6, v.VehicleBrandName, "1", "C", true)
		pdf.SetXY(startX+35, startY)
		pdf.CellFormat(75, rowHeight, v.VehicleModelName+" "+v.VehicleSubModelName, "1", 0, "L", true, 0, "")
		pdf.CellFormat(10, rowHeight, v.YearRegister, "1", 0, "C", true, 0, "")
		pdf.CellFormat(10, rowHeight, v.VehicleGradeID, "1", 0, "C", true, 0, "")
		pdf.CellFormat(15, rowHeight, v.VehicleColorName, "1", 0, "C", true, 0, "")

		pdf.SetFont("Sarabun", "B", 12)
		pdf.CellFormat(30, rowHeight, v.License, "1", 0, "C", true, 0, "")

		closeprice, _ := strconv.ParseFloat(v.ClosePrice, 64)
		imcvat := float64(float64(closeprice) * 1.07)
		formattedVatPrice := common.FloatWithCommas(imcvat, 0)
		formattedCloseprice := common.FloatWithCommas(closeprice, 0)
		pdf.CellFormat(17.5, rowHeight, formattedCloseprice, "1", 0, "R", true, 0, "")
		pdf.SetFont("Sarabun", "", 12)
		pdf.CellFormat(17.5, rowHeight, formattedVatPrice, "1", 0, "R", true, 0, "")

		pdf.Ln(-1)
		tableYz += rowHeight
		line++
		newpage := false

		if i == 31 {
			newpage = true
		} else if i > 31 {
			temp := i - 31
			if temp%perpage == 0 {
				newpage = true
			}
		}
		if newpage {
			line = 1
			page++
			if i < len(vehicles) {
				pdf.AddPage()
				tableYz = 18
				headertableReport(pdf, 0, 8)
			}
		}

		headerdown := filepathStr + "Bottom.png"
		pdf.Image(headerdown, 0, 279, 210, 18, false, "", 0, "")
	}
	pdf.Ln(-1)
	// filedesination := filepathStr + "ใบประกาศผล Result" + fileextention
	// err := pdf.OutputFileAndClose(filedesination)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// }
	// return filedesination, nil
	path, err := common.UploadPdfToGoogle(pdf, "ใบประกาศผลประมูล", "auction", "fourwd-auction")
	if err != nil {
		return "", err
	}
	return path, nil
}

func prepareResultByMonth(month string) ([]VehicleReport, Summary) {

	var vehicles []VehicleReport
	common.Database.Raw(`SELECT v.engine_no, v.contract_number as cn,v.vehicle_auction_receipt as receipt, DATE(a.start_date) as start_date_auction,DATE(a.end_date) as end_date_auction, b.text_color,b.background_color,s.name as source_name,vg.color_code as color_code,v.id as vehicle_id, v.license_receive_date,(r.crp-(r.crp*7/100)) AS crp_p_vat,	r.crp as crp,
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
						WHERE DATE_FORMAT(a.actual_end_date, '%Y-%m') = ? AND av.is_win = true
							ORDER BY av.vehicle_no ASC`, month).Scan(&vehicles)
	common.Print("year month", month)
	common.Print("vehicle", common.StructToString(vehicles))

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
	  DATE_FORMAT(a.actual_end_date, '%Y-%m') = ? AND av.is_win = true
	  GROUP BY DATE(a.start_date), a.name`, month).Scan(&summary)
	common.Print("year month", month)

	return vehicles, *summary
}

// unused fix think is work

func prepareVehicleResultByMonth(month string, vehicleID string) ([]VehicleReport, Summary) {

	var vehicles []VehicleReport
	if vehicleID == "" {
		common.Database.Raw(`SELECT v.engine_no,v.chassis_no,v.contract_number as cn,v.vehicle_auction_receipt as receipt, DATE(a.start_date) as start_date_auction,DATE(a.end_date) as end_date_auction, b.text_color,b.background_color,s.name as source_name,vg.color_code as color_code,v.id as vehicle_id, v.license_receive_date,(r.crp-(r.crp*7/100)) AS crp_p_vat,	r.crp as crp,
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
						WHERE DATE_FORMAT(a.actual_end_date, '%Y-%m') = ? AND av.is_win = true
							ORDER BY av.vehicle_no ASC`, month).Scan(&vehicles)
	} else {
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
						WHERE DATE_FORMAT(a.actual_end_date, '%Y-%m') = ? AND av.vehicle_id = ? AND av.is_win = true
							ORDER BY av.vehicle_no ASC`, month, vehicleID).Scan(&vehicles)
	}
	common.Print("year month", month)
	common.Print("vehicle", common.StructToString(vehicles))

	summary := new(Summary)
	common.Database.Raw(`
	SELECT MONTH(a.start_date) as start_date_auction,COUNT(*) AS total_car,
    SUM(av.close_price) AS total_price,
    CONCAT_WS(" ","รวมรอบเดือน ",MONTH(a.start_date) )AS auction_name
    FROM
        auction_vehicles av
        LEFT JOIN vehicles v on av.vehicle_id = v.id
        LEFT JOIN auctions a ON av.auction_id = a.id
	where
	DATE_FORMAT(a.actual_end_date, '%Y-%m') = ? AND av.is_win = true  GROUP BY MONTH(a.start_date) , auction_name`, month).Scan(&summary)
	common.Print("year month", month)

	return vehicles, *summary
}
*/