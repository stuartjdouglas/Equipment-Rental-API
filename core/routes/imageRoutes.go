package routes

import (
	"../router"
	"github.com/zenazn/goji/web"
"net/http"
	"../models"
	"../utils"
	"io/ioutil"
	"fmt"
	"mime/multipart"
	"os"
	"io"
	"path"
)

func writeToFile(image byte, filename string) {
	fmt.Println(filename);
	path:= "client/images/"
	ioutil.WriteFile(path + filename, []byte("hello"), 0644)
	fmt.Println("write file to desk");
}


func write(image multipart.File, filename string) {
	f, err := os.OpenFile("client/images/" + filename, os.O_WRONLY | os.O_CREATE, 0666)
	if err != nil {
		panic (err)
	}
	io.Copy(f, image)
	image.Close()
}

func generateImageRoutes(api router.API) {

//	TODO Upload


	type image struct {
		Image []byte `json:"image"`
	}

	api.Router.Post("/image/upload", func (c web.C, res http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("token")

		if token != "" {
			if models.IsSessionValid(api, token) {
				file, header, err:= r.FormFile("image")
				if err != nil {
					panic(err)
				}


				write(file, utils.GenerateToken(header.Filename) + path.Ext(header.Filename))

//				res.Header().Set("Content-Type", "application/json")
//				res.WriteHeader(200)
//				res.Write(data)
			} else {
				http.Error(res, "", http.StatusUnauthorized)
			}

		} else {
			http.Error(res, "", http.StatusUnauthorized)
		}
	})

//	TODO Delete

//	Should we store images raw or scale them (to save storage?)

}
