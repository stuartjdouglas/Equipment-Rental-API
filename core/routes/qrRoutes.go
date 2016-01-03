package routes

import (
	"github.com/zenazn/goji/web"
	"net/http"
	"bytes"
	"image/jpeg"
	"log"
	"strconv"
	"encoding/base64"
	"github.com/remony/Equipment-Rental-API/core/router"
	"github.com/remony/Equipment-Rental-API/core/models"
)


type QRcode struct {
	Code 	string 	`json:"code"`
	Width 	int 	`json:"width"`
	Height 	int 	`json:"height"`
}

func generateQrRoutes(api router.API) {
	api.Router.Get("/identify/qr/:type", func (c web.C, res http.ResponseWriter, req *http.Request) {
		qr := QRcode{}
		qr.Code = req.Header.Get("code")
		qr.Height, _ = strconv.Atoi(req.Header.Get("height"))
		qr.Width, _ = strconv.Atoi(req.Header.Get("width"))



		switch(c.URLParams["type"]) {
		case "user":
			qr.Code = string('@') + qr.Code
			break;
		}

		if len(qr.Code) < 0 || qr.Height <= 0 || qr.Width < 0 {
			qr = QRcode{
				Code: 	"You have given me wrong params; see docs for more information",
				Height:	500,
				Width:	500,
			}
		} else {

		}

		result := models.GenerateQR(qr.Code, qr.Height, qr.Width)

		buffer := new(bytes.Buffer)
		if err := jpeg.Encode(buffer, result, nil); err != nil {
			log.Println("unable to encode image.")
		}

		imgBase64Str := base64.StdEncoding.EncodeToString(buffer.Bytes())


		res.Header().Set("Content-Type", "image/jpeg")
		res.Header().Set("Content-Length", strconv.Itoa(len(imgBase64Str)))
		if _, err := res.Write([]byte(imgBase64Str)); err != nil {
			log.Println("unable to write image.")
		}

	})
}
