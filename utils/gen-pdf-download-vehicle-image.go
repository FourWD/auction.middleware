package utils

func GenZipDownloadVehicleImage(auctionID string, userID string) (string, error) {
	return "https://storage.googleapis.com/fourwd-auction/auction/2023/%E0%B8%AB%E0%B8%99%E0%B9%89%E0%B8%B2%E0%B8%81%E0%B8%A3%E0%B8%B0%E0%B9%82%E0%B8%9B%E0%B8%A3%E0%B8%87.rar", nil
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
