package API

import (
	"DBServer/DB"
	"DBServer/MemeSeeker"
	"DBServer/Structures"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var Router *gin.Engine

func init() {
	initRouter()
	setRouterPoints()
}

func initRouter() {
	gin.SetMode(gin.ReleaseMode)
	Router = gin.Default()
}

func setRouterPoints() {
	Router.GET("/userAuth", authenticateUser)

	Router.GET("/meme", findMeme)
	Router.POST("/meme", addMeme)
}

func authenticateUser(context *gin.Context) {
	receivedUserAuthData := getUserAuthDataForRequest(context)

	userId, err := DB.FindUser(receivedUserAuthData)
	if err != nil {
		sendUnauthorizedError(context, err)
	} else {
		context.String(http.StatusOK, "userId", userId)
	}
}

func getUserAuthDataForRequest(context *gin.Context) (receivedUserAuthData Structures.UserAuthData) {
	receivedUserAuthData = Structures.UserAuthData{
		Login:    context.Query("login"),
		Password: context.Query("pass"),
	}
	return
}

func sendUnauthorizedError(context *gin.Context, err error) {
	log.Println(err)
	context.Status(http.StatusUnauthorized)
}

func findMeme(context *gin.Context) {
	receivedMemeTags, err := getMemeTagsFromRequest(context)
	if err != nil {
		sendInternalServerError(context, err)
		return
	}

	memeFilePath, err := MemeSeeker.FindMeme(receivedMemeTags)
	if err != nil {
		sendInternalServerError(context, err)
		return
	}

	const MemeImageFolderPath string = "./MemeImages/"
	context.File(MemeImageFolderPath + memeFilePath)
}

func getMemeTagsFromRequest(context *gin.Context) (receivedMemeTags Structures.MemeTags, err error) {
	err = context.BindJSON(&receivedMemeTags)
	return
}

func addMeme(context *gin.Context) {

}

func sendInternalServerError(context *gin.Context, err error) {
	log.Println(err)
	context.Status(http.StatusInternalServerError)
}
