package utils

func GenPDFDownloadVehicleImage(auctionID string, userID string) (string, error) {
	return "https://storage.googleapis.com/fourwd-auction/auction/2023/Vehicle-Image.pdf", nil
}

func GenPDFVehicle(auctionID string, userID string, vehicles []string) (string, error) { //carlist
	return "https://storage.googleapis.com/fourwd-auction/auction/2023/0000.pdf", nil
}
func GenPDFAuctionResult(auctionID string) (string, error) {
	return "https://storage.googleapis.com/fourwd-auction/auction/2023/1.png", nil
}
func GenPDFPayment(auctionID string, userID string) (string, error) { //ใบแจ้งผู้ประมูลราคาสูงสุด
	return "https://storage.googleapis.com/fourwd-auction/auction/2023/2.pdf", nil
}
