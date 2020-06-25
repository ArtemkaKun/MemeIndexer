package Objects

import (
	"Back/Structures"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	userData Structures.UserAuthData
}

func (user *User) AuthenticateUser(context *gin.Context) {
	user.getAuthDataFromRequest(context)
	authRequest := prepareAuthRequest(user.userData)
	errorMessage := MakeAuthRequestToDBServer(authRequest)

	if errorMessage != "" {
		context.String(http.StatusInternalServerError, errorMessage)
		return
	}

	context.Status(http.StatusOK)
}

func (user *User) getAuthDataFromRequest(context *gin.Context) {
	user.userData.Login = context.Query("login")
	user.userData.Pass = context.Query("pass")
}

func prepareAuthRequest(authData Structures.UserAuthData) (authRequest string) {
	authRequest = fmt.Sprintf("userAuth?login=%v&pass=%v", authData.Login, authData.Pass)
	return
}
