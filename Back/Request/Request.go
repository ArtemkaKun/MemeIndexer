package Request

import (
	"Back/Error"
	"bytes"
	"io/ioutil"
	"net/http"
)

const DBServerAddress = "http://localhost:8001/"

func PrepareJSONRequest(requestType string, requestEndpoint string, memeDataInJSON []byte) (request *http.Request, errorMessage string) {
	request, err := http.NewRequest(requestType, DBServerAddress+requestEndpoint, bytes.NewBuffer(memeDataInJSON))
	if err != nil {
		return nil, Error.HandleCommonError(err)
	}
	request.Header.Set("Content-Type", "application/json")
	return
}

func MakeAuthRequestToDBServer(request string) (errorMessage string) {
	response, err := http.Get(DBServerAddress + request)
	if err != nil {
		return Error.HandleCommonError(err)
	}
	errorMessage = CheckResponseStatusCode(response)
	return
}

func MakeRequestToDBServer(request *http.Request) (serverResponse []byte, errorMessage string) {
	httpClient := &http.Client{}
	response, err := httpClient.Do(request)
	if err != nil {
		return nil, Error.HandleCommonError(err)
	}
	defer response.Body.Close()

	errorMessage = CheckResponseStatusCode(response)
	if errorMessage != "" {
		return
	}

	serverResponse, _ = ioutil.ReadAll(response.Body)
	return
}

func CheckResponseStatusCode(response *http.Response) (errorMessage string) {
	if response.StatusCode != 200 {
		errorFromHeader := response.Header.Get("err")
		return Error.HandleDBServerError(&errorFromHeader)
	}
	return
}
