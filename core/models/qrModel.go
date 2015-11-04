package models

import (
	"image"
	"../utils"
)



func GenerateQR(code string) image.Image{

	return utils.GenerateQRCode(code)
}


