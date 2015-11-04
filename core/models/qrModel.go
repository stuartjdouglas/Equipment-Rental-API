package models

import (
	"image"
	"../utils"
)



func GenerateQR(code string, height int, width int) image.Image{
	return utils.GenerateQRCode(code, height, width)
}


