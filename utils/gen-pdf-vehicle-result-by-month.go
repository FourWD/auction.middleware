package utils

import (
	"fmt"
	"strconv"
	"time"

	"github.com/FourWD/middleware/common"
	"github.com/jung-kurt/gofpdf"
)

func GenPDFResultByMonth(month string) (string, error) { //ใบประกาศ ผลประมูล

	auctionlist, _ := prepareAuctionListByMonth(month) // 2024-08

	filepathStr := "images/pdf/"
	// fileextention := ".pdf"
	pdf := gofpdf.New("P", "mm", "A4", "")

	pdf.AddUTF8Font("Sarabun", "", "fonts/THSarabun.ttf")
	pdf.AddUTF8Font("Sarabun", "B", "fonts/THSarabunBold.ttf")

	headerAdded := false
	page := 1
	line := 1
	tableYz := 88.00
	perPageFirst := 31
	perPageSubsequent := 43
	counter := 1

	for _, auction := range auctionlist {
		vehicles, _ := prepareResultByMonth(auction.AuctionID) // auction UUID by month

		for i, v := range vehicles {
			if !headerAdded {
				// Add the header only once on the first page
				pdf.AddPage()

				header := filepathStr + "Head-2.png"
				pdf.Image(header, 0, 0, 210, 70, false, "", 0, "")

				pdf.SetFont("Sarabun", "B", 24)
				pdf.SetTextColor(196, 16, 22)

				pdf.SetXY(150, 20)
				pdf.SetFillColor(255, 255, 255)
				pdf.SetDrawColor(0, 0, 0)
				pdf.CellFormat(35, 8, "ผลการประมูลรถ", "0", 0, "C", true, 0, "")

				pdf.SetXY(135, 30)
				pdf.SetFillColor(255, 255, 255)
				pdf.SetFont("Sarabun", "B", 18)
				pdf.SetDrawColor(0, 0, 0)
				pdf.CellFormat(30, 8, getDateAuction(auction.StartDateAuction), "0", 0, "L", true, 0, "")

				pdf.SetTextColor(0, 0, 0)

				headerAdded = true
				headertableReport(pdf, 70, 78)
			}

			if page > 1 && line == 1 {
				pdf.AddPage()
				tableYz = 18
				headertableReport(pdf, 0, 8)
			}

			fmt.Printf("index = %d, License = %s", i, v.License)
			fmt.Printf(", page = %d, line = %d", page, line)
			fmt.Println("========================")

			tableX := 0
			pdf.SetXY(float64(tableX), float64(tableYz))
			pdf.SetFont("Sarabun", "", 12)

			startX, startY := pdf.GetX(), pdf.GetY()
			pdf.MultiCell(25, 6, v.VehicleBrandName, "", "C", true)
			endY := pdf.GetY()
			brandCellHeight := endY - startY

			rowHeight := brandCellHeight
			if rowHeight < 6 {
				rowHeight = 6
			}
			pdf.SetXY(startX, startY)
			pdf.CellFormat(10, rowHeight, fmt.Sprintf("%d", counter), "1", 0, "C", true, 0, "")
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
			imcvat := float64(closeprice * 1.07)
			formattedVatPrice := common.FloatWithCommas(imcvat, 0)
			formattedCloseprice := common.FloatWithCommas(closeprice, 0)
			pdf.CellFormat(17.5, rowHeight, formattedCloseprice, "1", 0, "R", true, 0, "")
			pdf.SetFont("Sarabun", "", 12)
			pdf.CellFormat(17.5, rowHeight, formattedVatPrice, "1", 0, "R", true, 0, "")

			pdf.Ln(-1)
			tableYz += rowHeight
			line++
			counter++

			maxLines := perPageFirst
			if page > 1 {
				maxLines = perPageSubsequent
			}

			if line > maxLines {
				page++
				line = 1
				tableYz = 39
			}

			headerdown := filepathStr + "Bottom.png"
			pdf.Image(headerdown, 0, 279, 210, 18, false, "", 0, "")
		}
	}

	path, err := common.UploadPdfToGoogle(pdf, "ใบประกาศผลประมูล", "auction", "fourwd-auction")
	if err != nil {
		return "", err
	}
	return path, nil
}

func prepareAuctionListByMonth(month string) ([]AuctionListResult, Summary) {
	var auctions []AuctionListResult
	common.Database.Raw(`
	SELECT 
		start_date_auction,
		end_date_auction,
		auction_id,
		auction_name
	FROM 
		(SELECT 
			DATE(a.start_date) AS start_date_auction,
			DATE(a.end_date) AS end_date_auction,
			av.auction_id AS auction_id,
			r.name AS auction_name,
			v.vehicle_brand_name
		FROM 
			auction_vehicles AS av
		LEFT JOIN 
			auctions AS a ON av.auction_id = a.id
		LEFT JOIN 
			vehicles AS v ON v.id = av.vehicle_id
		LEFT JOIN 
			vehicle_grades AS vg ON v.vehicle_grade_id = vg.id
		LEFT JOIN 
			sources AS s ON v.source_id = s.id
		LEFT JOIN 
			(SELECT DISTINCT auction_id, chassis_no, license, crp, redbook FROM redbooks) r 
			ON av.auction_id = r.auction_id AND v.chassis_no = r.chassis_no AND v.license = r.license
		LEFT JOIN 
			branches b ON v.branch_id = b.id
		LEFT JOIN 
			vehicle_brands vb ON v.vehicle_brand_id = vb.id
		LEFT JOIN 
			vehicle_models vm ON v.vehicle_model_id = vm.id
		LEFT JOIN 
			vehicle_images AS vi ON av.vehicle_id = vi.vehicle_id AND vi.template_vehicle_image_id = '4b8fe630-a2b6-4470-9e86-3825c169a8f5' AND vi.is_delete = 0
		LEFT JOIN 
			rounds AS r ON a.round_id = r.id
		WHERE 
			DATE_FORMAT(a.actual_end_date, '%Y-%m') = ? AND av.is_win = true
		GROUP BY 
			DATE(a.start_date), DATE(a.end_date), av.auction_id, r.name, v.vehicle_brand_name) AS sub
	ORDER BY 
		vehicle_brand_name = "ISUZU" ASC`, month).Debug().Scan(&auctions)

	summary := new(Summary)
	common.Database.Raw(`
	SELECT 
		MONTH(a.start_date) as start_date_auction,
		COUNT(*) AS total_car,
		SUM(av.close_price) AS total_price,
		CONCAT_WS(" ", "รวมรอบเดือน", MONTH(a.start_date)) AS auction_name
	FROM
		auction_vehicles av
	LEFT JOIN 
		vehicles v ON av.vehicle_id = v.id
	LEFT JOIN 
		auctions a ON av.auction_id = a.id
	WHERE 
		DATE_FORMAT(a.actual_end_date, '%Y-%m') = ? AND av.is_win = true
	GROUP BY 
		MONTH(a.start_date), auction_name`, month).Debug().Scan(&summary)

	return auctions, *summary
}

func prepareResultByMonth(auctionID string) ([]Vehicle, Summary) {

	var vehicles []Vehicle
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
							ORDER BY av.vehicle_no ASC`, auctionID).Debug().Scan(&vehicles)

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

	return vehicles, *summary
}

func getDateAuction(date time.Time) string {
	return fmt.Sprintf("รอบการประมูล เดือน %s %d", getThaiMonthDate(date.Month()), date.Year()+543)

}

func getThaiMonthDate(month time.Month) string {
	months := []string{
		"มกราคม", "กุมภาพันธ์", "มีนาคม", "เมษายน", "พฤษภาคม", "มิถุนายน",
		"กรกฎาคม", "สิงหาคม", "กันยายน", "ตุลาคม", "พฤศจิกายน", "ธันวาคม",
	}
	if int(month) >= 1 && int(month) <= 12 {
		return months[month-1]
	}
	return ""
}

type AuctionListResult struct {
	AuctionID        string    `json:"auction_id"`
	AuctionName      string    `json:"auction_name"`
	StartDateAuction time.Time `json:"start_date_auction"`
	EndDateAuction   time.Time `json:"end_date_auction"`
}

// func prepareAuctionListByMonth(month string) ([]AuctionListResult, Summary) {

// 	var auctions []AuctionListResult
// 	common.Database.Raw(`SELECT  DATE(a.start_date) as start_date_auction,DATE(a.end_date) as end_date_auction,
// 	av.auction_id AS auction_id,r.name AS auction_name
// 	FROM auction_vehicles AS av
// 	LEFT JOIN auctions AS a ON av.auction_id = a.id
// 	LEFT JOIN vehicles AS v ON  v.id = av.vehicle_id
// 	LEFT JOIN vehicle_grades AS vg ON v.vehicle_grade_id = vg.id
// 	LEFT JOIN sources AS s ON v.source_id = s.id
// 	LEFT JOIN ( SELECT DISTINCT auction_id , chassis_no , license , crp , redbook FROM redbooks ) r
// 	ON av.auction_id = r.auction_id AND v.chassis_no = r.chassis_no AND v.license = r.license
// 	LEFT JOIN branches b ON v.branch_id = b.id
// 	LEFT JOIN vehicle_brands vb ON v.vehicle_brand_id = vb.id
// 	LEFT JOIN vehicle_models vm ON v.vehicle_model_id = vm.id
// 	LEFT JOIN vehicle_images AS vi ON av.vehicle_id = vi.vehicle_id AND vi.template_vehicle_image_id = '4b8fe630-a2b6-4470-9e86-3825c169a8f5' AND vi.is_delete = 0
// 	LEFT JOIN rounds AS r ON a.round_id = r.id
// 	WHERE DATE_FORMAT(a.actual_end_date, '%Y-%m') = ? AND av.is_win = true
// 	GROUP BY av.auction_id ORDER BY end_date_auction ASC`, month).Scan(&auctions)
// 	// common.Print(month, "month")
// 	// common.Print("vehicle", common.StructToString(auctions))

// 	summary := new(Summary)
// 	common.Database.Raw(`
// 	SELECT MONTH(a.start_date) as start_date_auction,COUNT(*) AS total_car,
//     SUM(av.close_price) AS total_price,
//     CONCAT_WS(" ","รวมรอบเดือน ",MONTH(a.start_date) )AS auction_name
//     FROM
//         auction_vehicles av
//         LEFT JOIN vehicles v on av.vehicle_id = v.id
//         LEFT JOIN auctions a ON av.auction_id = a.id
// 	where
// 	DATE_FORMAT(a.actual_end_date, '%Y-%m') = ? AND av.is_win = true  GROUP BY MONTH(a.start_date) , auction_name`, month).Scan(&summary)
// 	// common.Print("year month", month)

// 	return auctions, *summary
// }

// func prepareResultByMonth(auctionID string) ([]VehicleReport, Summary) {

// 	var vehicles []VehicleReport
// 	common.Database.Raw(`SELECT v.engine_no, v.contract_number as cn,v.vehicle_auction_receipt as receipt, DATE(a.start_date) as start_date_auction,DATE(a.end_date) as end_date_auction, b.text_color,b.background_color,s.name as source_name,vg.color_code as color_code,v.id as vehicle_id, v.license_receive_date,(r.crp-(r.crp*7/100)) AS crp_p_vat,	r.crp as crp,
// 	v.vehicle_grade_id, v.vehicle_gear_name, v.vehicle_brand_name,vb.name_en as vehicle_brand_name_th,v.vehicle_color_name,
// 	v.vehicle_model_name, av.open_price,av.close_price as close_price, v.year_manufacturing as years,v.year_register, vm.name_en as vehicle_model_name_th,v.vehicle_sub_model_name,v.license,
// 	v.license_province_id, v.license_province_name, v.mile AS mile, b.label as branch_label,v.remark,COALESCE(v.vehicle_grade_id, '') AS vehicle_grade_id,
// 	COALESCE(v.vehicle_grade_value, '') AS vehicle_grade_value,
// 	COALESCE(av.vehicle_no, 0) AS vehicle_no,
// 	vi.image_path AS image_preview_path ,av.auction_id AS auction_id
// 	FROM auction_vehicles AS av
// 	LEFT JOIN auctions AS a ON av.auction_id = a.id
// LEFT JOIN vehicles AS v ON  v.id = av.vehicle_id
// LEFT JOIN vehicle_grades AS vg ON v.vehicle_grade_id = vg.id
// LEFT JOIN sources AS s ON v.source_id = s.id
// LEFT JOIN (
// 	SELECT DISTINCT auction_id , chassis_no , license , crp , redbook
// FROM redbooks
//  ) r ON av.auction_id = r.auction_id AND v.chassis_no = r.chassis_no AND v.license = r.license
// LEFT JOIN branches b ON v.branch_id = b.id
// LEFT JOIN vehicle_brands vb ON v.vehicle_brand_id = vb.id
// LEFT JOIN vehicle_models vm ON v.vehicle_model_id = vm.id
// LEFT JOIN vehicle_images AS vi ON av.vehicle_id = vi.vehicle_id AND vi.template_vehicle_image_id = '4b8fe630-a2b6-4470-9e86-3825c169a8f5' AND vi.is_delete = 0
// 						WHERE av.auction_id = ? AND av.is_win = true
// 							ORDER BY av.vehicle_no ASC`, auctionID).Scan(&vehicles)
// 	// common.Print(auctionID, "auctionID")
// 	// common.Print("vehicle", common.StructToString(vehicles))

// 	summary := new(Summary)
// 	common.Database.Raw(`
// 	SELECT DATE(a.start_date) as start_date_auction,COUNT(*) AS total_car,
//     SUM(av.close_price) AS total_price,
//     MAX(a.name) AS auction_name
//     FROM
//         auction_vehicles av
//         LEFT JOIN vehicles v on av.vehicle_id = v.id
//         LEFT JOIN auctions a ON av.auction_id = a.id
// 	where
// 	  av.auction_id = ? AND av.is_win = true
// 	  GROUP BY DATE(a.start_date), a.name`, auctionID).Scan(&summary)
// 	// common.Print(auctionID, "auctionID")

// 	return vehicles, *summary
// }
