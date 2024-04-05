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

// func GenPDFImageCar(auctionID string) (string, error) { //หน้าใบปลิวรถ
// 	pdf := gofpdf.New("P", "mm", "A4", "")

// 	// auctionID := c.Params("auction_id")
// 	GenPDFImageCarDetail(pdf, auctionID)

// 	path, err := common.UploadPdfToGoogle(pdf, "ใบปลิวภาพรถ", "auction", "fourwd-auction")
// 	if err != nil {
// 		return "", err
// 	}
// 	return path, nil
// }

func GenPDFImageCarDetail(auctionID string) (string, error) { //หน้าใบปลิวรถ
	vehicles := prepareDetailList(auctionID)

	filepathStr := "images/pdf/"
	// fileextention := ".pdf"
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.AddUTF8Font("Sarabun", "", "fonts/THSarabun.ttf")
	pdf.AddUTF8Font("Sarabun", "B", "fonts/THSarabunBold.ttf")

	// กำหนดสีพื้นหลังเป็นสีแดง
	pdf.SetFillColor(196, 16, 22)
	pdf.Rect(0, 0, 210, 297, "F")

	x := 2.5
	y := 75.0
	imagesInRow := 0

	// location := filepathStr + "m-removebg-preview.png"
	// page := 1
	// imageperpage := 18
	// imageperpage2 := 24
	imageindex := 1
	page := 1
	// perpage := 18
	// line := 1
	for i, v := range vehicles {

		i++
		if page == 1 {

		} else {

			y = 10
			// perpage = 24
			pdf.SetFillColor(196, 16, 22)
			pdf.Rect(0, 0, 210, 297, "F")
		}

		// fmt.Printf("index = %d, License = %s", i, v.License)
		// fmt.Printf(", page = %d, perpage = %d", page, perpage)
		// fmt.Printf(", imageindex = %d, line = %d", imageindex, line)
		// fmt.Println("========================")

		if i == 1 {
			header := filepathStr + "bar-top.jpg"
			pdf.Image(header, 0, 0, 210, 70, false, "", 0, "")
		}

		if v.ImagePreviewPath != "" {

			pdf.SetFillColor(255, 255, 255)
			pdf.Rect(x, y, 67, 30, "F")
			pdf.Rect(x, y, 67, 30, "D")

			// ลงทะเบียนรูปภาพ
			pdf.SetFont("Sarabun", "", 7)
			httpimg.Register(pdf, v.ImagePreviewPath, "")
			pdf.Image(v.ImagePreviewPath, x+1, y+1, 32, 22, false, "", 0, "")
			pdf.SetTextColor(255, 255, 255)
			pdf.SetFillColor(133, 133, 133)
			pdf.RoundedRect(x+23, y+18.3, 5, 4.75, 0.4, "1,2", "F")
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
				openprice = "฿" + common.FloatWithCommas(openPriceFloat, 2)

			}
			pdf.Text(x+2.5, y+26.5, "ราคาเปิด ")
			pdf.Text(x+21, y+26.5, openprice)
			// pdf.CellFormat(25, 8, openPriceComma, "1", 0, "R", true, 0, "")

			// เพิ่มข้อความรายละเอียดรถ
			pdf.SetFont("Sarabun", "B", 12)
			pdf.SetTextColor(0, 0, 0)

			pdf.Text(x+34, y+4, "ปี "+v.YearRegister+" "+v.VehicleBrandName)
			pdf.SetFont("Sarabun", "", 8)
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

			// pdf.Text(x+34.5, y+19-2, "&#xf3c5")
			imagePaths := textcolor([]string{v.TextColor})

			for _, imagePath := range imagePaths {
				pdf.Image(
					imagePath, x+34.5, y+19-2, 2.5, 2.5, false, "", 0, "")
			}
			pdf.SetTextColor(0, 0, 0)

			pdf.Text(x+37.5, y+19, v.BranchLabel)
			pdf.SetTextColor(0, 0, 0)

			pdf.Text(x+34, y+23, "ผู้ถือกรรมสิทธิ์ : ")
			pdf.Text(x+34, y+27, v.SourceName)

			x += 69.0
			imagesInRow++
			//18+24

			if imagesInRow >= 3 {

				// รีเซ็ตค่า x และเพิ่มค่า y เพื่อเริ่มต้นรูปใหม่ในแถวถัดไป
				x = 2.5
				y += 32.0
				imagesInRow = 0

			}

			headerdown := filepathStr + "bar-down.jpg"
			pdf.Image(headerdown, 0, 274, 210, 23, false, "", 0, "")
		}

		imageindex++
		newpage := false
		if i == 18 {
			newpage = true

		} else if i > 18 {
			temp := i - 18
			if temp%24 == 0 {
				newpage = true
			}
		}
		if newpage {
			pdf.AddPage()
			pdf.SetFillColor(196, 16, 22)
			pdf.Rect(0, 0, 210, 297, "F")
			x = 2.5
			y = 15
			imageindex = 1
			// imagesInRow = 0

		}
		i++
	}
	// common.Print("len", fmt.Sprintf("%d", len(vehicles)))

	// filedesination := filepathStr + "ใบปลิวหน้ารถ" + fileextention
	// err := pdf.OutputFileAndClose(filedesination)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// }
	// return filedesination
	path, err := common.UploadPdfToGoogle(pdf, "ใบปลิวรถ", "auction", "fourwd-auction")
	if err != nil {
		return "", err
	}
	return path, nil
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
		output, err := os.CreateTemp("", color+".png")
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
