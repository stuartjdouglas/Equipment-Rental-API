package email
import (
	"github.com/go-gomail/gomail"
	"log"
	"github.com/remony/Equipment-Rental-API/core/router"
	"strconv"
)


type Email struct {
	Subject string `json:"string"`
	Receipt string `json:"receipt"`
	Body string `json:"body"`

}
// SendEmail sends an email using example from https://godoc.org/gopkg.in/gomail.v2#example-package
func SendEmail(api router.API, receipt string, subject string, body string) bool {
	m := gomail.NewMessage()
	m.SetHeader("From", "test@test.com")
	m.SetHeader("To", receipt)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)


	port, err := strconv.Atoi(api.Config.Development.Email.Port)
	if err != nil {
		log.Println(err)
	}


	d := gomail.NewPlainDialer(
		api.Config.Development.Email.Hostname,
		port,
		api.Config.Development.Email.Login,
		api.Config.Development.Email.Password,
	)

	if err := d.DialAndSend(m); err != nil {
		log.Println(err)
		return false
	}
	return true;
}