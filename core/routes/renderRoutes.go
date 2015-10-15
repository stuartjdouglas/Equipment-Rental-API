package routes

import (
	"../router"
//	"encoding/json"
	"net/http"
	"fmt"
	"github.com/zenazn/goji/web"
	//"html/template"
"html/template"
//	"encoding/json"
)

func CreateRenderRoutes(api router.API) {
	api.Router.Get("/hello/:name", func (c web.C, res http.ResponseWriter, r *http.Request) {
		message := hello{
			Message: c.URLParams["name"],
		}


//		data, err := json.Marshal(message)
//		if err != nil {
//			http.Error(res, err.Error(), http.StatusInternalServerError)
//			return
//		}

//		res.Header().Set("Content-Type", "application/json")
//		res.WriteHeader(200)
//		json.NewEncoder(res).Encode(data)
//		data := message.Message

		data := fmt.Sprintf("%s!", message.Message)

		t := template.New("fieldname example")
		t, _ = t.Parse("<h1><b>こんにちは,</b> {{.Body}}</h1>")
		p := Page{Title: "Welcome", Body: data}
		t.Execute(res, p)




//		res.Write(data)
//		t := template.New("some template") // Create a template.
//		t, _ = t.ParseFiles("core/client/hello.html")  // Parse template file.
//		user := GetUser() // Get current user infomration.
//		t.Execute(res, data)
	})
}


type Page struct {
	Title string
	Body  string
}