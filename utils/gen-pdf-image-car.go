package utils

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/FourWD/middleware/common"
	"github.com/jung-kurt/gofpdf"
	"github.com/jung-kurt/gofpdf/contrib/httpimg"
)

func GenPDFImageCarDetail(auctionID string) (string, error) { //หน้าใบปลิวรถ
	vehicles := prepareDetailList(auctionID)

	// filepathStr := "images/pdf/"
	// fileextention := ".pdf"
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.AddUTF8Font("Sarabun", "", "fonts/THSarabun.ttf")
	pdf.AddUTF8Font("Sarabun", "B", "fonts/THSarabunBold.ttf")

	// กำหนดสีพื้นหลังเป็นสีแดง
	pdf.SetFillColor(196, 16, 22)
	pdf.Rect(0, 0, 210, 297, "F")
	// // header := filepathStr + "top-list.jpg"
	// header := "https://storage.googleapis.com/fourwd-auction/app/pdf_resource/Bar%20%E0%B9%83%E0%B8%9A%E0%B8%9B%E0%B8%A5%E0%B8%B4%E0%B8%A7-156-1_0.jpg"
	// httpimg.Register(pdf, header, "")
	// // pdf.Image(header, 0, 0, 210, 35, false, "", 0, "")
	// pdf.Image(header, 0, 0, 210, 63, false, "", 0, "")

	header := ""
	sql := `select header_image_verical from auctions where id = ?`
	common.Database.Raw(sql, auctionID).Scan(&header)
	// header := "https://storage.googleapis.com/fourwd-auction/app/pdf_resource/top-list.jpg"
	httpimg.Register(pdf, header, "")
	pdf.Image(header, 0, 0, 210, 63, false, "", 0, "")

	// x := 2.5
	// // y := 75.0
	// y := 6.0
	x := 2.5
	y := 68.0

	imagesInRow := 0
	imageRowCount := 1
	page := 1

	vehicleCount := (len(vehicles))
	fmt.Println("vehicle count original ", vehicleCount)
	var emptypVehicle int
	if vehicleCount <= 21 {
		emptypVehicle = 21 - vehicleCount
	} else {
		vehicleCount2 := vehicleCount - 21
		fmt.Println("vehicle count 2 ", vehicleCount2)

		mod := vehicleCount2 % 27
		if mod != 0 {
			emptypVehicle = 27 - mod
		}
		fmt.Println("mod", mod)
	}

	for i := 1; i <= emptypVehicle; i++ {
		vnull := new(VehicleSummaryDetail)
		vehicles = append(vehicles, *vnull)
	}
	fmt.Println("emptyp", emptypVehicle)
	fmt.Println("vehicle count original + emptyp", len(vehicles))

	currentPage := 1

	for i, v := range vehicles {

		i++
		if page == 1 {

		} else {

			y = 0
			// perpage = 24
			pdf.SetFillColor(196, 16, 22)
			pdf.Rect(0, 0, 210, 297, "F")
		}
		if v.License != "" {
			pdf.SetFillColor(255, 255, 255)
			pdf.Rect(x, y, 67, 29, "F")
			pdf.Rect(x, y, 67, 29, "D")

			// ลงทะเบียนรูปภาพ
			pdf.SetFont("Sarabun", "", 7)
			if v.ImagePreviewPath != "" {

				httpimg.Register(pdf, v.ImagePreviewPath, "")
				pdf.Image(v.ImagePreviewPath, x+1, y+1, 32, 0, false, "", 0, "")
				pdf.SetTextColor(255, 255, 255)
				pdf.SetFillColor(133, 133, 133)
				pdf.RoundedRect(x+23, y+18.3, 5, 4.75, 0.4, "1,2", "F")
			} else {
				logooma := "https://storage.googleapis.com/fourwd-auction/app/pdf_resource/Logo-BG-003.png"
				httpimg.Register(pdf, logooma, "")
				// logo := filepathStr + "Logo-BG-003.png"
				// httpimg.Register(pdf, v.ImagePreviewPath, "")
				pdf.Image(logooma, x+1, y+1, 32, 22, false, "", 0, "")
			}
			pdf.Text(x+23.5, y+20, "เลขรถ")
			pdf.Text(x+25, y+22, v.VehicleNo)

			// ใช้ฟังก์ชัน hexToRGB เพื่อแปลงค่าสี HAX เป็น RGB
			r, g, b := hexToRGB(v.ColorCode)
			pdf.SetFillColor(r, g, b)
			pdf.RoundedRect(x+28, y+18.40, 5, 4.65, 0.4, "1", "F")
			pdf.Text(x+29, y+20, "เกรด")
			pdf.Text(x+30, y+22, v.VehicleGradeID)
			pdf.SetTextColor(0, 0, 0)

			pdf.SetFont("Sarabun", "", 10)
			pdf.SetTextColor(255, 255, 255)
			pdf.SetFillColor(46, 45, 45)
			pdf.RoundedRect(x+1, y+23.25, 32, 4.5, 0.4, "3,4", "F")

			openPriceFloat, _ := strconv.ParseFloat(v.OpenPrice, 64)
			// openPriceComma := common.FloatWithCommas(openPriceFloat, 0)

			var openprice string
			if openPriceFloat == 0 {
				pdf.SetFont("Sarabun", "", 6)
				openprice = "รอกำหนดราคา"
			} else {
				openprice = common.FloatWithCommas(openPriceFloat, 0)
			}
			pdf.Text(x+2.5, y+26.5, "ราคาเปิด ")
			pdf.Text(x+21, y+26.5, "฿"+openprice)
			// pdf.CellFormat(25, 8, openPriceComma, "1", 0, "R", true, 0, "")

			// เพิ่มข้อความรายละเอียดรถ
			pdf.SetFont("Sarabun", "B", 12)
			pdf.SetTextColor(0, 0, 0)

			pdf.Text(x+34, y+4, "ปี "+v.YearRegister+" "+v.VehicleBrandName)
			pdf.SetFont("Sarabun", "", 7)
			pdf.Text(x+34, y+7, v.VehicleModelName+" "+v.VehicleSubModelName)

			pdf.Text(x+34, y+11, "ทะเบียน : "+v.License+" "+v.LicenseProvinceName)
			mile, _ := strconv.ParseFloat(v.Mile, 64)
			milecomma := common.FloatWithCommas(mile, 0)
			pdf.Text(x+34, y+15, "เลขไมล์ : "+milecomma)

			a, c, d := hexToRGB(v.BackgroundColor)
			pdf.SetFillColor(a, c, d)
			// กำหนดความกว้างของข้อความ
			textWidth := pdf.GetStringWidth("ab" + " " + v.BranchLabel)
			pdf.RoundedRect(x+34, y+19-2.5, textWidth+2, 3.65, 0.75, "1,2,3,4", "F")

			// // pdf.Text(x+34.5, y+19-2, "&#xf3c5")
			// imagePaths := textcolor([]string{v.TextColor})

			// for _, imagePath := range imagePaths {
			// 	pdf.Image(
			// 		imagePath, x+34.5, y+19-2, 2.5, 2.5, false, "", 0, "")
			// }
			imagePaths := textcolor2(v.TextColor)
			httpimg.Register(pdf, imagePaths, "")
			println(imagePaths)
			// for _, imagePath := range imagePaths {
			pdf.Image(imagePaths, x+34.5, y+19-2, 2.5, 2.5, false, "", 0, "")
			// }

			pdf.SetTextColor(0, 0, 0)

			pdf.Text(x+37.5, y+19, v.BranchLabel)
			pdf.SetTextColor(0, 0, 0)

			pdf.Text(x+34, y+23, "ผู้ถือกรรมสิทธิ์ : ")
			pdf.Text(x+34, y+27, v.SourceName)

		} else {
			pdf.Rect(x, y, 67, 28.9, "D") // กล่องภาพ
			logoX := x
			logoY := y
			logoWidth := 67.0
			logoHeight := 28.9
			logo := "https://storage.googleapis.com/fourwd-auction/app/pdf_resource/logo-omakase.png"
			pdf.Image(logo, logoX, logoY, logoWidth, logoHeight, false, "", 0, "")
		}
		x += 69.0
		imagesInRow++
		//18+24
		if imagesInRow >= 3 {
			x = 2.5
			y += 31.5
			imagesInRow = 0
			imageRowCount++
		}
		newpage := false
		if i == 21 {
			if vehicleCount > 21 {
				newpage = true
			}
		} else if i > 21 {
			temp := i - 21
			if temp%27 == 0 && i < vehicleCount {
				newpage = true
			}
		}
		if newpage {
			currentPage++
			fmt.Println("currentPage", currentPage)

			fmt.Println("newpage", i)
			pdf.AddPage()
			pdf.SetFillColor(196, 16, 22)
			pdf.Rect(0, 0, 210, 297, "F")
			x = 2.5
			y = 3
			imagesInRow = 0
		}
		i++

		// headerdown := filepathStr + "bar-down.jpg"
		headerdown := ""
		sql := `select bottom_image_verical from auctions where id = ?`
		common.Database.Raw(sql, auctionID).Scan(&headerdown)
		// headerdownpath := "https://storage.googleapis.com/fourwd-auction/app/pdf_resource/bar-down.jpg"
		httpimg.Register(pdf, headerdown, "")
		pdf.Image(headerdown, 0, 279, 210, 18, false, "", 0, "")

		// headerdown := ""
		// sql := `select header_image_horizontal from auctions where id = ?`
		// common.Database.Raw(sql).Scan(&header)
		// // header := "https://storage.googleapis.com/fourwd-auction/app/pdf_resource/top-list.jpg"
		// httpimg.Register(pdf, header, "")
		// pdf.Image(header, 0, 0, 297, 53, false, "", 0, "")

	}

	fileName := generateFileNameImageCar(vehicles[0].AuctionName)

	path, err := common.UploadPdfToGoogle(pdf, fileName, "auction", "fourwd-auction")
	if err != nil {
		return "", err
	}
	return path, nil
}
func textcolor2(color string) string {
	colorWithPercent := strings.Replace(color, "#", "%23", -1)
	return fmt.Sprintf("https://storage.googleapis.com/fourwd-auction/app/pdf_resource/%s.png", colorWithPercent)
}
func generateFileNameImageCar(auctionName string) string {
	auctionName = strings.ReplaceAll(auctionName, " ", "_")
	fileName := "ใบปลิวรถ_" + auctionName
	return fileName
}
func hexToRGB(hex string) (int, int, int) {
	var r, g, b int
	fmt.Sscanf(hex, "#%02x%02x%02x", &r, &g, &b)
	return r, g, b
}

func hexToRGBA(hex string) color.RGBA {
	// ตัดเครื่องหมาย # ออกจากรหัสสี
	hex = strings.TrimPrefix(hex, "#")

	// แปลงสตริงสีเป็นตัวเลขเต็ม (int) 16 บิต
	value, err := strconv.ParseUint(hex, 16, 32)
	if err != nil {
		panic(err)
	}

	// แยกสีแต่ละส่วน (Red, Green, Blue, Alpha) จากค่าสีที่ได้
	r := uint8((value >> 16) & 0xFF)
	g := uint8((value >> 8) & 0xFF)
	b := uint8(value & 0xFF)
	a := uint8(255) // ตั้งค่า alpha เป็นค่าเริ่มต้น

	// สร้างและส่งคืนสี RGBA
	return color.RGBA{r, g, b, a}
}

func textcolor(colors []string) []string {
	var filenames []string
	filepathStr := "images/pdf/"
	location := filepathStr + "m-removebg-preview.png"

	for _, color := range colors {
		// Open the PNG file
		file, err := os.Open(location)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		// Decode the PNG file
		img, _, err := image.Decode(file)
		if err != nil {
			panic(err)
		}

		// Create a new empty image with the same bounds as the original image
		bounds := img.Bounds()
		newImg := image.NewRGBA(bounds)

		// Define the new color you want to set
		newColor := hexToRGBA(color)

		// Iterate over each pixel of the original image
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			for x := bounds.Min.X; x < bounds.Max.X; x++ {
				// Get the original pixel color
				originalColor := img.At(x, y)

				// If the original pixel is not fully transparent, set the new color
				if _, _, _, a := originalColor.RGBA(); a != 0 {
					newImg.Set(x, y, newColor)
				} else {
					// If the original pixel is fully transparent, keep it transparent in the new image
					newImg.Set(x, y, originalColor)
				}
			}
		}

		// Save the modified image to a new PNG file
		outputFilename := filepath.Join(filepathStr, color+".png")
		output, err := os.Create(outputFilename)
		if err != nil {
			panic(err)
		}
		defer output.Close()

		// Encode the new image as PNG
		if err := png.Encode(output, newImg); err != nil {
			panic(err)
		}
		filenames = append(filenames, outputFilename)
	}

	return filenames
}
