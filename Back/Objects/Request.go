package Objects

import (
	"net/http"
)

const DBServerAddress = "http://localhost:8001/"

func MakeGetRequestToDBServer(request string) (errorMessage string){
	response, err := http.Get(DBServerAddress + request)
	if err != nil {
		return HandleCommonError(err)
	}

	if response.StatusCode != 200 {
		errorFromHeader := response.Header.Get("err")
		return HandleDBServerError(&errorFromHeader)
	}

	return
}
