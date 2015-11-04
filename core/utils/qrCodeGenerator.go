package utils
import (
"image"
"github.com/boombuler/barcode/qr"
	"log"
	"github.com/boombuler/barcode"
)

// GenerateQRCode creates a QR code image with the given code and returns the image
func GenerateQRCode(code string, height int, width int) image.Image {
	// If the height and width is less than 300 set to 300
	if height < 300 || width < 300 {
		height = 300
		width = 300
	}

	image, err := qr.Encode(code, qr.L, qr.Unicode)

	if err != nil {
		log.Println("Error Generating QR code")
	}

	if code != image.Content() {
		log.Fatal("data if different")
	}

	image, err = barcode.Scale(image, height, width)
	if err != nil {
		log.Fatal(err)
	}

	return image
}