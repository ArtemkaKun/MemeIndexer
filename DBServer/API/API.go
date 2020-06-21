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
const MemeImageFolderPath string = "./MemeImages/"

func init() {
	gin.SetMode(gin.ReleaseMode)
	Router = gin.Default()

	setRoutingPoints()
}

func setRoutingPoints() {
	Router.GET("/userAuth", authenticateUser)

	Router.GET("/meme", findMeme)
}

func authenticateUser(context *gin.Context) {
	receivedUserAuthData := Structures.UserAuthData{
		Login: context.Query("login"),
		Password: context.Query("pass"),
	}

	userId, err := DB.FindUser(receivedUserAuthData)
	if err != nil {
		log.Println(err)
		context.Status(http.StatusUnauthorized)
	} else {
		context.String(http.StatusOK, "userId", userId)
	}
}

func findMeme(context *gin.Context) {
	var receivedMemeTags Structures.MemeTags

	err := context.BindJSON(&receivedMemeTags)
	if err != nil {
		sendInternalServerError(context, err)
		return
	}

	memeFilePath, err := MemeSeeker.FindMeme(receivedMemeTags)
	if err != nil {
		sendInternalServerError(context, err)
		return
	}

	context.File(MemeImageFolderPath + memeFilePath)
}

func sendInternalServerError(context *gin.Context, err error) {
	log.Println(err)
	context.Status(http.StatusInternalServerError)
}