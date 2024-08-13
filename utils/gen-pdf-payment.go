package utils

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	model "github.com/FourWD/auction.middleware/model"

	"github.com/FourWD/middleware/common"
	"github.com/jung-kurt/gofpdf"
	"github.com/jung-kurt/gofpdf/contrib/httpimg"
)

func DateToString(date time.Time) string {
	return date.Format("2006-01-02")
}

// return "https://storage.googleapis.com/fourwd-auction/auction/2023/2.pdf", nil //
func GenPDFPayment(auctionID string, userID string) (string, string, error) {
	user, vehicles, summary := preparePayment(auctionID, userID)

	pdf := gofpdf.New("P", "mm", "A4", "")

	genPaymentSummary(pdf, user, vehicles, summary)
	vehicleImage := prepareDataVehicle(auctionID, userID)
	i := 0
	for _, vehicle := range vehicles {
		genByVehicle(pdf, user, vehicle, summary)
		genVehicleImages(pdf, vehicleImage[i])
		i++
	}

	var fileName string
	if user.UserTypeID == "01" {
		fileName = fmt.Sprintf("%s_%s_%s_%s", user.UserFirstname, user.UserLastname, summary.RoundName, DateToString(vehicles[0].EndDate))
	} else {
		companyName := strings.ReplaceAll(user.CompanyName, " ", "_")
		fileName = fmt.Sprintf("%s_%s_%s", companyName, summary.RoundName, DateToString(vehicles[0].EndDate))
	}

	path, err := common.UploadPdfToGoogle(pdf, fileName, "auction", "fourwd-auction")
	if err != nil {
		return "", "", err
	}
	return path, fileName, nil
}

func genByVehicle(pdf *gofpdf.Fpdf, user UserSummary, vehicle VehicleSummary, summary Summary) {
	filepathStr := "images/pdf/"

	// Add a page
	pdf.AddPage()

	header(pdf)

	pdf.AddUTF8Font("Sarabun", "", "fonts/THSarabun.ttf")
	pdf.AddUTF8Font("Sarabun", "B", "fonts/THSarabunBold.ttf")

	// Set font for the entire document
	pdf.SetFont("Sarabun", "B", 30)

	pdf.Text(125, 30, "ใบแจ้งผู้ประมูลราคาสูงสุด")
	pdf.Ln(5)

	if user.UserTypeID == "01" {
		pdf.SetFont("Sarabun", "B", 14)
		pdf.Text(10, 30, "คุณ "+user.UserFirstname+" "+user.UserLastname)
		pdf.Ln(8)
	} else {
		pdf.SetFont("Sarabun", "B", 14)
		pdf.Text(10, 30, user.CompanyName)
		pdf.Ln(8)
	}

	if user.TaxProvinceName != "" || user.TaxDistrictName != "" {

		subname := "แขวง "
		districname := "เขต "
		if user.TaxProvinceID != "10" {
			subname = "ตำบล "
			districname = "อำเภอ "
		}
		pdf.SetFont("Sarabun", "", 10)
		pdf.Text(10, 35, user.TaxAddress+" "+user.TaxStreet+" "+subname+user.TaxSubDistrictName)
		pdf.Ln(4)
		pdf.Text(10, 40, districname+user.TaxDistrictName+" "+"จังหวัด "+user.TaxProvinceName+" "+"รหัสไปรษณีย์  "+user.TaxPostCode)
		pdf.Ln(4)
		pdf.Text(10, 45, "โทร:"+user.Mobile+" | "+"อีเมล:"+user.Email)
		pdf.Ln(4)
		pdf.SetTextColor(0, 0, 0)

	} else {
		subnamev2 := "แขวง "
		districnamev2 := "เขต "
		if user.TaxProvinceID != "10" {
			subnamev2 = "ตำบล "
			districnamev2 = "อำเภอ "
		}

		pdf.SetFont("Sarabun", "", 10)
		pdf.Text(10, 35, user.Address+" "+user.Street+" "+subnamev2+user.SubDistrictName)
		pdf.Ln(4)
		pdf.Text(10, 40, districnamev2+user.DistrictName+" "+"จังหวัด "+user.ProvinceName+" "+"รหัสไปรษณีย์  "+user.TaxPostCode)
		pdf.Ln(4)
		pdf.Text(10, 45, "โทร: "+user.Mobile+" | "+"อีเมล: "+user.Email)
		pdf.Ln(4)
		pdf.SetTextColor(0, 0, 0)
	}

	currentDate := vehicle.EndDate
	formattedDate := strconv.Itoa(currentDate.Day()) + " " + getThaiMonth(currentDate.Month().String()) + " พ.ศ." + strconv.Itoa(currentDate.Year()+543)
	pdf.SetFont("Sarabun", "B", 14)
	pdf.Text(152, 38, "วันที่")
	pdf.SetFont("Sarabun", "", 14)
	pdf.Text(162, 38, formattedDate)
	pdf.Ln(10)

	pdf.SetFont("Sarabun", "B", 14)
	pdf.Text(141, 45, "เลขที่เอกสาร ")
	pdf.SetFont("Sarabun", "", 14)
	pdf.Text(162, 45, user.Code+"-"+vehicle.CodeAuction+"/"+vehicle.VehicleNo)
	pdf.Ln(15)

	// Table Header
	pdf.SetFont("Sarabun", "", 10)
	pdf.Text(11, 52, "ชื่อ-นามสกุล ผู้ประมูล")
	pdf.SetFont("Sarabun", "B", 10)
	pdf.Text(11, 56, "คุณ "+user.UserFirstname+" "+user.UserLastname)

	pdf.SetFont("Sarabun", "", 10)
	pdf.SetXY(155, 52)
	pdf.SetFillColor(255, 255, 255)
	pdf.SetDrawColor(0, 0, 0)
	pdf.SetTextColor(0, 0, 0)
	pdf.MultiCell(40, 6, vehicle.AuctionName, "0", "R", true)
	// pdf.Text(140, 52, "รอบการประมูลของ "+vehicle.AuctionName)
	pdf.SetFont("Sarabun", "B", 10)

	pdf.SetLineWidth(0.2)
	pdf.SetDrawColor(0, 0, 0)
	pdf.Line(10, 48, 200, 48)
	pdf.Ln(1)

	bbg := filepathStr + "bbg.png"
	pdf.Image(bbg, 10, 65, 190, 100, false, "", 0, "")
	pdf.Ln(30)

	// // Data rows
	// pdf.SetTextColor(0, 0, 0)
	// pdf.SetFont("Sarabun", "B", 16)
	//+33
	pdf.SetFont("Sarabun", "B", 12)
	pdf.Text(13, 85, vehicle.VehicleBrandName)
	pdf.Text(46, 85, vehicle.VehicleModelName+" "+vehicle.VehicleSubModelName)
	pdf.Text(146, 85, vehicle.YearRegister)
	pdf.Text(166, 85, vehicle.License)
	pdf.Ln(1)
	pdf.Text(166, 90, vehicle.LicenseProvinceName)
	pdf.Ln(-1)
	pdf.Text(13, 107, vehicle.VehicleColorName)

	mile, _ := strconv.ParseFloat(vehicle.Mile, 64)
	mileComma := common.FloatWithCommas(mile, 0)

	pdf.Text(46, 107, mileComma)

	pdf.Text(81, 107, vehicle.EngineSize+" ซีซี")
	pdf.Text(113, 107, vehicle.EngineNo)
	pdf.Text(146, 107, vehicle.ChassisNo)
	pdf.Ln(-1)
	pdf.SetFont("Sarabun", "B", 9)
	// pdf.Text(28, 120, "ลานจอดรถ เพชรบุรี ระยะเวลาในการโอนเอกสาร 60 วัน หากเอกสารหมดอายุ มีค่าใช้จ่ายการออกเอกสารใหม่ 500 บาท รถวันนี้ติดค่าใช้จ่าย M Flow")
	pdf.SetXY(25, 117)
	pdf.SetFillColor(255, 255, 255)
	pdf.SetDrawColor(0, 0, 0)
	pdf.SetTextColor(0, 0, 0)

	pdf.MultiCell(150, 3, "ลานจอด "+" "+vehicle.BranchLabel+" "+vehicle.Remark, "0", "L", true)
	imcvat, _ := strconv.ParseFloat(vehicle.ClosePrice, 64)
	vatPrice := float64(float64(imcvat) * 1.07)

	vat := vatPrice - imcvat

	formattedVatPrice := common.FloatWithCommas(vatPrice, 2)
	formattedVatPriceTotal := common.FloatWithCommas(vatPrice, 2)

	formattedTotal := common.FloatWithCommas(imcvat, 2)
	vattotal := common.FloatWithCommas(vat, 2)

	// Table down

	// Row 1
	detailtextcar := filepathStr + "detail-text-car.png"
	pdf.Image(detailtextcar, 10, 170, 180, 15, false, "", 0, "")
	pdf.Ln(-1)

	pdf.Ln(40)
	downscan := filepathStr + "detail-price-new2.png"
	pdf.Image(downscan, 10, 190, 190, 25, false, "", 0, "")
	pdf.Ln(-1)

	qr := summary.BankQr
	httpimg.Register(pdf, summary.BankQr, "")
	pdf.Image(qr, 178, 195, 18, 18, false, "", 0, "")

	pdf.SetFont("Sarabun", "", 10)
	xTextBankName := 11
	yTextBankName := 196
	pdf.SetXY(float64(xTextBankName), float64(yTextBankName))
	bank := summary.BankName
	pdf.CellFormat(0, 5, bank, "", 0, "L", false, 0, "")
	pdf.SetTextColor(0, 0, 0)

	pdf.SetFont("Sarabun", "B", 24)
	xTextBankAccountNo := 65
	yTextBankAccountNo := 202
	pdf.SetXY(float64(xTextBankAccountNo), float64(yTextBankAccountNo))
	bankAccountNo := summary.BankAccountNo
	pdf.CellFormat(45, 5, bankAccountNo, "", 0, "C", false, 0, "")
	pdf.SetTextColor(0, 0, 0)

	pdf.SetFont("Sarabun", "B", 12)
	xTextBankAccount := 15
	yTextBankAccount := 204
	pdf.SetXY(float64(xTextBankAccount), float64(yTextBankAccount))
	bankAccount := summary.BankAccount
	pdf.CellFormat(50, 5, bankAccount, "", 0, "C", false, 0, "")
	pdf.SetTextColor(0, 0, 0)

	pdf.SetFont("Sarabun", "B", 18)
	xTextTotalPrice := 185
	yTextTotalPrice := 136
	pdf.SetXY(float64(xTextTotalPrice), float64(yTextTotalPrice))
	texttotalprice := formattedTotal
	pdf.CellFormat(10, 5, texttotalprice, "", 0, "R", false, 0, "")
	pdf.Ln(20)

	xTextVatPrice := 185
	yTextVatPrice := 146
	pdf.SetXY(float64(xTextVatPrice), float64(yTextVatPrice))
	textvatprice := vattotal
	pdf.CellFormat(10, 5, textvatprice, "", 0, "R", false, 0, "")
	pdf.Ln(1)

	xTextTotalVatPrice := 185
	yTextTotalVatPrice := 155
	pdf.SetXY(float64(xTextTotalVatPrice), float64(yTextTotalVatPrice))
	texttotalvatprice := formattedVatPrice
	pdf.CellFormat(10, 5, texttotalvatprice, "", 0, "R", false, 0, "")
	pdf.Ln(20)

	pdf.SetTextColor(255, 0, 0)
	pdf.SetFont("Sarabun", "B", 28)
	xTextTotal := 88
	yTextTotal := 202
	pdf.SetXY(float64(xTextTotal), float64(yTextTotal))
	texttotal := formattedVatPriceTotal
	pdf.CellFormat(0, 5, texttotal, "", 0, "C", false, 0, "")
	pdf.SetTextColor(0, 0, 0)
	pdf.Ln(20)

	pdf.SetFont("Sarabun", "", 9)
	xText := 88
	yText := 209
	pdf.SetXY(float64(xText), float64(yText))
	text := common.ConvertFloatToThaiText(vatPrice)
	pdf.CellFormat(0, 5, text, "", 0, "C", false, 0, "")

	textdown := filepathStr + "text-down.png"
	pdf.Image(textdown, 10, 220, 180, 60, false, "", 0, "")
	pdf.Ln(-1)
}

func GenPDFQRCode(url string) string {
	return "scan.png"
}

func genPDFLinkDownload(auctionID string, vehicleID string) (string, string) {
	url := "www.omakasecar.com/download/" + vehicleID
	qr := GenPDFQRCode(vehicleID)

	return url, qr
}
func genPaymentSummary(pdf *gofpdf.Fpdf, user UserSummary, vehicles []VehicleSummary, summary Summary) {

	filepathStr := "images/pdf/"

	// Add a page
	pdf.AddPage()
	header(pdf)

	pdf.SetLineWidth(0.2) // Set line width
	pdf.Line(10, 48, 200, 48)
	pdf.Ln(5)
	pdf.SetLineWidth(0) // Set line width

	// Set font for the entire document
	pdf.AddUTF8Font("Sarabun", "", "fonts/THSarabun.ttf")
	pdf.AddUTF8Font("Sarabun", "B", "fonts/THSarabunBold.ttf")

	// Set font for the entire document
	pdf.SetFont("Sarabun", "B", 30)

	pdf.Text(125, 30, "ใบแจ้งผู้ประมูลราคาสูงสุด")
	pdf.Ln(5)
	// common.Print("sq", user.UserTypeID)
	if user.UserTypeID == "01" {
		pdf.SetFont("Sarabun", "B", 14)
		pdf.Text(10, 30, "คุณ "+user.UserFirstname+" "+user.UserLastname)
		pdf.Ln(8)
	} else {
		pdf.SetFont("Sarabun", "B", 14)
		pdf.Text(10, 30, user.CompanyName)
		pdf.Ln(8)
	}

	if user.IsAddressTax == 0 {

		subname := " แขวง "
		districname := "เขต "
		if user.TaxProvinceID != "10" {
			subname = " ตำบล "
			districname = "อำเภอ "
		}

		taxStreet := ""
		if user.TaxStreet != "" {
			taxStreet = " ถนน " + user.TaxStreet
		}

		taxAddress := ""
		if user.TaxAddress != "" {
			taxAddress = user.TaxAddress
		}

		taxSub := ""
		if user.TaxSubDistrictName != "" {
			taxSub = user.TaxSubDistrictName
		}

		taxDis := ""
		if user.TaxDistrictName != "" {
			taxDis = user.TaxDistrictName
		}

		taxProvince := ""
		if user.TaxProvinceName != "" {
			taxProvince = "จังหวัด " + user.TaxProvinceName
		}

		taxPost := ""
		if user.TaxPostCode != "" {
			taxPost = " รหัสไปรษณีย์ " + user.TaxPostCode
		}

		mobile := ""
		if user.Mobile != "" {
			mobile = "โทร: " + user.Mobile
		}

		email := ""
		if user.Email != "" {
			email = "อีเมล: " + user.Email
		}

		pdf.SetFont("Sarabun", "", 10)
		// if taxSub == "" {
		pdf.Text(10, 35, taxAddress+taxStreet+subname+taxSub)
		pdf.Ln(4)
		// }
		pdf.Text(10, 40, districname+taxDis+" "+taxProvince+taxPost)
		pdf.Ln(4)
		pdf.Text(10, 45, mobile+" | "+email)
		pdf.Ln(4)
		pdf.SetTextColor(0, 0, 0)

	} else {

		subnamev2 := " แขวง "
		districnamev2 := "เขต "
		if user.ProvinceID != "10" && user.ProvinceID != "" {
			subnamev2 = " ตำบล "
			districnamev2 = "อำเภอ "
		}

		address := ""
		if user.Address != "" {
			address = user.Address
		}

		street := ""
		if user.TaxStreet != "" {
			street = " ถนน " + user.Street
		}

		subdis := ""
		if user.SubDistrictName != "" {
			subdis = user.SubDistrictName
		}

		dis := ""
		if user.DistrictName != "" {
			dis = user.DistrictName
		}

		province := ""
		if user.ProvinceName != "" {
			province = "จังหวัด " + user.ProvinceName
		}

		post := ""
		if user.Post != "" {
			post = "รหัสไปรษณีย์ " + user.Post
		}

		mobile := ""
		if user.Mobile != "" {
			mobile = "โทร: " + user.Mobile
		}

		email := ""
		if user.Email != "" {
			email = "อีเมล: " + user.Email
		}

		pdf.SetFont("Sarabun", "", 10)
		pdf.Text(10, 35, address+street+subnamev2+subdis)
		pdf.Ln(4)
		pdf.Text(10, 40, districnamev2+dis+" "+province+" "+post)
		pdf.Ln(4)
		pdf.Text(10, 45, mobile+" | "+email)
		pdf.Ln(4)
		pdf.SetTextColor(0, 0, 0)
	}

	for _, vehicle := range vehicles {

		currentDate := vehicle.EndDate
		formattedDate := strconv.Itoa(currentDate.Day()) + " " + getThaiMonth(currentDate.Month().String()) + " พ.ศ." + strconv.Itoa(currentDate.Year()+543)
		pdf.SetFont("Sarabun", "B", 14)
		pdf.Text(156, 38, "วันที่")
		pdf.SetFont("Sarabun", "", 14)
		pdf.Text(165, 38, formattedDate)
		pdf.Ln(10)

		pdf.SetFont("Sarabun", "B", 14)
		pdf.Text(145, 45, "เลขที่เอกสาร ")
		pdf.SetFont("Sarabun", "", 14)
		pdf.Text(165, 45, user.Code+"-"+vehicle.CodeAuction)
		pdf.Ln(5)
		break

	}

	// pdf.SetFont("Sarabun", "B", 14)
	// pdf.Text(10, 30, "บริษัท "+user.CompanyName)
	// pdf.Ln(4)

	pdf.SetFont("Sarabun", "", 16)
	pdf.Text(10, 55, summary.AuctionName)

	// pdf.SetFont("Sarabun", "", 16)
	// pdf.Text(10, 55, "รอบการประมูลของ"+summary.RoundName+" | "+summary.AuctionName)

	pdf.SetFont("Sarabun", "B", 18)
	pdf.Text(175, 55, "จำนวน "+strconv.Itoa(summary.TotalCar)+" คัน")
	pdf.Ln(8)

	// Define table headers
	pdf.SetFillColor(237, 28, 36)
	pdf.SetTextColor(255, 255, 255)
	pdf.SetFont("Sarabun", "B", 12)
	pdf.SetDrawColor(255, 255, 255) // Set border color to white
	pdf.CellFormat(5, 8, "#", "", 0, "C", true, 0, "")
	pdf.CellFormat(85, 8, "รายละเอียดรถยนต์", "", 0, "", true, 0, "")
	pdf.CellFormat(8, 8, "เกรด", "", 0, "C", true, 0, "")
	pdf.CellFormat(20, 8, "ลานจอด", "", 0, "C", true, 0, "")
	pdf.CellFormat(18, 6, "ราคาสูงสุด", "", 0, "C", true, 0, "")
	pdf.CellFormat(18, 6, "ราคารถ 10%", "", 0, "C", true, 0, "")
	pdf.CellFormat(18, 6, "ราคารถ 90%", "", 0, "C", true, 0, "")
	pdf.SetFillColor(196, 16, 22)
	pdf.CellFormat(18, 6, "รวมค่ารถ", "", 0, "C", true, 0, "")
	pdf.SetFillColor(0, 0, 0)

	pdf.Ln(-1)
	pdf.SetFillColor(237, 28, 36)

	// Add text below the table headers
	pdf.SetFont("Sarabun", "B", 8) // Set smaller font size
	pdf.CellFormat(5, 5, "", "", 0, "C", true, 0, "")
	pdf.CellFormat(85, 5, "", "", 0, "C", true, 0, "")
	pdf.CellFormat(8, 5, "", "", 0, "C", true, 0, "")
	pdf.CellFormat(20, 5, "", "", 0, "C", true, 0, "")
	pdf.CellFormat(18, 5, "(ไม่รวม VAT 7%)", "", 0, "C", true, 0, "")
	pdf.CellFormat(18, 5, "(รวม VAT 7%)", "", 0, "C", true, 0, "")
	pdf.CellFormat(18, 5, "(รวม VAT 7%)", "", 0, "C", true, 0, "")
	pdf.SetFillColor(196, 16, 22)
	pdf.CellFormat(18, 5, "(รวม VAT 7%)", "", 0, "C", true, 0, "")
	pdf.SetFillColor(0, 0, 0)

	pdf.Ln(-1)

	pdf.SetTextColor(0, 0, 0)

	pdf.SetFont("Sarabun", "", 10)
	counter := 01
	total := 0

	// //บรรทัดเดียว loop รายละเอียดรถยนต์
	// for _, vehicle := range vehicles {
	// 	pdf.SetFillColor(240, 240, 240)
	// 	pdf.SetDrawColor(255, 255, 255)
	// 	pdf.CellFormat(5, 8, vehicle.VehicleNo, "1", 0, "C", true, 0, "")

	// 	pdf.CellFormat(85, 8, vehicle.License+" "+vehicle.LicenseProvinceName+" "+
	// 		vehicle.YearManuFacturing+""+vehicle.VehicleBrandName+" "+vehicle.VehicleModelName+" "+vehicle.VehicleSubModelName+" "+vehicle.VehicleColorName, "1", 0, "", true, 0, "")

	// 	imcvat, _ := strconv.ParseFloat(vehicle.ClosePrice, 64)
	// 	vatPrice := float64(float64(imcvat) * 1.07)
	// 	c10 := float64(float64(vatPrice) * 0.1)
	// 	c90 := float64(float64(vatPrice) * 0.9)

	// 	formattedVatPrice := common.FloatWithCommas(vatPrice, 2)
	// 	formattedTotal := common.FloatWithCommas(imcvat, 2)
	// 	formattedTotal10 := common.FloatWithCommas(c10, 2)
	// 	formattedTotal90 := common.FloatWithCommas(c90, 2)

	// 	pdf.CellFormat(8, 8, vehicle.VehicleGradeID, "1", 0, "C", true, 0, "")
	// 	pdf.CellFormat(20, 8, vehicle.BranchLabel, "1", 0, "C", true, 0, "")
	// 	pdf.CellFormat(18, 8, formattedTotal, "1", 0, "R", true, 0, "")
	// 	pdf.CellFormat(18, 8, formattedTotal10, "1", 0, "R", true, 0, "")
	// 	pdf.CellFormat(18, 8, formattedTotal90, "1", 0, "R", true, 0, "")
	// 	pdf.CellFormat(18, 8, formattedVatPrice, "1", 0, "R", true, 0, "")

	// 	pdf.Ln(-1)

	// 	pdf.SetFillColor(240, 240, 240)

	// 	temp, _ := strconv.Atoi(vehicle.ClosePrice)
	// 	total += temp
	// 	counter++
	// }

	// pdf.Ln(-1)

	//code รายละเอียดรถยนต์ 2 บรรทัด
	for _, vehicle := range vehicles {
		pdf.SetFillColor(240, 240, 240) // Light gray color for the cell background
		pdf.SetDrawColor(255, 255, 255) // Set border color to white

		// Initial position
		x, y := pdf.GetXY()

		// Draw the first cell
		pdf.CellFormat(5, 8, strconv.Itoa(counter), "1", 0, "C", true, 0, "")

		startY := y
		pdf.SetXY(x+5, y)

		// Calculate the height required for the MultiCell content
		pdf.MultiCell(85, 4, vehicle.License+" "+vehicle.LicenseProvinceName+
			vehicle.YearManuFacturing+" "+"\n"+vehicle.VehicleBrandName+" "+
			vehicle.VehicleModelName+" "+vehicle.VehicleSubModelName+" "+vehicle.VehicleColorName, "", "", false)

		currentY := pdf.GetY()
		height := currentY - startY

		// Adjust the height for all cells in the same row to match the height of the MultiCell
		pdf.SetXY(x+5, startY)
		pdf.Rect(x+5, startY, 85, height, "F") // Draw the background rectangle
		pdf.SetXY(x+5, startY)
		pdf.MultiCell(85, 4, vehicle.License+" "+vehicle.LicenseProvinceName+
			vehicle.YearManuFacturing+" "+"\n"+vehicle.VehicleBrandName+" "+
			vehicle.VehicleModelName+" "+vehicle.VehicleSubModelName+" "+vehicle.VehicleColorName, "", "", false)
		// Draw other cells with the calculated height
		pdf.SetXY(x+90, startY)
		pdf.CellFormat(8, height, vehicle.VehicleGradeID, "1", 0, "C", true, 0, "")
		pdf.CellFormat(20, height, vehicle.BranchLabel, "1", 0, "C", true, 0, "")

		imcvat, _ := strconv.ParseFloat(vehicle.ClosePrice, 64)
		vatPrice := float64(float64(imcvat) * 1.07)
		c10 := float64(float64(vatPrice) * 0.1)
		c90 := float64(float64(vatPrice) * 0.9)

		formattedVatPrice := common.FloatWithCommas(vatPrice, 2)
		formattedTotal := common.FloatWithCommas(imcvat, 2)
		formattedTotal10 := common.FloatWithCommas(c10, 2)
		formattedTotal90 := common.FloatWithCommas(c90, 2)

		pdf.SetFont("Sarabun", "", 10)

		pdf.CellFormat(18, height, formattedTotal, "1", 0, "R", true, 0, "")
		pdf.CellFormat(18, height, formattedTotal10, "1", 0, "R", true, 0, "")
		pdf.CellFormat(18, height, formattedTotal90, "1", 0, "R", true, 0, "")
		pdf.CellFormat(18, height, formattedVatPrice, "1", 0, "R", true, 0, "")

		pdf.Ln(-1)
		pdf.SetFillColor(240, 240, 240)

		temp, _ := strconv.Atoi(vehicle.ClosePrice)
		total += temp
		counter++
	}

	vatPrice := float64(float64(total) * 1.07)
	total10 := float64(float64(vatPrice) * 0.1)
	total90 := float64(float64(vatPrice) * 0.9)
	formattedVatPriceTotal := common.FloatWithCommas(vatPrice, 2)
	formattedVatPriceTotalVat := common.FloatWithCommas(vatPrice, 2)
	formattedTotal10 := common.FloatWithCommas(total10, 2)
	formattedTotal90 := common.FloatWithCommas(total90, 2)

	// แสดงผลใน PDF
	pdf.SetFont("Sarabun", "B", 16)
	totalprice := filepathStr + "sumprice.png"
	pdf.Image(totalprice, 50, 230, 150, 15, false, "", 0, "")

	pdf.SetFont("Sarabun", "B", 20)
	xTextTotal10 := 7
	yTextTotal10 := 236
	pdf.SetXY(float64(xTextTotal10), float64(yTextTotal10))
	texttotal10 := formattedTotal10
	pdf.CellFormat(0, 5, texttotal10, "", 0, "C", false, 0, "")
	pdf.Ln(-1)

	pdf.SetFont("Sarabun", "B", 20)
	xTextTotal90 := 83
	yTextTotal90 := 236
	pdf.SetXY(float64(xTextTotal90), float64(yTextTotal90))
	texttotal90 := formattedTotal90
	pdf.CellFormat(0, 5, texttotal90, "", 0, "C", false, 0, "")
	pdf.Ln(-1)

	pdf.SetFont("Sarabun", "B", 20)
	xTextTotalVat := 164
	yTextTotalVat := 236
	pdf.SetXY(float64(xTextTotalVat), float64(yTextTotalVat))
	texttotalvat := formattedVatPriceTotal
	pdf.CellFormat(0, 5, texttotalvat, "", 0, "C", false, 0, "")
	pdf.SetTextColor(0, 0, 0)

	pdf.Ln(-1)

	//สแกนจ่าย
	downscan := filepathStr + "detail-price-new1.png"
	pdf.Image(downscan, 10, 248, 190, 25, false, "", 0, "")
	pdf.Ln(-1)

	pdf.SetFont("Sarabun", "", 10)
	xTextBankName := 11
	yTextBankName := 252
	pdf.SetXY(float64(xTextBankName), float64(yTextBankName))
	bank := summary.BankName
	pdf.CellFormat(0, 5, bank, "", 0, "L", false, 0, "")
	pdf.SetTextColor(0, 0, 0)

	pdf.SetFont("Sarabun", "B", 24)
	xTextBankAccountNo := 65
	yTextBankAccountNo := 258
	pdf.SetXY(float64(xTextBankAccountNo), float64(yTextBankAccountNo))
	bankAccountNo := summary.BankAccountNo
	pdf.CellFormat(45, 5, bankAccountNo, "", 0, "C", false, 0, "")
	pdf.SetTextColor(0, 0, 0)

	pdf.SetFont("Sarabun", "B", 12)
	xTextBankAccount := 15
	yTextBankAccount := 259
	pdf.SetXY(float64(xTextBankAccount), float64(yTextBankAccount))
	bankAccount := summary.BankAccount
	pdf.CellFormat(50, 5, bankAccount, "", 0, "C", false, 0, "")
	pdf.SetTextColor(0, 0, 0)

	qr := summary.BankQr
	httpimg.Register(pdf, qr, "")
	pdf.Image(qr, 178, 252, 17, 17, false, "", 0, "")

	pdf.SetTextColor(255, 0, 0)
	pdf.SetFont("Sarabun", "B", 28)
	xTextTotal := 85
	yTextTotal := 258
	pdf.SetXY(float64(xTextTotal), float64(yTextTotal))
	texttotal := formattedVatPriceTotalVat
	pdf.CellFormat(0, 5, texttotal, "", 0, "C", false, 0, "")
	pdf.SetTextColor(0, 0, 0)

	pdf.SetFont("Sarabun", "", 9)
	xText := 85
	yText := 264
	pdf.SetXY(float64(xText), float64(yText))
	text := common.ConvertFloatToThaiText(vatPrice)
	pdf.CellFormat(0, 5, text, "", 0, "C", false, 0, "")
	pdf.Ln(20)

	pdf.SetTextColor(0, 0, 0)

	line := filepathStr + "sign.png"
	pdf.Image(line, 160, 270, 35, 15, false, "", 0, "")
	pdf.Ln(-1)

}
func preparePayment(auctionID string, userID string) (UserSummary, []VehicleSummary, Summary) {
	//

	user := new(UserSummary)
	common.Database.Raw(`select u.id,up.province_id,up.tax_province_id, u.user_type_id, u.code,u.firstname as user_firstname,u.lastname as user_lastname,
	up.company_name,up.address,up.tax_address,up.street,p.name as province_name,tp.name as tax_province_name,d.name as district_name,td.name as tax_district_name,
	sd.name as sub_district_name,tsd.name as tax_sub_district_name,sd.postcode as post,up.tax_postcode as tax_post_code,
	u.email,u.mobile,up.company_phone,up.is_address_tax
	from users u
	left join user_profiles up on u.id = up.user_id
	left join districts d on up.district_id = d.id 
    left join districts td on up.tax_district_id = td.id 
	left join provinces p on up.province_id = p.id
    left join provinces tp on up.tax_province_id = tp.id
	left join sub_districts sd on up.sub_district_id = sd.id
	left join sub_districts tsd on up.tax_sub_district_id = tsd.id
where
	u.id = ?
						`, userID).Scan(&user)
	// common.Print("resultddddddddddddddd", common.StructToString(user))
	var vehicles []VehicleSummary
	common.Database.Raw(`SELECT v.id,DATE(av.end_date) as end_date ,a.name as auction_name,a.code as code_auction,av.vehicle_no as vehicle_no,b.label as branch_label,v.license, v.license_province_name, v.vehicle_grade_id, v.branch_name, av.close_price,
	v.vehicle_brand_name,v.vehicle_model_name,v.vehicle_sub_model_name,v.year_manufacturing,v.year_register,v.vehicle_color_name,
	v.mile,v.engine_size,v.engine_no,v.chassis_no,v.remark,m1.image_path as car_image,m2.image_path as license_image from auction_vehicles av
	LEFT JOIN vehicles v ON av.vehicle_id = v.id
	LEFT JOIN auctions a ON av.auction_id = a.id
	LEFT JOIN branches b ON v.branch_id = b.id
	LEFT JOIN vehicle_images m1 ON m1.vehicle_id = av.vehicle_id AND m1.id = "7fc5b459-d01b-45aa-a871-86156d827342"
	LEFT JOIN vehicle_images m2 ON m2.vehicle_id = av.vehicle_id AND m2.id = "1665c937-403f-42cc-9e74-a3c4a5195981"
	LEFT JOIN vehicle_grade_remarks vgr ON v.chassis_no = vgr.chassis_number AND v.license = vgr.license
	WHERE av.auction_id = ? AND av.winner_user_id = ? AND av.is_win = 1 order by vehicle_no
						`, auctionID, userID).Debug().Scan(&vehicles)
	// common.Print("vehicle", common.StructToString(vehicles))

	summary := new(Summary)
	common.Database.Raw(`SELECT MAX(r.name) AS round_name, COUNT(*) AS total_car, SUM(av.close_price) AS total_price,  MAX(a.name) AS auction_name,
	rp.bank_name, rp.bank_account_no,  rp.bank_account, rp.bank_qr
	FROM auction_vehicles av
    LEFT JOIN vehicles v ON av.vehicle_id = v.id
    LEFT JOIN auctions a ON av.auction_id = a.id
    LEFT JOIN rounds r ON a.round_id = r.id
    LEFT JOIN round_payments rp ON r.round_payment_id = rp.id
	WHERE av.auction_id = ? AND av.winner_user_id = ? AND av.is_win = 1
	GROUP BY  rp.bank_name, rp.bank_account_no,  rp.bank_account,  rp.bank_qr`, auctionID, userID).Debug().Scan(&summary)
	// common.Print(auctionID, "auctionID")

	return *user, vehicles, *summary
}

type Image struct {
	VehicleID string `json:"vehicle_id"`
	ImageName string `json:"image_name"`
	ImagePath string `json:"image_path"`
}

type VehicleImage struct {
	Vehicle model.VehiclePDF
	Images  []Image `json:"images"`
}

func header(pdf *gofpdf.Fpdf) {

	filepathStr := "images/pdf/"
	pdf.AddUTF8Font("Sarabun", "", "fonts/THSarabun.ttf")
	pdf.AddUTF8Font("Sarabun", "B", "fonts/THSarabunBold.ttf")

	// pdf.SetLineWidth(0.2) // Set line width
	// pdf.Line(10, 48, 200, 48)
	// pdf.Ln(5)
	// pdf.SetLineWidth(0) // Set line width

	x := 8.00
	y := 4.00
	width := 30.00
	height := 18.00

	imagePath := filepathStr + "simple.png"
	pdf.Image(imagePath, x, y, width, height, false, "", 0, "")

	topHandle := filepathStr + "top.png"
	pdf.Image(topHandle, 90, -3, 130, 40, false, "", 0, "")
	pdf.SetFont("Sarabun", "B", 18)
	pdf.Text(45, 14, "บริษัท โอมาคาเสะ คาร์ จำกัด")

	// boxLicense := filepathStr + "text-box-top.png"
	// pdf.Image(boxLicense, 10, 25, 190, 15, false, "", 0, "")

	pdf.SetTextColor(169, 169, 169)
	pdf.SetFont("Sarabun", "B", 14)
	pdf.Text(45, 20, "Omakase Car Online Auction")
	pdf.SetTextColor(0, 0, 0)

	imageDown := filepathStr + "down.png"
	pdf.Image(imageDown, 0, 285, 210, 20, false, "", 0, "")

	pdf.SetFont("Sarabun", "B", 12)
	pdf.SetTextColor(255, 255, 255)

	imageMap := filepathStr + "map W.png"
	pdf.Image(imageMap, 10, 292, 3, 3, false, "", 0, "")
	pdf.Text(14, 295, "2 ซอยลาดพร้าว 21 ถนนลาดพร้าว แขวงจอมพล เขตจตุจักร กรุงเทพมหานคร 10900")

	imagePhone := filepathStr + "phone W.png"
	pdf.Image(imagePhone, 125, 292, 3, 3, false, "", 0, "")
	pdf.Text(130, 295, "082-250-8888")

	imageWebsite := filepathStr + "website W.png"
	pdf.Image(imageWebsite, 165, 292, 3, 3, false, "", 0, "")
	pdf.Text(170, 295, "www.omakasecar.com")

	pdf.SetTextColor(0, 0, 0)

}

func genVehicleImages(pdf *gofpdf.Fpdf, v VehicleImage) {
	filepathStr := "images/pdf/"

	pdf.AddPage()
	header(pdf)
	boxLicense := filepathStr + "text-box-top.png"
	pdf.Image(boxLicense, 10, 25, 190, 15, false, "", 0, "")
	pdf.AddUTF8Font("Sarabun", "B", "fonts/THSarabunBold.ttf")

	pdf.SetTextColor(105, 105, 105)
	pdf.Text(20, 30, "เลขทะเบียนรถยนต์")
	pdf.SetFont("Sarabun", "B", 18)
	pdf.SetTextColor(0, 0, 0)
	pdf.SetFont("Sarabun", "B", 22)
	pdf.Text(20, 38, v.Vehicle.License+" "+v.Vehicle.LicenseProvinceName)

	// //หมายเหตุ
	// borderBox := filepathStr + "border-down.png"
	// pdf.Image(borderBox, 12, 258, 180, 25, false, "", 0, "")

	// pdf.SetTextColor(169, 169, 169)
	// pdf.SetFont("Sarabun", "B", 12)
	// pdf.Text(15, 265, "หมายเหตุ")
	// pdf.SetTextColor(0, 0, 0)

	// pdf.Text(15, 270, "คุณสามารถดาวน์โหลดรูปภาพได้จากลิงก์")

	// url, qr := genPDFLinkDownload(v.Vehicle.ID, v.Vehicle.ID)
	// pdf.SetTextColor(255, 0, 0)
	// pdf.Text(15, 275, url)
	// pdf.SetTextColor(0, 0, 0)

	// pdf.SetFont("Sarabun", "B", 12)
	// pdf.SetTextColor(169, 169, 169)
	// pdf.Text(169, 263, "หรือสแกนได้ที่นี่")
	// webScan := filepathStr + qr
	// pdf.Image(webScan, 170, 265, 15, 15, false, "", 0, "")

	//down
	pdf.SetTextColor(255, 255, 255)

	downHandle := filepathStr + "down.png"
	pdf.Image(downHandle, 0, 285, 210, 20, false, "", 0, "")

	mapIconPath := filepathStr + "map W.png"
	pdf.SetFont("Sarabun", "B", 12)
	pdf.Image(mapIconPath, 10, 292, 3, 3, false, "", 0, "")
	pdf.Text(14, 295, "2 ซอยลาดพร้าว 21 ถนนลาดพร้าว แขวงจอมพล เขตจตุจักร กรุงเทพมหานคร 10900")

	phoneIconPath := filepathStr + "phone W.png"
	pdf.SetFont("Sarabun", "B", 12)
	pdf.Image(phoneIconPath, 125, 292, 3, 3, false, "", 0, "")
	pdf.Text(130, 295, "082-250-8888")

	webIconPath := filepathStr + "website W.png"
	pdf.SetFont("Sarabun", "B", 12)
	pdf.Image(webIconPath, 165, 292, 3, 3, false, "", 0, "")
	pdf.Text(170, 295, "www.omakasecar.com")
	pdf.SetTextColor(0, 0, 0)

	y := 4.00
	y += 5.00
	pdf.Ln(5)
	//loop page

	// check if ImageCar is not empty

	row := 0
	for _, i := range v.Images {

		// url := "https://cdn.recast.ai/newsletter/city-01.png"
		httpimg.Register(pdf, i.ImagePath, "")
		if i.ImageName != "ใบเกรด 1" && i.ImageName != "ใบเกรด 2" && i.ImageName != "ใบเกรด 3" && i.ImageName != "สำเนาเล่มทะเบียน" && i.ImageName != "สำเนาเล่มทะเบียน 14" && i.ImageName != "สำเนาเล่มทะเบียน 16" && i.ImageName != "สำเนาเล่มทะเบียน 18" {

			if row == 0 {
				pdf.SetXY(10, 45)

				pdf.CellFormat(60, 8, i.ImageName, "0", 0, "C", true, 0, "")
				pdf.Image(i.ImagePath, 10, 52, 60, 40, false, "", 0, "")
			}
			if row == 1 {

				pdf.SetXY(75, 45)

				pdf.CellFormat(60, 8, i.ImageName, "0", 0, "C", true, 0, "")

				pdf.Image(i.ImagePath, 75, 52, 60, 40, false, "", 0, "")
			}
			if row == 2 {

				pdf.SetXY(140, 45)

				pdf.CellFormat(60, 8, i.ImageName, "0", 0, "C", true, 0, "")
				pdf.Image(i.ImagePath, 140, 52, 60, 40, false, "", 0, "")
			}

			if row == 3 {
				pdf.SetXY(10, 97)

				pdf.CellFormat(60, 8, i.ImageName, "", 0, "C", true, 0, "")
				pdf.Image(i.ImagePath, 10, 105, 60, 40, false, "", 0, "")
			}
			if row == 4 {
				pdf.SetXY(75, 97)

				pdf.CellFormat(60, 8, i.ImageName, "", 0, "C", true, 0, "")
				pdf.Image(i.ImagePath, 75, 105, 60, 40, false, "", 0, "")
			}
			if row == 5 {

				pdf.SetXY(140, 97)

				pdf.CellFormat(60, 8, i.ImageName, "", 0, "C", true, 0, "")
				pdf.Image(i.ImagePath, 140, 105, 60, 40, false, "", 0, "")
			}

			if row == 6 {
				pdf.SetXY(10, 150)

				pdf.CellFormat(60, 8, i.ImageName, "", 0, "C", true, 0, "")

				pdf.Image(i.ImagePath, 10, 158, 60, 40, false, "", 0, "")
			}
			if row == 7 {
				pdf.SetXY(75, 150)

				pdf.CellFormat(60, 8, i.ImageName, "", 0, "C", true, 0, "")

				pdf.Image(i.ImagePath, 75, 158, 60, 40, false, "", 0, "")
			}
			if row == 8 {
				pdf.SetXY(140, 150)

				pdf.CellFormat(60, 8, i.ImageName, "", 0, "C", true, 0, "")

				pdf.Image(i.ImagePath, 140, 158, 60, 40, false, "", 0, "")
			}

			if row == 9 {
				pdf.SetXY(10, 204)

				pdf.CellFormat(60, 8, i.ImageName, "", 0, "C", true, 0, "")
				pdf.Image(i.ImagePath, 10, 211, 60, 40, false, "", 0, "")
			}
			if row == 10 {
				pdf.SetXY(75, 204)

				pdf.CellFormat(60, 8, i.ImageName, "", 0, "C", true, 0, "")
				pdf.Image(i.ImagePath, 75, 211, 60, 40, false, "", 0, "")
			}
			if row == 11 {
				pdf.SetXY(140, 204)

				pdf.CellFormat(60, 8, i.ImageName, "", 0, "C", true, 0, "")
				pdf.Image(i.ImagePath, 140, 211, 60, 40, false, "", 0, "")
			}
		}

		if i.ImageName != "ใบเกรด 1" && i.ImageName != "ใบเกรด 2" && i.ImageName != "ใบเกรด 3" {

		} else {
			if i.ImageName == "ใบเกรด 1" {
				pdf.AddPage()
				header(pdf)
				boxLicenseT := filepathStr + "text-box-top.png"
				pdf.Image(boxLicenseT, 10, 25, 190, 15, false, "", 0, "")

				pdf.SetTextColor(169, 169, 169)
				pdf.Text(20, 30, "เลขทะเบียนรถยนต์")
				pdf.SetFont("Sarabun", "B", 16)
				pdf.SetTextColor(0, 0, 0)
				pdf.SetFont("Sarabun", "B", 22)
				pdf.Text(20, 38, v.Vehicle.License+" "+v.Vehicle.LicenseProvinceName)

				//รายละเอียดตามเกรดรถยนต์
				boxLicense := filepathStr + "text-licens.png"
				pdf.Image(boxLicense, 10, 50, 190, 15, false, "", 0, "")

				pdf.SetFont("Sarabun", "B", 20)
				pdf.Text(80, 60, "รายละเอียดตามเกรดรถยนต์")

			}
			// pdf.Image(i.ImagePath, 10, 70, 94, 125, false, "", 0, "")
			if i.ImageName == "ใบเกรด 1" {
				pdf.Image(i.ImagePath, 10, 70, 94, 125, false, "", 0, "")

			} else if i.ImageName == "ใบเกรด 2" {

				pdf.Image(i.ImagePath, 106, 70, 94, 125, false, "", 0, "")

			} else if i.ImageName == "ใบเกรด 3" {
				pdf.AddPage()
				header(pdf)
				boxLicenseT := filepathStr + "text-box-top.png"
				pdf.Image(boxLicenseT, 10, 25, 190, 15, false, "", 0, "")

				pdf.SetTextColor(169, 169, 169)
				pdf.Text(20, 30, "เลขทะเบียนรถยนต์")
				pdf.SetFont("Sarabun", "B", 16)
				pdf.SetTextColor(0, 0, 0)
				pdf.SetFont("Sarabun", "B", 22)
				pdf.Text(20, 38, v.Vehicle.License+" "+v.Vehicle.LicenseProvinceName)

				//รายละเอียดตามเกรดรถยนต์
				boxLicense := filepathStr + "text-licens.png"
				pdf.Image(boxLicense, 10, 50, 190, 15, false, "", 0, "")

				pdf.SetFont("Sarabun", "B", 20)
				pdf.Text(80, 60, "รายละเอียดตามเกรดรถยนต์")
				pdf.Image(i.ImagePath, 10, 70, 94, 125, false, "", 0, "")

			}
		}

		if i.ImageName != "สำเนาเล่มทะเบียน" && i.ImageName != "สำเนาเล่มทะเบียน 14" && i.ImageName != "สำเนาเล่มทะเบียน 16" && i.ImageName != "สำเนาเล่มทะเบียน 18" {

		} else {
			if i.ImageName == "สำเนาเล่มทะเบียน" {

				pdf.AddPage()
				boxLicenseT := filepathStr + "text-box-top.png"
				pdf.Image(boxLicenseT, 10, 25, 190, 15, false, "", 0, "")
				header(pdf)
				pdf.SetTextColor(169, 169, 169)
				pdf.Text(20, 30, "เลขทะเบียนรถยนต์")
				pdf.SetFont("Sarabun", "B", 16)
				pdf.SetTextColor(0, 0, 0)
				pdf.SetFont("Sarabun", "B", 22)
				pdf.Text(20, 38, v.Vehicle.License+" "+v.Vehicle.LicenseProvinceName)

				//รายละเอียดตามเกรดรถยนต์
				boxLicense := filepathStr + "text-licens.png"
				pdf.Image(boxLicense, 10, 50, 190, 15, false, "", 0, "")

				pdf.SetFont("Sarabun", "B", 20)
				pdf.Text(80, 60, "เล่มทะเบียนรถยนต์")

			}
			if i.ImageName == "สำเนาเล่มทะเบียน" {
				pdf.Image(i.ImagePath, 8, 70, 94, 125, false, "", 0, "")

			} else if i.ImageName == "สำเนาเล่มทะเบียน 14" {

				pdf.Image(i.ImagePath, 104, 70, 94, 125, false, "", 0, "")

			} else if i.ImageName == "สำเนาเล่มทะเบียน 16" {

				pdf.AddPage()
				boxLicenseT := filepathStr + "text-box-top.png"
				pdf.Image(boxLicenseT, 10, 25, 190, 15, false, "", 0, "")
				header(pdf)
				pdf.SetTextColor(169, 169, 169)
				pdf.Text(20, 30, "เลขทะเบียนรถยนต์")
				pdf.SetFont("Sarabun", "B", 16)
				pdf.SetTextColor(0, 0, 0)
				pdf.SetFont("Sarabun", "B", 22)
				pdf.Text(20, 38, v.Vehicle.License+" "+v.Vehicle.LicenseProvinceName)

				//รายละเอียดตามเกรดรถยนต์
				boxLicense := filepathStr + "text-licens.png"
				pdf.Image(boxLicense, 10, 50, 190, 15, false, "", 0, "")

				pdf.SetFont("Sarabun", "B", 20)
				pdf.Text(80, 60, "เล่มทะเบียนรถยนต์")
				pdf.Image(i.ImagePath, 10, 70, 94, 125, false, "", 0, "")

			} else if i.ImageName == "สำเนาเล่มทะเบียน 18" {

				pdf.Image(i.ImagePath, 106, 70, 94, 125, false, "", 0, "")

			}

		}

		row++

	}

}
func prepareDataVehicle(auctionID string, userID string) []VehicleImage {

	var vehicles []model.VehiclePDF
	common.Database.Raw(`select v.*,v.license_province_name  from vehicles v
	left join auction_vehicles av on v.id = av.vehicle_id 
	where av.auction_id = ? AND av.winner_user_id = ?
	AND av.is_win = 1 order by vehicle_no`, auctionID, userID).Scan(&vehicles)
	// common.Print(auctionID, userID)
	var images []Image
	common.Database.Raw(`
	SELECT vem.vehicle_id,vem.image_path,vm.name as image_name from template_download_vehicle_images dm
	left join template_vehicle_images vm on dm.vehicle_image_id = vm.id
	left join vehicle_images vem on vm.id = vem.template_vehicle_image_id
	where vem.vehicle_id IN (select v.id from vehicles v
		left join auction_vehicles av on v.id = av.vehicle_id 
		where av.auction_id = ? AND av.winner_user_id = ?
		AND av.is_win = 1) and vem.is_delete = 0 
	ORDER BY vehicle_id , dm.row_order`, auctionID, userID).Scan(&images)
	// common.Print("vehicle", common.StructToString(vehicles))
	// common.Print("image", common.StructToString(images))

	var vehicleImage []VehicleImage
	for _, v := range vehicles {
		vi := new(VehicleImage)
		vi.Vehicle = v

		var is []Image
		for _, i := range images {
			if i.VehicleID == v.ID {
				is = append(is, i)
			}
		}
		vi.Images = is

		vehicleImage = append(vehicleImage, *vi)
	}
	// common.Print("finalarray", common.StructToString(vehicleImage))

	return vehicleImage
}
func getThaiMonth(month string) string {
	thaiMonths := map[string]string{
		"January":   "มกราคม",
		"February":  "กุมภาพันธ์",
		"March":     "มีนาคม",
		"April":     "เมษายน",
		"May":       "พฤษภาคม",
		"June":      "มิถุนายน",
		"July":      "กรกฎาคม",
		"August":    "สิงหาคม",
		"September": "กันยายน",
		"October":   "ตุลาคม",
		"November":  "พฤศจิกายน",
		"December":  "ธันวาคม",
	}
	return thaiMonths[month]
}
