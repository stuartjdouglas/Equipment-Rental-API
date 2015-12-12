package routes
import(

	"encoding/json"
"net/http"
	"net"
	"github.com/zenazn/goji/web"
	"strings"
	"github/remony/Equipment-Rental-API/core/router"
	"github/remony/Equipment-Rental-API/core/models"
)

func generateAuthRoutes(api router.API)	{
	api.Router.Post("/logout", func (c web.C, res http.ResponseWriter, r *http.Request) {
		//		Not yet implemented
		//		Call method to remove token
	})




	api.Router.Post("/login", func (c web.C, res http.ResponseWriter, r *http.Request) {
		var loginDetails = login{
			Username: r.FormValue("username"),
			Password: r.FormValue("password"),
		}

		if len(loginDetails.Username) == 0 || len(loginDetails.Password) == 0{
			err := json.NewDecoder(r.Body).Decode(&loginDetails)
			if err != nil {
				http.Error(res, err.Error(), http.StatusBadRequest)
			}
		}
//		TODO Remove IP address recording
		ip_address, _, _ := net.SplitHostPort(r.RemoteAddr)

		var login models.Auth
		login = models.LoginUser(api, strings.ToLower(loginDetails.Username), loginDetails.Password, ip_address)

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