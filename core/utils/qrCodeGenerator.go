package utils
import (
"image"
"github.com/boombuler/barcode/qr"
	"log"
	"github.com/boombuler/barcode"
)


func GenerateQRCode(id string) image.Image {

	code, err := qr.Encode(id, qr.L, qr.Unicode)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Encoded data: ", code.Content())

	if id != code.Content() {
		log.Fatal("data differs")
	}

	code, err = barcode.Scale(code, 300, 300)
	if err != nil {
		log.Fatal(err)
	}

	return code
}