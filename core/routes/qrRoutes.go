package routes

import (
	"../router"
	"github.com/zenazn/goji/web"
	"net/http"
	"../models"
	"bytes"
	"image/jpeg"
	"log"
	"strconv"
)


type QRcode struct {
	Code string `json:"code"`
}

func generateQrRoutes(api router.API) {
	api.Router.Get("/qr", func (c web.C, res http.ResponseWriter, r *http.Request) {

		code := r.Header.Get("code")
			result := models.GenerateQR(code)

		buffer := new(bytes.Buffer)
		if err := jpeg.Encode(buffer, result, nil); err != nil {
			log.Println("unable to encode image.")
		}

		res.Header().Set("Content-Type", "image/jpeg")
		res.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
		if _, err := res.Write(buffer.Bytes()); err != nil {
			log.Println("unable to write image.")
		}
	})
}
