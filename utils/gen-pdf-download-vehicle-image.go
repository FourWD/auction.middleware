package utils

func GenPDFDownloadVehicleImage(auctionID string, userID string) (string, error) {
	return "https://storage.googleapis.com/fourwd-auction/app/example_file/Payment-Term-Customer.pdf", nil
}

func GenPDFVehicle(auctionID string, userID string, vehicles []string) (string, error) { //carlist
	return "https://storage.googleapis.com/fourwd-auction/app/example_file/Payment-Term-Customer.pdf", nil
}
func GenPDFAuctionResult(auctionID string) (string, error) {
	return "https://storage.googleapis.com/fourwd-auction/app/example_file/Payment-Term-Customer.pdf", nil
}
func GenPDFPayment(auctionID string, userID string) (string, error) { //ใบแจ้งผู้ประมูลราคาสูงสุด
	return "https://storage.googleapis.com/fourwd-auction/app/example_file/Payment-Term-Customer.pdf", nil
}
