package models

import (
	"image"
	"github.com/remony/Equipment-Rental-API/core/utils"
)

func GenerateQR(code string, height int, width int) image.Image {
	return utils.GenerateQRCode(code, height, width)
}


