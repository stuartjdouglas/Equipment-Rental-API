package routes
import(
	"../router"
	"../models"
	"encoding/json"
"net/http"
	"net"
	"fmt"
	"github.com/zenazn/goji/web"
)

func generateAuthRoutes(api router.API)	{
	api.Router.Post("/logout", func (c web.C, res http.ResponseWriter, r *http.Request) {
		//		Not yet implemented
		//		Call method to remove token
	})




	api.Router.Post("/login", func (c web.C, res http.ResponseWriter, r *http.Request) {
		var loginDetails = login{
			Username:r.FormValue("username"),
			Password: r.FormValue("password"),
		}


		if len(loginDetails.Username) == 0 || len(loginDetails.Password) == 0{
			fmt.Println("Get user info from json");

			err := json.NewDecoder(r.Body).Decode(&loginDetails)
			if err != nil {
				http.Error(res, err.Error(), http.StatusBadRequest)
			}
		}
		ip_address, _, _ := net.SplitHostPort(r.RemoteAddr)
		println("username: " + r.FormValue("username") + "  |  Password: " + r.FormValue("password"))

		var login models.Auth
		login = models.LoginUser(api, loginDetails.Username, loginDetails.Password, ip_address)

		if(login.Success) {
			data, err := json.Marshal(login)
			if err != nil {
				http.Error(res, err.Error(), http.StatusInternalServerError)
				return
			}
			res.Header().Set("Content-Type", "application/json")
			res.WriteHeader(http.StatusOK)
			res.Write(data)
		} else {
			res.Header().Set("Content-Type", "application/json")
			res.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(res).Encode(error_response{Message:"Invalid Username or Password"})
		}



	})
}