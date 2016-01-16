package database

import (
	"fmt"
	"github.com/remony/Equipment-Rental-API/core/router"
)

func GetHello(api router.API, token string) hello {

	author := GetUsername(api, GetUserIdFromToken(api, token))
	message := fmt.Sprintf("こんにちは, %s!", author)

	return hello{Message:message}

}
