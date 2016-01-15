package routes
import(

	"encoding/json"
	"net/http"
	"github.com/zenazn/goji/web"
	"strings"
	"github.com/remony/Equipment-Rental-API/core/router"
	"github.com/remony/Equipment-Rental-API/core/models/sessions"
	"github.com/remony/Equipment-Rental-API/core/database"
)
type tokenremoved struct {
	ID	string `json:"id"`
	Message string `json:"message"`
}

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

		var login database.Auth
		login = database.LoginUser(api, strings.ToLower(loginDetails.Username), loginDetails.Password)

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

	api.Router.Delete("/session", func (c web.C, res http.ResponseWriter, r *http.Request) {
		var idenf = r.Header.Get("id")
		var token = r.Header.Get("token")


		if (token != "" && idenf != "") {
			if (sessions.IsSessionValid(api, token)) {


				removal := sessions.DisableToken(api, idenf)

				if (removal) {

					data, err := json.Marshal(tokenremoved{ID:idenf, Message:"Session removed."})
					if err != nil {
						http.Error(res, err.Error(), http.StatusInternalServerError)
						return
					}
					res.Header().Set("Content-Type", "application/json")
					res.WriteHeader(http.StatusOK)
					res.Write(data)
				} else {
					res.Header().Set("Content-Type", "application/json")
					res.WriteHeader(http.StatusInternalServerError)
					json.NewEncoder(res).Encode(error_response{Message:"Server error"})
				}
			} else {
				res.Header().Set("Content-Type", "application/json")
				res.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(res).Encode(error_response{Message:"Invalid Username or Password"})
			}
		}else {
			res.Header().Set("Content-Type", "application/json")
			res.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(res).Encode(error_response{Message:"Missing parameters id and/or token"})
		}


	})
}